package main

import (
	"fmt"
)

type pathSegment rune

type maze map[coord]pathSegment

func (m *maze) getStart() coord {
	x := 0
	for {
		if ps, found := (*m)[coord{x, 0}]; found && ps == pathSegment('|') {
			return coord{x, 0}
		}
		x++
	}
}

func (m *maze) traverse() (string, int) {
	currentPos := m.getStart()
	prevDelta := south
	letters := ""
	steps := 0
	for {
		steps++
		segment := (*m)[currentPos]
		if segment == '+' {
			// must change direction
			if prevDelta == north || prevDelta == south {
				// try east/west
				_, foundEast := (*m)[currentPos.applyDelta(east)]
				_, foundWest := (*m)[currentPos.applyDelta(west)]
				if foundEast && foundWest {
					panic(fmt.Errorf("found multiple options at %v", currentPos))
				}
				if foundEast {
					currentPos = currentPos.applyDelta(east)
					prevDelta = east
				} else {
					currentPos = currentPos.applyDelta(west)
					prevDelta = west
				}
			} else {
				// try north/south
				_, foundNorth := (*m)[currentPos.applyDelta(north)]
				_, foundSouth := (*m)[currentPos.applyDelta(south)]
				if foundNorth && foundSouth {
					panic(fmt.Errorf("found multiple options at %v", currentPos))
				}
				if !foundNorth && !foundSouth {
					break
				}
				if foundNorth {
					currentPos = currentPos.applyDelta(north)
					prevDelta = north
				} else {
					currentPos = currentPos.applyDelta(south)
					prevDelta = south
				}
			}
		} else {
			// just keep going
			// if it's a letter keep track of it
			if segment >= 'A' && segment <= 'Z' {
				letters += string(segment)
			}
			currentPos = currentPos.applyDelta(prevDelta)
			_, foundNext := (*m)[currentPos]
			if !foundNext {
				break
			}
		}
	}
	return letters, steps
}
