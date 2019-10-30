package closedrange

type Closedrange struct {
	lower int
	upper int
}

func NewClosedrange(lower int, upper int) *Closedrange {
	c := &Closedrange{lower, upper}
	return c
}
