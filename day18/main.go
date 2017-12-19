package main

import (
	"fmt"

	multicoretablet "github.com/fivegreenapples/adventofcode2017/day18/multicore-tablet"
	"github.com/fivegreenapples/adventofcode2017/day18/tablet"
)

func main() {
	testTablet := tablet.Make(testInput)
	testTablet.Reset()
	lastSound := testTablet.Run()
	fmt.Printf("Test part 1 - last sound played is %d\n", lastSound)

	part1Tablet := tablet.Make(input)
	part1Tablet.Reset()
	part1LastSound := part1Tablet.Run()
	fmt.Printf("Part 1 - last sound played is %d\n", part1LastSound)

	fmt.Printf("\nPart 2 test:\n")
	testPart2Tablet := multicoretablet.New(testInputPart2)
	testPart2Tablet.Run()

	fmt.Printf("\nPart 2 actual:\n")
	part2Tablet := multicoretablet.New(input)
	part2Tablet.Run()
}
