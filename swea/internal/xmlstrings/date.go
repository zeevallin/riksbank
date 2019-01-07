package xmlstrings

import (
	"strings"
	"time"

	"github.com/zeeraw/riksbank/date"
)

const (
	// DateLayout is the layout for the XML dates
	DateLayout = "2006-01-02"
)

// ParseDate attempts to turn a date string into a time object
func ParseDate(s string) time.Time {
	return date.ParseSafe(strings.TrimSpace(s))
}

// ParseDatePeriod attempts to turn a date and priod into usable values
func ParseDatePeriod(d, p string) (time.Time, string) {
	p = strings.TrimSpace(p)
	t := ParseDate(d)
	if p == "" {
		p = date.Format(t)
	}
	return t, p
}
