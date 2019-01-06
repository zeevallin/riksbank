package mock

import (
	"context"

	"github.com/zeeraw/riksbank/swea"
)

// API is a dud API that returns predetermined structs
type API struct {
	GetCalendarDaysResponse                  *swea.GetCalendarDaysResponse
	GetAllCrossNamesResponse                 *swea.GetAllCrossNamesResponse
	GetCrossRatesResponse                    *swea.GetCrossRatesResponse
	GetInterestAndExchangeRatesResponse      *swea.GetInterestAndExchangeRatesResponse
	GetInterestAndExchangeGroupNamesResponse *swea.GetInterestAndExchangeGroupNamesResponse
	GetInterestAndExchangeNamesResponse      *swea.GetInterestAndExchangeNamesResponse
}

// New creates a default mock of the Swea API
func New() *API {
	return &API{
		GetCalendarDaysResponse: &swea.GetCalendarDaysResponse{
			Days: []swea.DayInfo{},
		},
		GetAllCrossNamesResponse: &swea.GetAllCrossNamesResponse{
			Series: []swea.CrossSeriesInfo{},
		},
		GetCrossRatesResponse: &swea.GetCrossRatesResponse{
			CrossRates: []swea.CrossRateInfo{},
		},
		GetInterestAndExchangeRatesResponse: &swea.GetInterestAndExchangeRatesResponse{
			Rates: []swea.RateInfo{},
		},
		GetInterestAndExchangeGroupNamesResponse: &swea.GetInterestAndExchangeGroupNamesResponse{
			Groups: []swea.GroupInfo{},
		},
		GetInterestAndExchangeNamesResponse: &swea.GetInterestAndExchangeNamesResponse{
			Series: []swea.SeriesInfo{},
		},
	}
}

// GetCalendarDays returns the GetCalendarDaysResponse from the mock
func (api *API) GetCalendarDays(ctx context.Context, req *swea.GetCalendarDaysRequest) (*swea.GetCalendarDaysResponse, error) {
	res := api.GetCalendarDaysResponse
	res.From = req.From
	res.To = req.To
	return res, nil
}

// GetAllCrossNames returns the GetAllCrossNamesResponse from the mock
func (api *API) GetAllCrossNames(ctx context.Context, req *swea.GetAllCrossNamesRequest) (*swea.GetAllCrossNamesResponse, error) {
	res := api.GetAllCrossNamesResponse
	res.Language = req.Language
	return res, nil
}

// GetCrossRates returns the GetCrossRatesResponse from the mock
func (api *API) GetCrossRates(ctx context.Context, req *swea.GetCrossRatesRequest) (*swea.GetCrossRatesResponse, error) {
	res := api.GetCrossRatesResponse
	res.CrossPairs = req.CrossPairs
	res.From = req.From
	res.To = req.To
	res.AggregateMethod = req.AggregateMethod
	res.Language = req.Language
	return res, nil
}

// GetInterestAndExchangeRates returns the GetInterestAndExchangeRatesResponse from the mock
func (api *API) GetInterestAndExchangeRates(ctx context.Context, req *swea.GetInterestAndExchangeRatesRequest) (*swea.GetInterestAndExchangeRatesResponse, error) {
	res := api.GetInterestAndExchangeRatesResponse
	res.Series = req.Series
	res.Average = req.Average
	res.Min = req.Min
	res.Max = req.Max
	res.Ultimo = req.Ultimo
	res.From = req.From
	res.To = req.To
	res.AggregateMethod = req.AggregateMethod
	res.Language = req.Language
	return res, nil
}

// GetInterestAndExchangeGroupNames returns the GetInterestAndExchangeGroupNamesResponse from the mock
func (api *API) GetInterestAndExchangeGroupNames(ctx context.Context, req *swea.GetInterestAndExchangeGroupNamesRequest) (*swea.GetInterestAndExchangeGroupNamesResponse, error) {
	res := api.GetInterestAndExchangeGroupNamesResponse
	res.Language = req.Language
	return res, nil
}

// GetInterestAndExchangeNames returns the GetInterestAndExchangeNamesResponse from the mock
func (api *API) GetInterestAndExchangeNames(ctx context.Context, req *swea.GetInterestAndExchangeNamesRequest) (*swea.GetInterestAndExchangeNamesResponse, error) {
	res := api.GetInterestAndExchangeNamesResponse
	res.Language = req.Language
	return res, nil
}
