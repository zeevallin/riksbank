package swea

import (
	"context"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/zeeraw/riksbank/swea/internal/validutf8"
	"github.com/zeeraw/riksbank/swea/internal/xmlstrings"
	"github.com/zeeraw/riksbank/swea/responses"
)

const (
	scheme      = "http"
	host        = "swea.riksbank.se"
	path        = "/sweaWS/services/SweaWebServiceHttpSoap12Endpoint"
	contentType = "text/xml"
)

var (
	sweaURL = &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
)

// Config represents the configuration for the Riksbank API client
type Config struct {
	HTTPClient *http.Client
}

// LiveAPI is a collection of the methods for the live Riksbank API
type LiveAPI struct {
	client *http.Client
	url    *url.URL
}

// New constructs and returns a new Swea client
func New(config Config) *LiveAPI {
	// Setup the HTTP client
	client := http.DefaultClient
	if config.HTTPClient != nil {
		client = config.HTTPClient
	}
	return &LiveAPI{
		client: client,
		url:    sweaURL,
	}
}

// GetCalendarDays returns the business days between two dates
func (api *LiveAPI) GetCalendarDays(ctx context.Context, req *GetCalendarDaysRequest) (*GetCalendarDaysResponse, error) {
	body, err := build(tmpl("get_calendar_days"), req)
	if err != nil {
		return nil, err
	}
	env := &responses.GetCalendarDaysResponseEnvelope{}
	err = api.call(ctx, body, env)
	if err != nil {
		return nil, err
	}
	res := &GetCalendarDaysResponse{
		From: req.From,
		To:   req.To,
		Days: make([]DayInfo, len(env.Body.GetCalendarDaysResponse.Return)),
	}
	for idx, r := range env.Body.GetCalendarDaysResponse.Return {
		date := xmlstrings.ParseDate(r.Caldate.Text)
		week, err := strconv.Atoi(r.Week.Text)
		if err != nil {
			return nil, err
		}
		weekYear, err := strconv.Atoi(r.Weekyear.Text)
		if err != nil {
			return nil, err
		}
		res.Days[idx] = DayInfo{
			Date:      date,
			Week:      week,
			WeekYear:  weekYear,
			IsBankDay: xmlstrings.ParseBool(r.Bankday.Text),
		}
	}
	return res, nil
}

// GetAllCrossNames returns the series names for exhcnage rates to SEK
func (api *LiveAPI) GetAllCrossNames(ctx context.Context, req *GetAllCrossNamesRequest) (*GetAllCrossNamesResponse, error) {
	body, err := build(tmpl("get_all_cross_names"), req)
	if err != nil {
		return nil, err
	}
	env := &responses.GetAllCrossNamesResponseEnvelope{}
	err = api.call(ctx, body, env)
	if err != nil {
		return nil, err
	}
	res := &GetAllCrossNamesResponse{
		Language: req.Language,
		Series:   make([]CrossSeriesInfo, len(env.Body.GetAllCrossNamesResponse.Return)),
	}
	for idx, r := range env.Body.GetAllCrossNamesResponse.Return {
		res.Series[idx] = CrossSeriesInfo{
			ID:          r.Seriesid.Text,
			Name:        r.Seriesname.Text,
			Description: r.Seriesdescription.Text,
		}
	}
	return res, nil
}

// GetCrossRates returns the exchange rates for series
func (api *LiveAPI) GetCrossRates(ctx context.Context, req *GetCrossRatesRequest) (*GetCrossRatesResponse, error) {
	body, err := build(tmpl("get_cross_rates"), req)
	if err != nil {
		return nil, err
	}
	env := &responses.GetCrossRatesResponseEnvelope{}
	err = api.call(ctx, body, env)
	if err != nil {
		return nil, err
	}
	res := &GetCrossRatesResponse{
		CrossPairs:      req.CrossPairs,
		From:            req.From,
		To:              req.To,
		Language:        req.Language,
		AggregateMethod: req.AggregateMethod,
	}

	var crossRates = []CrossRateInfo{}
	for _, s := range env.Body.GetCrossRatesResponse.Return.Groups.Series {
		for _, rr := range s.Resultrows {
			date, period := xmlstrings.ParseDatePeriod(rr.Date, rr.Period.Text)
			var value string
			switch req.AggregateMethod {
			case Weekly, Monthly, Quarterly, Yearly:
				value = strings.TrimSpace(rr.Average.Text)
			default:
				value = strings.TrimSpace(rr.Value)
			}
			cr := CrossRateInfo{
				Base:    ParseCurrency(strings.TrimSpace(s.Seriesid1)),
				Counter: ParseCurrency(strings.TrimSpace(s.Seriesid2)),
				Date:    date,
				Value:   value,
				Period:  period,
			}
			crossRates = append(crossRates, cr)
		}
	}
	res.CrossRates = crossRates
	return res, nil
}

