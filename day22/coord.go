package main

type coord struct {
	x int
	y int
}

func (c *coord) stepInDirection(d direction) {
	c.x += d.x
	c.y += d.y
}

type direction struct {
	x int
	y int
}

func (d *direction) turnLeft() {
	if (*d) == north {
		d.x = west.x
		d.y = west.y
	} else if (*d) == east {
		d.x = north.x
		d.y = north.y
	} else if (*d) == south {
		d.x = east.x
		d.y = east.y
	} else if (*d) == west {
		d.x = south.x
		d.y = south.y
	}
}
func (d *direction) turnRight() {
	if (*d) == north {
		d.x = east.x
		d.y = east.y
	} else if (*d) == east {
		d.x = south.x
		d.y = south.y
	} else if (*d) == south {
		d.x = west.x
		d.y = west.y
	} else if (*d) == west {
		d.x = north.x
		d.y = north.y
	}
}

var north = direction{0, -1}
var east = direction{1, 0}
var south = direction{0, 1}
var west = direction{-1, 0}
