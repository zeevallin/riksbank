package xmlstrings

import (
	"strings"
	"time"
)

const (
	// DateLayout is the layout for the XML dates
	DateLayout = "2006-01-02"
)

// ParseDate attempts to turn a date string into a time object
func ParseDate(s string) time.Time {
	t, err := time.Parse(DateLayout, strings.TrimSpace(s))
	if err != nil {
		return time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	}
	return t
}

// ParseDatePeriod attempts to turn a date and priod into usable values
func ParseDatePeriod(d, p string) (time.Time, string) {
	p = strings.TrimSpace(p)
	date := ParseDate(d)
	if p == "" {
		p = date.Format(DateLayout)
	}
	return date, p
}
