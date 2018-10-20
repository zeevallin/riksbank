package calc

// Method is the method by which to compile numbers for a period
type Method int

func (m Method) String() string {
	return MethodName(m)
}

// Period is the time in which to compile entries
type Period int

func (p Period) String() string {
	return PeriodName(p)
}

const (
	// PeriodKey represents the API query key for calculation period
	PeriodKey = "f"
	// MethodKey represents the API query key for calculation method
	MethodKey = "c"

	// Ultimo takes the number for the last bank day of a period
	Ultimo Method = iota
	// Average takes the avarage of all numbers for a period
	Average
	// Min takes the lowest number for a period
	Min
	// Max takes the highest number for a period
	Max

	// Day is a period of one day
	Day Period = iota
	// Week is a period of one calendar week
	Week
	// Month is a period of one calendar month
	Month
	// Quarter is a period of three calendar months
	Quarter
	// Year is a period of one whole calendar year
	Year
)

// DayMethods are the allowed methods for the day period
var DayMethods = map[Method]struct{}{
	Average: struct{}{},
}

// WeekMethods are the allowed methods for the week period
var WeekMethods = map[Method]struct{}{
	Min:     struct{}{},
	Max:     struct{}{},
	Average: struct{}{},
}

// MonthMethods are the allowed methods for the month period
var MonthMethods = map[Method]struct{}{
	Min:     struct{}{},
	Max:     struct{}{},
	Average: struct{}{},
	Ultimo:  struct{}{},
}

// QuarterMethods are the allowed methods for the quarter period
var QuarterMethods = map[Method]struct{}{
	Min:     struct{}{},
	Max:     struct{}{},
	Average: struct{}{},
}

// YearMethods are the allowed methods for the year period
var YearMethods = map[Method]struct{}{
	Min:     struct{}{},
	Max:     struct{}{},
	Average: struct{}{},
}

// PeriodMethods are the allowed methods for all given periods
var PeriodMethods = map[Period]map[Method]struct{}{
	Day:     DayMethods,
	Week:    WeekMethods,
	Month:   MonthMethods,
	Quarter: QuarterMethods,
	Year:    YearMethods,
}

// MethodNames are the API names for the methods
var MethodNames = map[Method]string{
	Ultimo:  "cUltimo",
	Average: "cAverage",
	Min:     "cMin",
	Max:     "cMax",
}

// MethodName returns the API name corresponding to the method
func MethodName(m Method) string {
	return MethodNames[m]
}

// PeriodNames are the API names of the periods
var PeriodNames = map[Period]string{
	Day:     "Day",
	Week:    "Week",
	Month:   "Month",
	Quarter: "Quarter",
	Year:    "Year",
}

// PeriodName returns the API name corresponding to the period
func PeriodName(p Period) string {
	return PeriodNames[p]
}
