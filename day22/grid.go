package main

type nodestate int
type grid map[coord]nodestate

const (
	clean nodestate = iota
	weakened
	infected
	flagged
)

func makeGrid(in []string) grid {
	g := grid{}

	for y, row := range in {
		for x, r := range row {
			if r == '#' {
				g[coord{x, y}] = infected
			} else {
				g[coord{x, y}] = clean
			}
		}
	}

	return g
}

func (g *grid) countInfected() int {
	out := 0
	for _, cell := range *g {
		if cell == infected {
			out++
		}
	}
	return out
}
func (g *grid) center() coord {
	minX, minY, maxX, maxY := 10000000, 10000000, 0, 0
	for pos := range *g {
		if pos.x < minX {
			minX = pos.x
		}
		if pos.y < minY {
			minY = pos.y
		}
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}
	return coord{
		x: (minX + maxX) / 2,
		y: (minY + maxY) / 2,
	}
}

func (g *grid) isInfected(c coord) bool {
	return (*g)[c] == infected
}
func (g *grid) isClean(c coord) bool {
	return (*g)[c] == clean
}
func (g *grid) isWeakened(c coord) bool {
	return (*g)[c] == weakened
}
func (g *grid) isFlagged(c coord) bool {
	return (*g)[c] == flagged
}
func (g *grid) setState(c coord, s nodestate) {
	(*g)[c] = s
}
