package swea

import (
	"fmt"
	"strings"
)

// AggregateMethod represents an enumeration of available aggregate methods when calculating values
type AggregateMethod string

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

// ErrUnknownAggregateMethod happens when a string is not a known method
type ErrUnknownAggregateMethod struct {
	am AggregateMethod
}

func (e ErrUnknownAggregateMethod) Error() string {
	return fmt.Sprintf("cannot parse aggregate method: %s", e.am)
}

// AggregateNames are the real names of all known aggregate methods
var AggregateNames = map[AggregateMethod]string{
	Daily:     "daily",
	Weekly:    "weekly",
	Monthly:   "monthly",
	Quarterly: "quarterly",
	Yearly:    "yearly",
}

// AggregateMethods are all known aggregate methods
var AggregateMethods = map[string]AggregateMethod{
	"daily":     Daily,
	"weekly":    Weekly,
	"monthly":   Monthly,
	"quarterly": Quarterly,
	"yearly":    Yearly,
}

// ParseAggregate will turn a string into an aggregate method
func ParseAggregate(s string) (AggregateMethod, error) {
	if method, ok := AggregateMethods[strings.ToLower(s)]; ok {
		return method, nil
	}
	am := AggregateMethod(s)
	return am, ErrUnknownAggregateMethod{am}
}

// AggregateName returns the name for an aggregate method
func AggregateName(a AggregateMethod) string {
	if name, ok := AggregateNames[a]; ok {
		return name
	}
	return string(a)
}
