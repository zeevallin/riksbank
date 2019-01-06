package swea

import (
	"context"
)

// Swea represents the available methods in the Swea API
type Swea interface {
	GetCalendarDays(ctx context.Context, req *GetCalendarDaysRequest) (*GetCalendarDaysResponse, error)
	GetAllCrossNames(ctx context.Context, req *GetAllCrossNamesRequest) (*GetAllCrossNamesResponse, error)
	GetCrossRates(ctx context.Context, req *GetCrossRatesRequest) (*GetCrossRatesResponse, error)
	GetInterestAndExchangeRates(ctx context.Context, req *GetInterestAndExchangeRatesRequest) (*GetInterestAndExchangeRatesResponse, error)
	GetInterestAndExchangeGroupNames(ctx context.Context, req *GetInterestAndExchangeGroupNamesRequest) (*GetInterestAndExchangeGroupNamesResponse, error)
	GetInterestAndExchangeNames(ctx context.Context, req *GetInterestAndExchangeNamesRequest) (*GetInterestAndExchangeNamesResponse, error)
}
