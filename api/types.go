package api

import (
	"fmt"
	"time"
)

// Date is a specific calendar date
type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}
