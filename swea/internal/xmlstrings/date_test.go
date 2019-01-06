package xmlstrings_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeeraw/riksbank/swea/internal/xmlstrings"
)

func Test_ParseDate(t *testing.T) {
	cases := []struct {
		value    string
		expected time.Time
	}{
		{"2018-01-01", time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"1991-09-13", time.Date(1991, 9, 13, 0, 0, 0, 0, time.UTC)},
		{"foobar", time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)},
	}
	for _, c := range cases {
		t.Run(c.value, func(t *testing.T) {
			actual := xmlstrings.ParseDate(c.value)
			if actual != c.expected {
				t.Errorf("expected %q to be %v, was %v", c.value, c.expected, actual)
			}
		})
	}
}

func Test_ParseDatePeriod(t *testing.T) {
	cases := []struct {
		date           string
		period         string
		expectedPeriod string
		expectedDate   time.Time
	}{
		{"2018-01-01", "2018-01-01", "2018-01-01", time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"1991-09-13", "", "1991-09-13", time.Date(1991, 9, 13, 0, 0, 0, 0, time.UTC)},
		{"foobar", "", "-0001-11-30", time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%s & %s", c.date, c.period), func(t *testing.T) {
			date, period := xmlstrings.ParseDatePeriod(c.date, c.period)
			if date != c.expectedDate {
				t.Errorf("expected %q to be %q, was %q", c.date, c.expectedDate, date)
			}
			if period != c.expectedPeriod {
				t.Errorf("expected %q to be %q, was %q", c.period, c.expectedPeriod, period)
			}
		})
	}
}
