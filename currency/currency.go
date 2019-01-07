package currency

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

// Parse will attempt to turn a string into a currency
func Parse(s string) Currency {
	s = strings.ToUpper(s)
	if strings.HasSuffix(s, pmi) {
		return Currency(s[3:6])
	}
	return Currency(s)
}

// Series returns the currency series name
func (c Currency) Series() string {
	if c == sek {
		return string(c)
	}
	return fmt.Sprintf("%s%s%s", sek, c, pmi)
}

// Pair represents a currency conversion
type Pair struct {
	Base    Currency
	Counter Currency
}

// ParsePair will attempt to turn a string into a currency pair
func ParsePair(s string) Pair {
	split := pcpRegex.Split(s, -1)
	if len(split) > 1 {
		return Pair{
			Base:    Parse(split[0]),
			Counter: Parse(split[1]),
		}
	}
	return Pair{
		Base:    Parse(split[0]),
		Counter: Parse(split[0]),
	}
}

func (cp Pair) String() string {
	return fmt.Sprintf("%s/%s", cp.Base, cp.Counter)
}
