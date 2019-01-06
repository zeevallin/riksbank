package swea

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	pmi = "PMI"
	sek = "SEK"
)

var (
	pcpRegex = regexp.MustCompile(`[/\-:]`)
)

// Currency is a ISO standard name for a currency
type Currency string

// ParseCurrency will attempt to turn a string into a currency
func ParseCurrency(s string) Currency {
	if strings.HasSuffix(s, pmi) {
		return Currency(s[3:6])
	}
	return Currency(s)
}

// ToSeries converts the currency to a series
func (c Currency) ToSeries() Series {
	if c == sek {
		return Series(c)
	}
	return Series(fmt.Sprintf("%s%s%s", sek, c, pmi))
}

// CurrencyPair represents a currency conversion
type CurrencyPair struct {
	Base    Currency
	Counter Currency
}

// ParseCurrencyPair will attempt to turn a string into a currency pair
func ParseCurrencyPair(s string) CurrencyPair {
	split := pcpRegex.Split(s, -1)
	if len(split) > 1 {
		return CurrencyPair{
			Base:    ParseCurrency(split[0]),
			Counter: ParseCurrency(split[1]),
		}
	}
	return CurrencyPair{
		Base:    ParseCurrency(split[0]),
		Counter: ParseCurrency(split[0]),
	}
}

func (cp CurrencyPair) String() string {
	return fmt.Sprintf("%s/%s", cp.Base, cp.Counter)
}

// ToCrossPair converts a currency pair to a cross pair
func (cp CurrencyPair) ToCrossPair() CrossPair {
	return CrossPair{
		Base:    cp.Base.ToSeries(),
		Counter: cp.Counter.ToSeries(),
	}
}
