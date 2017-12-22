package main

type virus struct {
	infectionCount   int
	grid             *grid
	current          coord
	currentDirection direction
}

func (v *virus) burst() {
	if v.grid.isInfected(v.current) {
		v.currentDirection.turnRight()
		v.grid.setState(v.current, clean)
	} else {
		v.currentDirection.turnLeft()
		v.grid.setState(v.current, infected)
		v.infectionCount++
	}
	v.current.stepInDirection(v.currentDirection)
}

func (v *virus) burst2() {
	if v.grid.isClean(v.current) {
		v.currentDirection.turnLeft()
		v.grid.setState(v.current, weakened)
	} else if v.grid.isWeakened(v.current) {
		v.grid.setState(v.current, infected)
		v.infectionCount++
	} else if v.grid.isInfected(v.current) {
		v.currentDirection.turnRight()
		v.grid.setState(v.current, flagged)
	} else if v.grid.isFlagged(v.current) {
		v.currentDirection.turnRight()
		v.currentDirection.turnRight()
		v.grid.setState(v.current, clean)
	}
	v.current.stepInDirection(v.currentDirection)
}
