package main

import "fmt"

func main() {

	testInfectionsEvents, _ := makeGridAndInfect(testInput, 10000)
	fmt.Printf("Test part 1 - number bursts that caused an infection is %d\n", testInfectionsEvents)

	part1InfectionsEvents, _ := makeGridAndInfect(input, 10000)
	fmt.Printf("Part 1 - number bursts that caused an infection is %d\n", part1InfectionsEvents)

	testInfectionsEvents2, _ := makeGridAndInfect2(testInput, 10000000)
	fmt.Printf("Test part 2 - number bursts that caused an infection is %d\n", testInfectionsEvents2)

	part1InfectionsEvents2, _ := makeGridAndInfect2(input, 10000000)
	fmt.Printf("Part 2 - number bursts that caused an infection is %d\n", part1InfectionsEvents2)

}

func makeGridAndInfect(input []string, bursts int) (infectiousEvents, totalNowInfected int) {
	theGrid := makeGrid(input)
	testVirus := virus{
		grid:             &theGrid,
		current:          theGrid.center(),
		currentDirection: north,
	}
	rounds := bursts
	for rounds > 0 {
		testVirus.burst()
		rounds--
	}
	return testVirus.infectionCount, theGrid.countInfected()
}

func makeGridAndInfect2(input []string, bursts int) (infectiousEvents, totalNowInfected int) {
	theGrid := makeGrid(input)
	testVirus := virus{
		grid:             &theGrid,
		current:          theGrid.center(),
		currentDirection: north,
	}
	rounds := bursts
	for rounds > 0 {
		testVirus.burst2()
		rounds--
	}
	return testVirus.infectionCount, theGrid.countInfected()
}
