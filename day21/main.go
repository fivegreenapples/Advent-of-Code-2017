package main

import "fmt"

func main() {

	expandedTestRules := expandRules(testInput)
	currentTest := makeGrid(startingGrid)
	currentTest = iterate(currentTest, expandedTestRules)
	currentTest = iterate(currentTest, expandedTestRules)
	fmt.Printf("Test part 1 - result after 2 iterations is:\n%s\n", currentTest)
	fmt.Printf("Test part 1 - number of pixels on is %d\n", currentTest.numPixelsOn())

	expandedPart1Rules := expandRules(input)
	currentPart1 := makeGrid(startingGrid)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	fmt.Printf("Part 1 - result after 5 iterations is:\n%s\n", currentPart1)
	fmt.Printf("Part 1 - number of pixels on is %d\n", currentPart1.numPixelsOn())
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	currentPart1 = iterate(currentPart1, expandedPart1Rules)
	fmt.Printf("Part 2 - number of pixels on is %d\n", currentPart1.numPixelsOn())
}

func iterate(pg pixelGrid, rules map[string]string) pixelGrid {
	var subGridLength int
	if len(pg)%2 == 0 {
		subGridLength = 2
	} else {
		subGridLength = 3
	}
	numSubGrids := len(pg) / subGridLength
	allSubgrids := []pixelGrid{}
	for y := 0; y < numSubGrids; y++ {
		for x := 0; x < numSubGrids; x++ {
			subGrid := pg.getSubGrid(x, y, subGridLength)
			result, foundMatch := rules[subGrid.key()]
			if !foundMatch {
				panic("eek no match for subgrid: " + subGrid.key())
			}
			subGrid = makeGrid(result)
			allSubgrids = append(allSubgrids, subGrid)
		}
	}
	return joinGrids(allSubgrids, numSubGrids*(subGridLength+1))
}
func expandRules(in map[string]string) map[string]string {
	out := map[string]string{}
	for k, v := range in {
		grid := makeGrid(k)
		gridOptions := grid.makeKeys()
		for _, s := range gridOptions {
			out[s] = v
		}
	}
	return out
}
