package separators

// Separator is by what means to separate data in the return values
type Separator int

func (s Separator) String() string {
	return Names[s]
}

const (
	// Dot is a full stop separator
	Dot Separator = iota
	// Comma is a comma separator
	Comma

	// Key is the API key name of the separator
	Key = "s"
)

// Separators represents all the valid separators
var Separators = map[Separator]struct{}{
	Dot:   struct{}{},
	Comma: struct{}{},
}

// Names are the API names of the valid separators
var Names = map[Separator]string{
	Dot:   "Dot",
	Comma: "Comma",
}
