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

// Day represents a date in the context of the central bank
type Day struct {
	Date      civil.Date
	Week      int
	WeekYear  int
	IsBankDay bool
}

// Series represents a interest or currency conversion series
type Series struct {
	ID          string
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
	Series   []Series
}
