package riksbank

import (
	"fmt"
	"strings"
)

const (
	// Daily aggregation method (default)
	Daily = AggregateMethod("D")

	// Weekly aggregation method
	Weekly = AggregateMethod("W")

	// Monthly aggregation method
	Monthly = AggregateMethod("M")

	// Quarterly aggregation method
	Quarterly = AggregateMethod("Q")

	// Yearly aggregation method
	Yearly = AggregateMethod("Y")
)

var (
	// AggregateNames are the real names of all known aggregate methods
	AggregateNames = map[AggregateMethod]string{
		Daily:     "daily",
		Weekly:    "weekly",
		Monthly:   "monthly",
		Quarterly: "quarterly",
		Yearly:    "yearly",
	}

	// AggregateMethods are all known aggregate methods
	AggregateMethods = map[string]AggregateMethod{
		"daily":     Daily,
		"weekly":    Weekly,
		"monthly":   Monthly,
		"quarterly": Quarterly,
		"yearly":    Yearly,
	}
)

// ParseAggregate will turn a string into an aggregate method
func ParseAggregate(s string) (AggregateMethod, error) {
	if method, ok := AggregateMethods[strings.ToLower(strings.TrimSpace(s))]; ok {
		return method, nil
	}
	method := AggregateMethod(s)
	return method, UnknownAggregateMethodError{method}
}

// UnknownAggregateMethodError happens when a string is not a known method
type UnknownAggregateMethodError struct {
	am AggregateMethod
}

func (e UnknownAggregateMethodError) Error() string {
	return fmt.Sprintf("cannot parse aggregate method: %s", e.am)
}

// AggregateMethod represents an enumeration of available aggregate methods when calculating values
type AggregateMethod string

func (am AggregateMethod) String() string {
	return string(am)
}

// Name returns the human readable name of the aggregate method
func (am AggregateMethod) Name() string {
	if name, ok := AggregateNames[am]; ok {
		return name
	}
	return string(am)
}
