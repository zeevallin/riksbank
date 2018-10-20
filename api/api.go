package api

import (
	"encoding/csv"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/zeeraw/riksbank/api/calc"
	"github.com/zeeraw/riksbank/api/separators"
	"github.com/zeeraw/riksbank/api/series"
)

const (
	// Schema is the API host schema
	Schema = "https"
	// Host is the API host of Riksbanken
	Host = "www.riksbank.se"
	// SVPath is the API path for exchange rates in swedish
	SVPath = "/sv/statistik/sok-rantor--valutakurser/"

	// ENPath is the API path for exchange rates in english
	ENPath = "/en-gb/statistics/search-interest--exchange-rates/"

	exportKey = "export"
	exportVal = "csv"

	fromKey = "from"
	toKey   = "to"

	seriesVal = "on"
)

// RatesRequest represents the API request parameters for rates
type RatesRequest struct {
	Series  []series.Series
	Methods []calc.Method
	Period  calc.Period
	From    Date
	To      Date
}

// RatesResponse represents the API reponse for rates
type RatesResponse struct {
	Rates Rates
}

// Rates represents the association between a series and its value
type Rates map[series.Series]*Rate

// Rate represents a rate for a given series on a given period
type Rate struct {
	Period string
	Group  string
	Series string

	Average float64
	Min     float64
	Max     float64
	Ultimo  float64
}

// GetRates returns the rates for the given series
func GetRates(req *RatesRequest) (res *RatesResponse, err error) {
	_, err = ratesRequest(req)
	if err != nil {
		return nil, err
	}
	return &RatesResponse{}, nil
}

func ratesRequest(r *RatesRequest) (interface{}, error) {
	if err := ValidatePeriodMethods(r.Period, r.Methods...); err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Add(separators.Key, separators.Comma.String())
	query.Add(calc.PeriodKey, r.Period.String())
	query.Add(fromKey, r.From.String())

	for _, s := range r.Series {
		query.Add(s.String(), seriesVal)
	}

	for _, m := range r.Methods {
		query.Add(calc.MethodKey, m.String())
	}

	query.Add(toKey, r.To.String())
	query.Add(exportKey, exportVal)

	u := &url.URL{
		Path:     ENPath,
		Scheme:   Schema,
		Host:     Host,
		RawQuery: query.Encode(),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "*")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// bytes, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Print(string(bytes))

	reader := csv.NewReader(res.Body)
	reader.Comma = ';'

	var fields []string

	i := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {

			return nil, err
		}

		if i == 0 {
			for idx, r := range record {
				record[idx] = strings.TrimSpace(r)
				record[idx] = strings.TrimLeft(record[idx], "\ufeff")
				record[idx] = strings.TrimRight(record[idx], "\ufeff")
			}
			fields = record
			i++
			continue
		}

		rate := &Rate{}
		for idx, val := range record {
			switch fields[idx] {
			case "Period":
				rate.Period = val
			case "Group":
				rate.Group = val
			case "Series":
				rate.Series = val
			case "Average", "Min", "Max", "Ultimo":
				val = strings.Replace(val, ",", ".", -1)
				f, err := strconv.ParseFloat(val, 64)
				if err != nil {
					return nil, err
				}
				switch fields[idx] {
				case "Average":
					rate.Average = f
				case "Min":
					rate.Min = f
				case "Max":
					rate.Max = f
				case "Ultimo":
					rate.Ultimo = f
				}
			}
		}

		spew.Dump(rate)

		i++
	}

	return nil, nil
}
