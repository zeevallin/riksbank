package riksbank

import (
	"context"
	"strconv"
	"time"

	"github.com/zeeraw/riksbank/currency"
	"github.com/zeeraw/riksbank/period"
	"github.com/zeeraw/riksbank/swea"
)

// Config represents riksbank API configuration
type Config struct {
	SweaClient swea.Swea
}

// Riksbank represents the API client
type Riksbank struct {
	Swea swea.Swea
}

// New starts a new instance of riksbank
func New(config Config) *Riksbank {
	sweaClient := config.SweaClient
	if sweaClient == nil {
		sweaClient = swea.New(swea.Config{})
	}
	return &Riksbank{
		Swea: sweaClient,
	}
}

// RatesRequest represents parameters for requesting interest and exchange rates
type RatesRequest struct {
	Series          []string
	From            time.Time
	To              time.Time
	AggregateMethod AggregateMethod
	AnalysisMethod  AnalysisMethod
}

// RatesResponse represents the data retrieved from requesting interest and exchange rates
type RatesResponse struct {
	Rates Rates
}

// Rates return interest or change rates for one or more series between different dates
func (rb *Riksbank) Rates(ctx context.Context, req *RatesRequest) (*RatesResponse, error) {
	const (
		defaultGroupID = "3" // Set this to 3 until we're able to fetch series and groups
	)
	series := make([]swea.SearchGroupSeries, len(req.Series))
	for idx, s := range req.Series {
		series[idx] = swea.SearchGroupSeries{
			SeriesID: s,
			GroupID:  defaultGroupID,
		}
	}
	sweaReq := &swea.GetInterestAndExchangeRatesRequest{
		From:            req.From,
		To:              req.To,
		AggregateMethod: req.AggregateMethod.String(),
		Language:        swea.English,
		Series:          series,
	}
	switch req.AnalysisMethod {
	case Mean:
		sweaReq.Average = true
	case Min:
		sweaReq.Min = true
	case Max:
		sweaReq.Max = true
	case Ultimo:
		sweaReq.Ultimo = true
	}
	sweaRes, err := rb.Swea.GetInterestAndExchangeRates(ctx, sweaReq)
	if err != nil {
		return nil, err
	}
	rates := make(Rates, len(sweaRes.Rates))
	for idx, ri := range sweaRes.Rates {
		var value string
		switch req.AnalysisMethod {
		case Mean:
			value = ri.Average
		case Min:
			value = ri.Min
		case Max:
			value = ri.Max
		case Ultimo:
			value = ri.Ultimo
		default:
			value = ri.Value
		}
		rates[idx] = Rate{
			Date:   ri.Date,
			Period: period.Parse(ri.Period),
			Group: RateGroup{
				ID:   ri.GroupID,
				Name: ri.GroupName,
			},
			Series: RateSeries{
				ID:   ri.SeriesID,
				Name: ri.SeriesName,
			},
			Value: parseFloat(value),
		}
	}
	res := &RatesResponse{
		Rates: rates,
	}
	return res, nil
}

// ExchangeRatesRequest represents parameters for requesting exchange rates
type ExchangeRatesRequest struct {
	CurrencyPairs   []currency.Pair
	AggregateMethod AggregateMethod
	From            time.Time
	To              time.Time
}

// ExchangeRatesResponse represents the data retrieved from requesting exchange rates
type ExchangeRatesResponse struct {
	ExchangeRates []ExchangeRate
}

// ExchangeRates retrieves exchange rates for currency pairs
func (rb *Riksbank) ExchangeRates(ctx context.Context, req *ExchangeRatesRequest) (*ExchangeRatesResponse, error) {
	pairs := make([]swea.CrossPair, len(req.CurrencyPairs))
	for idx, cp := range req.CurrencyPairs {
		pairs[idx] = swea.CrossPair{
			BaseSeriesID:    cp.Base.Series(),
			CounterSeriesID: cp.Counter.Series(),
		}
	}
	sweaReq := &swea.GetCrossRatesRequest{
		CrossPairs:      pairs,
		AggregateMethod: req.AggregateMethod.String(),
		From:            req.From,
		To:              req.To,
		Language:        swea.English,
	}
	sweaRes, err := rb.Swea.GetCrossRates(ctx, sweaReq)
	if err != nil {
		return nil, err
	}
	exchangeRates := make(ExchangeRates, len(sweaRes.CrossRates))
	for idx, cr := range sweaRes.CrossRates {
		var value string
		switch req.AggregateMethod {
		case Daily:
			value = cr.Value
		default:
			value = cr.Average
		}
		exchangeRates[idx] = ExchangeRate{
			Date:    cr.Date,
			Period:  period.Parse(cr.Period),
			Base:    currency.Parse(cr.Base),
			Counter: currency.Parse(cr.Counter),
			Value:   parseFloat(value),
		}
	}
	res := &ExchangeRatesResponse{
		ExchangeRates: exchangeRates,
	}
	return res, nil
}

