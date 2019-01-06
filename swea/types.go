package swea

import (
	"cloud.google.com/go/civil"
)

// Language is the container for a language ID
type Language string

const (
	// English is the english language ID
	English = Language("en")

	// Swedish is the swedish language ID
	Swedish = Language("sv")
)

// Series is the container for a series ID
type Series string

// ToCurrency converts a series into a currency
func (s Series) ToCurrency() Currency {
	return ParseCurrency(string(s))
}

// SearchGroupSeries represents searchable group series
type SearchGroupSeries struct {
	GroupID  string
	SeriesID string
}

// Day represents a date in the context of the central bank
type Day struct {
	Date      civil.Date
	Week      int
	WeekYear  int
	IsBankDay bool
}

// CrossRate is an exchange rate between two currencies
type CrossRate struct {
	Base    Currency
	Counter Currency
	Date    civil.Date
	Period  string
	Value   string
}

// CrossPair are the series to compare in a currency exchange
type CrossPair struct {
	Base    Series
	Counter Series
}

// ToCurrencyPair converts
func (cp CrossPair) ToCurrencyPair() CurrencyPair {
	return CurrencyPair{
		Base:    cp.Base.ToCurrency(),
		Counter: cp.Counter.ToCurrency(),
	}
}

// CrossSeriesInfo represents a interest or currency conversion series information
type CrossSeriesInfo struct {
	ID          string
	Name        string
	Description string
}

// SeriesInfo represents a interest or currency conversion series information
type SeriesInfo struct {
	ID              string
	GroupID         string
	Name            string
	Description     string
	LongDescription string
	Source          string
	Type            string
	From            *civil.Date
	To              *civil.Date
}

// RateInfo represents information about a rate for a series in a period
type RateInfo struct {
	GroupID   string
	GroupName string

	SeriesID   string
	SeriesName string

	Date   civil.Date
	Period string

	Average string
	Min     string
	Max     string
	Ultimo  string
	Value   string
}

// GroupInfo represents a grouping of interest or exchange rates
type GroupInfo struct {
	ID       string
	ParentID string

	Name        string
	Description string
}

// GetCalendarDaysRequest represents the parameters to get all business days between two dates
type GetCalendarDaysRequest struct {
	From civil.Date
	To   civil.Date
}

// GetCalendarDaysResponse contains the
type GetCalendarDaysResponse struct {
	From civil.Date
	To   civil.Date
	Days []Day
}

// GetAllCrossNamesRequest represents the parameters get all the exchange rate series suitable for cross rate names
type GetAllCrossNamesRequest struct {
	Language Language
}

// GetAllCrossNamesResponse contains the currency conversion series
type GetAllCrossNamesResponse struct {
	Language Language
	Series   []CrossSeriesInfo
}

// GetCrossRatesRequest represents the parameters to get all change rates
type GetCrossRatesRequest struct {
	CrossPairs []CrossPair

	From            civil.Date
	To              civil.Date
	Language        Language
	AggregateMethod AggregateMethod
}

// GetCrossRatesResponse contains exchange rates
type GetCrossRatesResponse struct {
	CrossRates []CrossRate
	CrossPairs []CrossPair

	From            civil.Date
	To              civil.Date
	Language        Language
	AggregateMethod AggregateMethod
}

// GetInterestAndExchangeRatesRequest represents the parameters to get exchange and interest rates
type GetInterestAndExchangeRatesRequest struct {
	Series []SearchGroupSeries

	From            civil.Date
	To              civil.Date
	Language        Language
	AggregateMethod AggregateMethod

	Average bool
	Min     bool
	Max     bool
	Ultimo  bool
}

// GetInterestAndExchangeRatesResponse contains interest and exchange rates
type GetInterestAndExchangeRatesResponse struct {
	Rates []RateInfo

	Series []SearchGroupSeries

	From            civil.Date
	To              civil.Date
	Language        Language
	AggregateMethod AggregateMethod

	Average bool
	Min     bool
	Max     bool
	Ultimo  bool
}

// GetInterestAndExchangeGroupNamesRequest represents the parameters to get a list of all groups
type GetInterestAndExchangeGroupNamesRequest struct {
	Language Language
}

// GetInterestAndExchangeGroupNamesResponse contains all groups
type GetInterestAndExchangeGroupNamesResponse struct {
	Groups   []GroupInfo
	Language Language
}

// GetInterestAndExchangeNamesRequest represents the parameters to get all series for a group
type GetInterestAndExchangeNamesRequest struct {
	GroupID  string
	Language Language
}

// GetInterestAndExchangeNamesResponse contains all series for a group
type GetInterestAndExchangeNamesResponse struct {
	Series   []SeriesInfo
	GroupID  string
	Language Language
}
