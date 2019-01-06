package cli

import "time"

const (
	dateLayout = "2006-01-02"
)

func parseDate(s string) (time.Time, error) {
	return time.Parse(dateLayout, s)
}

func formatDate(t time.Time) string {
	return t.Format(dateLayout)
}

func boolToYesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
