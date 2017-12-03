package main

import (
	"fmt"
	"math"
	"os"
)

type coord struct {
	x int
	y int
}

func main() {
	fmt.Println("Test (1):", calcManhattan(1))
	fmt.Println("Test (12):", calcManhattan(12))
	fmt.Println("Test (23):", calcManhattan(23))
	fmt.Println("Test (1024):", calcManhattan(1024))
	fmt.Println("Answer1:", calcManhattan(input))

	allocations := map[coord]int{}
	allocations[coord{0, 0}] = 1

	c := 2
	for {
		cCoord := coordsForCell(c)
		adjCoords := adjacentCoordsForCoord(cCoord)
		sum := 0
		for _, aC := range adjCoords {
			sum += allocations[aC]
		}
		allocations[cCoord] = sum
		fmt.Println(c, sum)
		if sum > input {
			fmt.Println("First allocation larger than input is: ", sum)
			os.Exit(0)
		}
		c++
	}
}

func calcManhattan(input int) int {
	coords := coordsForCell(input)
	return int(math.Abs(float64(coords.x)) + math.Abs(float64(coords.y)))
}

func coordsForCell(c int) coord {
	// returns the x,y coords of a cell given its "id"

	if c == 1 {
		return coord{0, 0}
	}

	// If i is center distance, then cells exist from ((2i - 1)^2 + 1) to (2i + 1)^2
	// Find i assuming input is max in spiral ring. Then take ceil of that to get proper center distance
	// So cD == ceil((sqrt(INPUT) - 1) / 2)
	fInput := float64(c)
	centerDistance := math.Ceil((math.Sqrt(fInput) - 1.0) / 2.0)
	// With centerDistance find min and max for the ring of cells.
	minCell := math.Pow((2.0*centerDistance-1.0), 2) + 1.0
	// find side length, which is twice center distance
	sideLength := centerDistance * 2.0
	// find which side our input is on, the center of our side, and the distance from our side-center
	side := math.Floor((fInput - minCell) / sideLength)
	sideCenter := (minCell - 1.0) + ((side + 0.5) * sideLength)
	sideCenterDistance := fInput - sideCenter

	// fmt.Println(c, centerDistance, minCell, sideLength, side, sideCenter, sideCenterDistance)

	switch side {
	case 0.0:
		return coord{int(centerDistance), int(sideCenterDistance)}
	case 1.0:
		return coord{-int(sideCenterDistance), int(centerDistance)}
	case 2.0:
		return coord{-int(centerDistance), -int(sideCenterDistance)}
	case 3.0:
		return coord{int(sideCenterDistance), -int(centerDistance)}
	default:
		panic("unhandled side")
	}
}

func adjacentCoordsForCoord(c coord) []coord {
	// given a particular coordinate return the 8 adjacent coords
	// we go counter clockwise from cell immediately to right
	return []coord{
		coord{c.x + 1, c.y},
		coord{c.x + 1, c.y + 1},
		coord{c.x, c.y + 1},
		coord{c.x - 1, c.y + 1},
		coord{c.x - 1, c.y},
		coord{c.x - 1, c.y - 1},
		coord{c.x, c.y - 1},
		coord{c.x + 1, c.y - 1},
	}
}
