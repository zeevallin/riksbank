package xmlstrings_test

import (
	"testing"

	"github.com/zeeraw/riksbank/swea/internal/xmlstrings"
)

func Test_ParseBool(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{"YES", true},
		{"Y", true},
		{"1", true},
		{"TRUE", true},
		{"NO", false},
		{"N", false},
		{"0", false},
		{"FALSE", false},
		{"SOMETHING", false},
		{"  true \n", true},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			actual := xmlstrings.ParseBool(c.value)
			if actual != c.expected {
				t.Errorf("expected %q to be %v, was %v", c.value, c.expected, actual)
			}
		})
	}
}
