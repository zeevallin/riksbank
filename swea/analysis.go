package swea

import (
	"fmt"
	"strings"
)

// AnalysisMethod represents the analysis method for comparing values in a period
type AnalysisMethod int

func (am AnalysisMethod) String() string {
	if name, ok := AnalysisMethodNames[am]; ok {
		return name
	}
	return ""
}

const (
	// Real is the actual value for the period
	Real AnalysisMethod = iota + 1
	// Mean is the average value of all values in a period
	Mean
	// Min is the smallest value of all values in a period
	Min
	// Max is the largest value of all values in a period
	Max
	// Ultimo is the value of the last banking day in the period
	Ultimo
)

var (
	// AnalysisMethods are all known analysis methods
	AnalysisMethods = map[string]AnalysisMethod{
		"real":   Real,
		"mean":   Mean,
		"min":    Min,
		"max":    Max,
		"ultimo": Ultimo,
	}

	// AnalysisMethodNames are all known analysis method names
	AnalysisMethodNames = map[AnalysisMethod]string{
		Real:   "real",
		Mean:   "mean",
		Min:    "min",
		Max:    "max",
		Ultimo: "ultimo",
	}

	// DailyAnalysis are the analysis methods allowed for daily aggregation
	DailyAnalysis = map[AnalysisMethod]struct{}{
		Real: struct{}{},
	}

	// WeeklyAnalysis are the analysis methods allowed for weekly aggregation
	WeeklyAnalysis = map[AnalysisMethod]struct{}{
		Mean: struct{}{},
		Min:  struct{}{},
		Max:  struct{}{},
	}

	// MonthlyAnalysis are the analysis methods allowed for monthly aggregation
	MonthlyAnalysis = map[AnalysisMethod]struct{}{
		Mean: struct{}{},
		Min:  struct{}{},
		Max:  struct{}{},
	}

	// QuarterlyAnalysis are the analysis methods allowed for quarterly aggregation
	QuarterlyAnalysis = map[AnalysisMethod]struct{}{
		Mean: struct{}{},
		Min:  struct{}{},
		Max:  struct{}{},
	}

	// YearlyAnalysis are the analysis methods allowed for yearly aggregation
	YearlyAnalysis = map[AnalysisMethod]struct{}{
		Mean:   struct{}{},
		Min:    struct{}{},
		Max:    struct{}{},
		Ultimo: struct{}{},
	}
)

// ErrUnknownAnalysisForAggregate happens when the analysis method cannot be used for the aggregate method
type ErrUnknownAnalysisForAggregate struct {
	anm AnalysisMethod
	agm AggregateMethod
}

func (e ErrUnknownAnalysisForAggregate) Error() string {
	return fmt.Sprintf("%s aggregate does not support analysis method: %v", AggregateName(e.agm), e.anm)
}

// ErrUnknownAnalysis happens when the analysis method does not exist at all
type ErrUnknownAnalysis struct {
	s string
}

func (e ErrUnknownAnalysis) Error() string {
	return fmt.Sprintf("analysis method does not exist: %s", e.s)
}

// ParseAnalysisForAggregate will attempt to parse a string with an aggregate
func ParseAnalysisForAggregate(s string, aggregate AggregateMethod) (AnalysisMethod, error) {
	analysis, ok := AnalysisMethods[strings.ToLower(s)]
	if !ok {
		return AnalysisMethod(0), ErrUnknownAnalysis{s}
	}
	err := ErrUnknownAnalysisForAggregate{analysis, aggregate}
	switch aggregate {
	case Daily:
		if _, ok := DailyAnalysis[analysis]; !ok {
			return analysis, err
		}
	case Weekly:
		if _, ok := WeeklyAnalysis[analysis]; !ok {
			return analysis, err
		}
	case Monthly:
		if _, ok := MonthlyAnalysis[analysis]; !ok {
			return analysis, err
		}
	case Quarterly:
		if _, ok := QuarterlyAnalysis[analysis]; !ok {
			return analysis, err
		}
	case Yearly:
		if _, ok := YearlyAnalysis[analysis]; !ok {
			return analysis, err
		}
	default:
		return analysis, err
	}
	return analysis, nil
}
