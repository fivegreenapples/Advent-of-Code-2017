package main

import "fmt"

func main() {
	fmt.Println("day 19")
	testMaze := parse(testMap)
	testLetters, testSteps := testMaze.traverse()
	fmt.Printf("Test part 1 - the packet sees %s after %d steps\n", testLetters, testSteps)

	part1Maze := parse(inputMap)
	part1Letters, part1Steps := part1Maze.traverse()
	fmt.Printf("Part 1 - the packet sees %s after %d steps\n", part1Letters, part1Steps)
}

func parse(in []string) maze {
	m := maze{}
	for y, s := range in {
		for x, r := range s {
			if r == ' ' {
				continue
			}
			m[coord{x, y}] = pathSegment(r)
		}
	}
	return m
}

/*
	switch r {
	case "-":
	case "|":
	case "+":
	case "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
	case " ":
	default:
		panic(fmt.Errorf("unhandled map character: %c", r))
	}

*/
