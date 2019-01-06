package period_test

import (
	"testing"

	"github.com/zeeraw/riksbank/period"
)

func Test_Parse(t *testing.T) {
	cases := []struct {
		value    string
		expected string
	}{
		{"2018", "2018"},
		{"2018 Quarter 1", "2018 Quarter 1"},
		{"2018 January", "2018 January"},
		{"2018 Week 1", "2018 Week 1"},
		{"2018-01-01", "2018-01-01"},
	}
	for _, c := range cases {
		actual := period.Parse(c.value).String()
		if actual != c.expected {
			t.Errorf("expected %q was %q", c.expected, actual)
		}
	}
}