// ExchangeCurrenciesRequest represents parameters for requesting a list of all exchange currencies
type ExchangeCurrenciesRequest struct{}

// ExchangeCurrenciesResponse represents the data retrieved from requesting a list of all exchange currencies
type ExchangeCurrenciesResponse struct {
	Currencies ExchangeCurrencies
}

// ExchangeCurrencies retrieves exchange rates for currency pairs
func (rb *Riksbank) ExchangeCurrencies(ctx context.Context, req *ExchangeCurrenciesRequest) (*ExchangeCurrenciesResponse, error) {
	sweaReq := &swea.GetAllCrossNamesRequest{
		Language: swea.English,
	}
	sweaRes, err := rb.Swea.GetAllCrossNames(ctx, sweaReq)
	if err != nil {
		return nil, err
	}
	currencies := make(ExchangeCurrencies, len(sweaRes.Series))
	for idx, s := range sweaRes.Series {
		currencies[idx] = ExchangeCurrency{
			SeriesID:    s.ID,
			Currency:    currency.Parse(s.Name),
			Description: s.Description[4:],
		}
	}
	res := &ExchangeCurrenciesResponse{
		Currencies: currencies,
	}
	return res, nil
}

// DaysRequest represents parameters for requesting information about days
type DaysRequest struct {
	From time.Time
	To   time.Time
}

// DaysResponse represents the data retrieved from requesting information about days
type DaysResponse struct {
	Days Days
}

// Days retrieves exchange rates for currency pairs
func (rb *Riksbank) Days(ctx context.Context, req *DaysRequest) (*DaysResponse, error) {
	sweaReq := &swea.GetCalendarDaysRequest{
		From: req.From,
		To:   req.To,
	}
	sweaRes, err := rb.Swea.GetCalendarDays(ctx, sweaReq)
	if err != nil {
		return nil, err
	}
	days := make(Days, len(sweaRes.Days))
	for idx, day := range sweaRes.Days {
		days[idx] = Day{
			Date:      day.Date,
			IsBankDay: day.IsBankDay,
			Week:      day.Week,
			Year:      day.WeekYear,
		}
	}
	res := &DaysResponse{
		Days: days,
	}
	return res, nil
}

// GroupsRequest represents parameters for requesting a list of groups
type GroupsRequest struct{}

// GroupsResponse represents the data retrieved from requesting a list of groups
type GroupsResponse struct {
	Groups Groups
}

// Groups returns a list of all interest and exchange rate series groups
func (rb *Riksbank) Groups(ctx context.Context, req *GroupsRequest) (*GroupsResponse, error) {
	sweaRes, err := rb.Swea.GetInterestAndExchangeGroupNames(ctx, &swea.GetInterestAndExchangeGroupNamesRequest{
		Language: swea.English,
	})
	if err != nil {
		return nil, err
	}
	groups := make(Groups, len(sweaRes.Groups))
	for idx, group := range sweaRes.Groups {
		groups[idx] = sweaGroupToGroup(group)
	}
	res := &GroupsResponse{
		Groups: groups,
	}
	return res, nil
}

// SeriesRequest represents parameters for requesting grouped series
type SeriesRequest struct {
	Groups []string
}

// SeriesResponse represents the data retrieved from requesting grouped series
type SeriesResponse struct {
	Groups SeriesGroups
}

// Series returns a list of series grouped by their group
func (rb *Riksbank) Series(ctx context.Context, req *SeriesRequest) (*SeriesResponse, error) {
	sweaGRes, err := rb.Swea.GetInterestAndExchangeGroupNames(ctx, &swea.GetInterestAndExchangeGroupNamesRequest{
		Language: swea.English,
	})
	if err != nil {
		return nil, err
	}
	groups := SeriesGroups{}
	for _, g := range sweaGRes.Groups {
		if len(req.Groups) < 1 || isInSlice(req.Groups, g.ID) {
			sweaENReq := &swea.GetInterestAndExchangeNamesRequest{
				Language: swea.English,
				GroupID:  g.ID,
			}
			sweaENRes, err := rb.Swea.GetInterestAndExchangeNames(ctx, sweaENReq)
			if err != nil {
				return nil, err
			}
			series := make([]Series, len(sweaENRes.Series))
			for idx, s := range sweaENRes.Series {
				series[idx] = Series{
					ID:              s.ID,
					GroupID:         s.GroupID,
					Name:            s.Name,
					Description:     s.Description,
					LongDescription: s.LongDescription,
					Source:          s.Source,
					From:            s.From,
					To:              s.To,
				}
			}
			groups = append(groups, SeriesGroup{
				Group:  sweaGroupToGroup(g),
				Series: series,
			})
		}
	}
	res := &SeriesResponse{
		Groups: groups,
	}
	return res, nil
}

func parseFloat(s string) *float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil
	}
	return &f
}

func isInSlice(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func sweaGroupToGroup(g swea.GroupInfo) Group {
	return Group{
		ID:          g.ID,
		ParentID:    g.ParentID,
		Name:        g.Name,
		Description: g.Description,
	}
}
