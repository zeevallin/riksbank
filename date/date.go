package date

import (
	"time"
)

const (
	// DateLayout is the date format use in the context of the Swedish central bank
	DateLayout = "2006-01-02"
)

// Parse attempts to turn a string into to a date
func Parse(s string) (time.Time, error) {
	return time.Parse(DateLayout, s)
}

// ParseSafe attempts to turn a string into to a date and otherwise returns an empty time
func ParseSafe(s string) time.Time {
	t, err := time.Parse(DateLayout, s)
	if err != nil {
		return time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	}
	return t
}

// Format turns a time into a date string
func Format(t time.Time) string {
	return t.Format(DateLayout)
}
