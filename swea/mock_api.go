package swea

import "context"

// MockAPI is a dud API that returns dud structs
type MockAPI struct {
	GetCalendarDaysResponse                  *GetCalendarDaysResponse
	GetAllCrossNamesResponse                 *GetAllCrossNamesResponse
	GetCrossRatesResponse                    *GetCrossRatesResponse
	GetInterestAndExchangeRatesResponse      *GetInterestAndExchangeRatesResponse
	GetInterestAndExchangeGroupNamesResponse *GetInterestAndExchangeGroupNamesResponse
	GetInterestAndExchangeNamesResponse      *GetInterestAndExchangeNamesResponse
}

// NewMock creates a default mock
func NewMock() *MockAPI {
	return &MockAPI{
		GetCalendarDaysResponse: &GetCalendarDaysResponse{
			Days: []DayInfo{},
		},
		GetAllCrossNamesResponse: &GetAllCrossNamesResponse{
			Series: []CrossSeriesInfo{},
		},
		GetCrossRatesResponse: &GetCrossRatesResponse{
			CrossRates: []CrossRateInfo{},
		},
		GetInterestAndExchangeRatesResponse: &GetInterestAndExchangeRatesResponse{
			Rates: []RateInfo{},
		},
		GetInterestAndExchangeGroupNamesResponse: &GetInterestAndExchangeGroupNamesResponse{
			Groups: []GroupInfo{},
		},
		GetInterestAndExchangeNamesResponse: &GetInterestAndExchangeNamesResponse{
			Series: []SeriesInfo{},
		},
	}
}

// GetCalendarDays returns the GetCalendarDaysResponse from the mock
func (api *MockAPI) GetCalendarDays(ctx context.Context, req *GetCalendarDaysRequest) (*GetCalendarDaysResponse, error) {
	res := api.GetCalendarDaysResponse
	res.From = req.From
	res.To = req.To
	return res, nil
}

// GetAllCrossNames returns the GetAllCrossNamesResponse from the mock
func (api *MockAPI) GetAllCrossNames(ctx context.Context, req *GetAllCrossNamesRequest) (*GetAllCrossNamesResponse, error) {
	res := api.GetAllCrossNamesResponse
	res.Language = req.Language
	return res, nil
}

// GetCrossRates returns the GetCrossRatesResponse from the mock
func (api *MockAPI) GetCrossRates(ctx context.Context, req *GetCrossRatesRequest) (*GetCrossRatesResponse, error) {
	res := api.GetCrossRatesResponse
	res.CrossPairs = req.CrossPairs
	res.From = req.From
	res.To = req.To
	res.AggregateMethod = req.AggregateMethod
	res.Language = req.Language
	return res, nil
}

// GetInterestAndExchangeRates returns the GetInterestAndExchangeRatesResponse from the mock
func (api *MockAPI) GetInterestAndExchangeRates(ctx context.Context, req *GetInterestAndExchangeRatesRequest) (*GetInterestAndExchangeRatesResponse, error) {
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
func (api *MockAPI) GetInterestAndExchangeGroupNames(ctx context.Context, req *GetInterestAndExchangeGroupNamesRequest) (*GetInterestAndExchangeGroupNamesResponse, error) {
	res := api.GetInterestAndExchangeGroupNamesResponse
	res.Language = req.Language
	return res, nil
}

// GetInterestAndExchangeNames returns the GetInterestAndExchangeNamesResponse from the mock
func (api *MockAPI) GetInterestAndExchangeNames(ctx context.Context, req *GetInterestAndExchangeNamesRequest) (*GetInterestAndExchangeNamesResponse, error) {
	res := api.GetInterestAndExchangeNamesResponse
	res.Language = req.Language
	return res, nil
}