// GetInterestAndExchangeRates returns values that are aggregated and grouped according to selected aggregate method
func (api *LiveAPI) GetInterestAndExchangeRates(ctx context.Context, req *GetInterestAndExchangeRatesRequest) (*GetInterestAndExchangeRatesResponse, error) {
	body, err := build(tmpl("get_interest_and_exchange_rates"), req)
	if err != nil {
		return nil, err
	}
	env := &responses.GetInterestAndExchangeRatesResponseEnvelope{}
	err = api.call(ctx, body, env)
	if err != nil {
		return nil, err
	}
	res := &GetInterestAndExchangeRatesResponse{
		From:            req.From,
		To:              req.To,
		Average:         req.Average,
		Min:             req.Min,
		Max:             req.Max,
		Language:        req.Language,
		AggregateMethod: req.AggregateMethod,
		Series:          req.Series,
	}
	ret := env.Body.GetInterestAndExchangeRatesResponse.Return
	rates := []RateInfo{}
	for _, gs := range ret.Groups {
		groupID := strings.TrimSpace(gs.Groupid.Text)
		groupName := strings.TrimSpace(gs.Groupname.Text)
		for _, ss := range gs.Series {
			seriesID := strings.TrimSpace(ss.Seriesid.Text)
			seriesName := strings.TrimSpace(ss.Seriesname.Text)
			for _, rr := range ss.Resultrows {
				date, period := xmlstrings.ParseDatePeriod(rr.Date.Text, rr.Period.Text)
				ri := RateInfo{
					GroupID:    groupID,
					GroupName:  groupName,
					SeriesID:   seriesID,
					SeriesName: seriesName,
					Date:       date,
					Period:     period,
					Average:    strings.TrimSpace(rr.Average.Text),
					Min:        strings.TrimSpace(rr.Min.Text),
					Max:        strings.TrimSpace(rr.Max.Text),
					Ultimo:     strings.TrimSpace(rr.Ultimo.Text),
					Value:      strings.TrimSpace(rr.Value.Text),
				}
				rates = append(rates, ri)
			}
		}
	}
	res.Rates = rates
	return res, nil
}

// GetInterestAndExchangeGroupNames returns all the groups of interest and exchange rates
func (api *LiveAPI) GetInterestAndExchangeGroupNames(ctx context.Context, req *GetInterestAndExchangeGroupNamesRequest) (*GetInterestAndExchangeGroupNamesResponse, error) {
	body, err := build(tmpl("get_interest_and_exchange_group_names"), req)
	if err != nil {
		return nil, err
	}
	env := &responses.GetInterestAndExchangeGroupNamesResponseEnvelope{}
	err = api.call(ctx, body, env)
	if err != nil {
		return nil, err
	}
	groups := make(GroupsInfo, len(env.Body.GetInterestAndExchangeGroupNamesResponse.Return))
	for idx, g := range env.Body.GetInterestAndExchangeGroupNamesResponse.Return {
		groups[idx] = GroupInfo{
			ID:          strings.TrimSpace(g.Groupid.Text),
			ParentID:    strings.TrimSpace(g.Parentgroupid.Text),
			Name:        strings.TrimSpace(g.Groupname.Text),
			Description: strings.TrimSpace(g.Groupdescription.Text),
		}
	}
	sort.Sort(groups)
	res := &GetInterestAndExchangeGroupNamesResponse{
		Language: req.Language,
		Groups:   groups,
	}
	return res, nil
}

// GetInterestAndExchangeNames returns the series for the selected group
func (api *LiveAPI) GetInterestAndExchangeNames(ctx context.Context, req *GetInterestAndExchangeNamesRequest) (*GetInterestAndExchangeNamesResponse, error) {
	body, err := build(tmpl("get_interest_and_exchange_names"), req)
	if err != nil {
		return nil, err
	}
	env := &responses.GetInterestAndExchangeNamesResponseEnvelope{}
	err = api.call(ctx, body, env)
	if err != nil {
		return nil, err
	}
	series := make([]SeriesInfo, len(env.Body.GetInterestAndExchangeNamesResponse.Return))
	for idx, s := range env.Body.GetInterestAndExchangeNamesResponse.Return {
		series[idx] = SeriesInfo{
			ID:      strings.TrimSpace(s.Seriesid.Text),
			GroupID: strings.TrimSpace(s.Groupid.Text),

			Name:            strings.TrimSpace(s.Shortdescription.Text),
			Description:     strings.TrimSpace(s.Description.Text),
			LongDescription: strings.TrimSpace(s.Longdescription.Text),
			Source:          strings.TrimSpace(s.Source.Text),
			Type:            strings.TrimSpace(s.Type.Text),
			From:            xmlstrings.ParseDate(s.Datefrom.Text),
			To:              xmlstrings.ParseDate(s.Dateto.Text),
		}
	}
	res := &GetInterestAndExchangeNamesResponse{
		GroupID:  req.GroupID,
		Language: req.Language,
		Series:   series,
	}
	return res, nil
}

func (api *LiveAPI) call(ctx context.Context, body io.Reader, v interface{}) error {
	// Build the request
	req, err := http.NewRequest(http.MethodPost, api.url.String(), body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", contentType)
	req = req.WithContext(ctx)

	// Perform the request
	res, err := api.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Read the response
	bts, err := ioutil.ReadAll(validutf8.NewReader(res.Body))
	if err != nil {
		return err
	}
	err = xml.Unmarshal(bts, v)
	return err
}
