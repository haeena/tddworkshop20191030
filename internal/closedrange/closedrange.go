package closedrange

type Closedrange struct {
	lower int
	upper int
}

func NewClosedrange(lower int, upper int) *Closedrange {
	if lower > upper {
		return nil
	}
	c := &Closedrange{lower, upper}
	return c
}

func (c *Closedrange) Include(number int) bool {
	return c.lower <= number && number <= c.upper
}
