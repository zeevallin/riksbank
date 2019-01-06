package cli

import (
	"fmt"
	"time"
)

const (
	dateLayout = "2006-01-02"
)

func parseDate(s string) (time.Time, error) {
	return time.Parse(dateLayout, s)
}

func formatDate(t time.Time) string {
	return t.Format(dateLayout)
}

func formatFloat(f *float64) string {
	if f == nil {
		return "-"
	}
	return fmt.Sprintf("%f", *f)
}

func formatInt(i int) string {
	return fmt.Sprintf("%d", i)
}

func boolToYesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
