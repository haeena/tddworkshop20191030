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

func (c1 *Closedrange) Equal(c2 *Closedrange) bool {
	if c2 == nil {
		return false
	}
	return c1.lower == c2.lower && c1.upper == c2.upper
}

func (c1 *Closedrange) IncludeClosedrange(c2 *Closedrange) bool {
	if c1.Equal(c2) {
		return true
	}
	if c2.lower >= c1.lower && c2.upper <= c1.upper {
		return true
	}
	return false
}
