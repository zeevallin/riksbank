package swea

import (
	"strings"
)

func isTrue(s string) bool {
	switch strings.ToUpper(s) {
	case "Y", "YES", "TRUE", "1":
		return true
	}
	return false
}
