package main

import (
	"fmt"
	"math"
)

func main() {

	for i, test := range tests {
		pos := hexCoord{0, 0, 0}
		for _, s := range test {
			pos.addDelta(deltaFromStep(s))
		}
		fmt.Println("Test", i+1, pos, pos.toCartesian().toHexagonal(), pos.toCartesian().toHexagonal().manhattanDistance())
	}

	pos := hexCoord{0, 0, 0}
	maxManhattan := 0
	for _, s := range input {
		pos.addDelta(deltaFromStep(s))
		pos = pos.toCartesian().toHexagonal()
		if pos.manhattanDistance() > maxManhattan {
			maxManhattan = pos.manhattanDistance()
		}
	}
	fmt.Println(pos, pos.manhattanDistance(), maxManhattan)

}

type hexCoord struct {
	n, ne, se int
}
type cartCoord struct {
	x, y int
}

func (c *hexCoord) addDelta(d delta) {
	c.n += d.n
	c.ne += d.ne
	c.se += d.se
}

func (c hexCoord) toCartesian() cartCoord {
	return cartCoord{
		x: c.ne + c.se,
		y: (2 * c.n) + c.ne - c.se,
	}
}
func (c hexCoord) manhattanDistance() int {
	return int(math.Abs(float64(c.n)) + math.Abs(float64(c.ne)) + math.Abs(float64(c.se)))
}

func (c cartCoord) toHexagonal() hexCoord {
	absX, absY := c.x, c.y
	if absX < 0 {
		absX = -absX
	}
	if absY < 0 {
		absY = -absY
	}
	if absX > absY {
		// hex coord involves only ne/se components
		// x=ne+se; y=ne-se
		// ne=(x+y)/2; se=(x-y)/2
		return hexCoord{
			n:  0,
			ne: (c.x + c.y) / 2,
			se: (c.x - c.y) / 2,
		}
	}

	// hex coord involves n component (possibly zero)
	if c.x < 0 && c.y < 0 || c.x > 0 && c.y > 0 {
		// hex coord involves ne component (possibly zero)
		// x=ne; y=2n+ne
		// ne=x; n=(y-x)/2
		return hexCoord{
			n:  (c.y - c.x) / 2,
			ne: c.x,
			se: 0,
		}
	}

	// hex coord involves se component (possibly zero)
	// x=se; y=2n-se
	// se=x; n=(y+x)2
	return hexCoord{
		n:  (c.y + c.x) / 2,
		ne: 0,
		se: c.x,
	}
}

type delta struct {
	n, ne, se int
}

func deltaFromStep(st step) delta {
	switch st {
	case n:
		return delta{1, 0, 0}
	case ne:
		return delta{0, 1, 0}
	case se:
		return delta{0, 0, 1}
	case s:
		return delta{-1, 0, 0}
	case sw:
		return delta{0, -1, 0}
	case nw:
		return delta{0, 0, -1}
	default:
		panic("unsupported step")
	}
}
