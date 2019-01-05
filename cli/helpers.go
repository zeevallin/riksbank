package cli

func boolToYesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
