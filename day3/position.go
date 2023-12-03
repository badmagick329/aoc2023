package day3

type Pos struct {
	x int
	y int
}

func (p *Pos) adjacentPositions(width, height int) []Pos {
	x := p.x
	y := p.y
	adj := []Pos{}
	// anti clockwise starting at up
	if y > 0 {
		adj = append(adj, Pos{x, y - 1})
	}

	if x > 0 {
		if y > 0 {
			adj = append(adj, Pos{x - 1, y - 1})
		}
		adj = append(adj, Pos{x - 1, y})
	}

	if y < height-1 {
		if x > 0 {
			adj = append(adj, Pos{x - 1, y + 1})
		}
		adj = append(adj, Pos{x, y + 1})
	}

	if x < width-1 {
		if y < height-1 {
			adj = append(adj, Pos{x + 1, y + 1})
		}
		adj = append(adj, Pos{x + 1, y})
		if y > 0 {
			adj = append(adj, Pos{x + 1, y - 1})
		}
	}
	return adj
}
