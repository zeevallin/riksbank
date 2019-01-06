package xmlstrings

import (
	"strings"
)

// ParseBool returns the boolean value represented by the string
func ParseBool(s string) bool {
	switch strings.TrimSpace(strings.ToUpper(s)) {
	case "Y", "YES", "TRUE", "1":
		return true
	}
	return false
}
