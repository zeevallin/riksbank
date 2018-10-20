package api

import (
	"fmt"

	"github.com/zeeraw/riksbank/api/calc"
	"github.com/zeeraw/riksbank/api/separators"
)

// ValidatePeriodMethods looks at the period and methods to see if they're a valid combination
func ValidatePeriodMethods(period calc.Period, methods ...calc.Method) error {
	if cm, ok := calc.PeriodMethods[period]; ok {
		for _, method := range methods {
			if _, ok := cm[method]; !ok {
				return fmt.Errorf("%s is not a valid compile method for the %s period", method, period)
			}
		}
	} else {
		return fmt.Errorf("%s is not a valid period", period)
	}
	return nil
}

// ValidateSeparator takes a string and checks if it is a valid separator
func ValidateSeparator(sep separators.Separator) error {
	if _, ok := separators.Separators[sep]; !ok {
		return fmt.Errorf("%s is not a valid separator", sep)
	}
	return nil
}
