package day2

type Colors struct {
	r int
	g int
	b int
}

func (c *Colors) contains(other Colors) bool {
	return other.r <= c.r && other.g <= c.g && other.b <= c.b
}

func (c *Colors) power() int {
	return c.r * c.g * c.b
}

func minRequiredFor(played []Colors) Colors {
	minColors := Colors{}
	for _, c := range played {
		if minColors.r < c.r {
			minColors.r = c.r
		}
		if minColors.g < c.g {
			minColors.g = c.g
		}
		if minColors.b < c.b {
			minColors.b = c.b
		}
	}
	return minColors
}
