package main

type coord struct {
	x int
	y int
}

func (c coord) applyDelta(d delta) coord {
	return coord{
		c.x + d.x,
		c.y + d.y,
	}
}

type delta struct {
	x int
	y int
}

var north = delta{0, -1}
var east = delta{1, 0}
var south = delta{0, 1}
var west = delta{-1, 0}
