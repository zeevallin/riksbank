package period

// Period represets a period between two dates
type Period struct {
	name string
}

// Parse will attempt to turn a string into a period
func Parse(s string) Period {
	return Period{
		name: s,
	}
}

func (p Period) String() string {
	return p.name
}
