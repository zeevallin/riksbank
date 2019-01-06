package riksbank

import (
	"time"

	"github.com/zeeraw/riksbank/currency"
	"github.com/zeeraw/riksbank/period"
)

// RateGroup represents a group for a rate
type RateGroup struct {
	ID   string
	Name string
}

// RateSeries represents a series for a rate
type RateSeries struct {
	ID   string
	Name string
}

// Rate represents a interest or exchange series rate for a period
type Rate struct {
	Group  RateGroup
	Series RateSeries
	Date   time.Time
	Period period.Period
	Value  *float64
}

// Rates represents multiple instances of Rate
type Rates []Rate

// Day represents information about a specific day in banking context
type Day struct {
	Date      time.Time
	Year      int
	Week      int
	IsBankDay bool
}

// Days represents multiple instances of Day
type Days []Day

// Group represents information about an interest and exchange rate grouping
type Group struct {
	ID          string
	ParentID    string
	Name        string
	Description string
}

// Groups represents multiple instances of Group
type Groups []Group

// ExchangeRate represents an exchange rate between two currencies
type ExchangeRate struct {
	Date    time.Time
	Period  period.Period
	Base    currency.Currency
	Counter currency.Currency
	Value   *float64
}

// ExchangeRates represents multiple instance of ExchangeRate
type ExchangeRates []ExchangeRate

// ExchangeCurrency represents the exchange information for a currency
type ExchangeCurrency struct {
	SeriesID    string
	Code        string
	Currency    currency.Currency
	Description string
}

// ExchangeCurrencies represents multiple instances of ExchangeCurrency
type ExchangeCurrencies []ExchangeCurrency

// Series represents information about a series
type Series struct {
	ID              string
	GroupID         string
	Name            string
	Description     string
	LongDescription string
	Source          string
	From            time.Time
	To              time.Time
}

// SeriesGroup represents a group of series and information about the group
type SeriesGroup struct {
	Group  Group
	Series []Series
}

// SeriesGroups represent multiple instances of SeriesGroup
type SeriesGroups []SeriesGroup
