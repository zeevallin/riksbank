package date_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/zeeraw/riksbank/date"
)

func Test_Parse(t *testing.T) {
	cases := []struct {
		expected time.Time
		date     string
		errors   bool
	}{
		{time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC), "2018-01-01", false},
		{time.Date(1982, 6, 30, 0, 0, 0, 0, time.UTC), "1982-06-30", false},
		{time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), "hello world", true},
	}
	for _, c := range cases {
		tm, err := date.Parse(c.date)
		if c.errors {
			notequal(t, nil, err)
		} else {
			equal(t, c.expected, tm)
		}
	}
}

func Test_ParseSafe(t *testing.T) {
	cases := []struct {
		expected time.Time
		date     string
	}{
		{time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC), "2018-01-01"},
		{time.Date(1982, 6, 30, 0, 0, 0, 0, time.UTC), "1982-06-30"},
		{time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), "hello world"},
	}
	for _, c := range cases {
		equal(t, c.expected, date.ParseSafe(c.date))
	}
}

func Test_Format(t *testing.T) {
	cases := []struct {
		time     time.Time
		expected string
	}{
		{time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC), "2018-01-01"},
		{time.Date(1982, 6, 30, 0, 0, 0, 0, time.UTC), "1982-06-30"},
	}
	for _, c := range cases {
		equal(t, c.expected, date.Format(c.time))
	}
}

func equal(t testing.TB, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nexpected: %v\n     got: %v", expected, actual)
	}
}

func notequal(t testing.TB, expected, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("\nexpected: %v\n     got: %v", expected, actual)
	}
}
