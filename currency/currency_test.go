package currency_test

import (
	"testing"

	"github.com/zeeraw/riksbank/currency"
)

func Test_Parse(t *testing.T) {
	cases := []struct {
		value    string
		expected currency.Currency
	}{
		{"SEK", currency.Currency("SEK")},
	}
	for _, c := range cases {
		actual := currency.Parse(c.value)
		if actual != c.expected {
			t.Errorf("expected %q was %q", c.expected, actual)
		}
	}
}
func Test_Currency_Series(t *testing.T) {
	cases := []struct {
		currency currency.Currency
		expected string
	}{
		{currency.Currency("SEK"), "SEK"},
		{currency.Currency("NOK"), "SEKNOKPMI"},
		{currency.Currency("GBP"), "SEKGBPPMI"},
	}
	for _, c := range cases {
		actual := c.currency.Series()
		if actual != c.expected {
			t.Errorf("expected %q was %q", c.expected, actual)
		}
	}
}
func Test_ParsePair(t *testing.T) {
	cases := []struct {
		value    string
		expected currency.Pair
	}{
		{"SEK/NOK", currency.Pair{
			Base:    currency.Currency("SEK"),
			Counter: currency.Currency("NOK"),
		}},
		{"NOK:SEK", currency.Pair{
			Base:    currency.Currency("NOK"),
			Counter: currency.Currency("SEK"),
		}},
		{"EUR-GBP", currency.Pair{
			Base:    currency.Currency("EUR"),
			Counter: currency.Currency("GBP"),
		}},
	}
	for _, c := range cases {
		actual := currency.ParsePair(c.value)
		if actual != c.expected {
			t.Errorf("expected %q was %q", c.expected, actual)
		}
	}
}
func Test_Pair_String(t *testing.T) {
	cases := []struct {
		pair     currency.Pair
		expected string
	}{
		{currency.Pair{
			Base:    currency.Currency("SEK"),
			Counter: currency.Currency("NOK"),
		}, "SEK/NOK"},
		{currency.Pair{
			Base:    currency.Currency("NOK"),
			Counter: currency.Currency("SEK"),
		}, "NOK/SEK"},
		{currency.Pair{
			Base:    currency.Currency("EUR"),
			Counter: currency.Currency("GBP"),
		}, "EUR/GBP"},
	}
	for _, c := range cases {
		actual := c.pair.String()
		if actual != c.expected {
			t.Errorf("expected %q was %q", c.expected, actual)
		}
	}
}
