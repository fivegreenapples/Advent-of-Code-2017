package main

import "fmt"

func main() {

	testAllComponents := convertToMap(testInput)
	testPart1Strongest, _ := findStrongest(0, testAllComponents, false)
	testPart2Strongest, _ := findStrongest(0, testAllComponents, true)
	fmt.Printf("Test part 1 - strongest is %d\n", testPart1Strongest)
	fmt.Printf("Test part 2 - strongest is %d\n", testPart2Strongest)

	allComponents := convertToMap(input)
	part1Strongest, _ := findStrongest(0, allComponents, false)
	part2Strongest, _ := findStrongest(0, allComponents, true)
	fmt.Printf("Part 1 - strongest is %d\n", part1Strongest)
	fmt.Printf("Part 2 - strongest is %d\n", part2Strongest)
}

func findStrongest(startingPortType int, remainingComponents map[int]map[int]int, prioritiseLength bool) (strengthOfStrongest, lengthOfStrongest int) {

	strengthOfStrongest = 0
	lengthOfStrongest = 0

	// loop over components with the given starting port type
	for k, count := range remainingComponents[startingPortType] {
		// ignore if we don't have any of this component remaining
		if count == 0 {
			continue
		}

		// decrement counts for this component (once for each end) to
		// signal we just used the component
		remainingComponents[startingPortType][k]--
		remainingComponents[k][startingPortType]--

		thisStrength := k + startingPortType
		thisLength := 1

		// recurse to find strongest from the other end of our component
		subStrength, subLength := findStrongest(k, remainingComponents, prioritiseLength)

		thisStrength += subStrength
		thisLength += subLength

		// decide whether to bump our current leaders
		if prioritiseLength {
			if thisLength > lengthOfStrongest {
				lengthOfStrongest = thisLength
				strengthOfStrongest = thisStrength
			} else if thisLength == lengthOfStrongest {
				if thisStrength > strengthOfStrongest {
					strengthOfStrongest = thisStrength
				}
			}
		} else if thisStrength > strengthOfStrongest {
			lengthOfStrongest = thisLength
			strengthOfStrongest = thisStrength
		}

		// increment counts (put the component back on the heap)
		remainingComponents[startingPortType][k]++
		remainingComponents[k][startingPortType]++
	}

	return
}

func convertToMap(in []component) map[int]map[int]int {

	out := map[int]map[int]int{}
	for _, c := range in {
		if _, found := out[c[0]]; !found {
			out[c[0]] = map[int]int{}
		}
		if _, found := out[c[1]]; !found {
			out[c[1]] = map[int]int{}
		}
		out[c[0]][c[1]]++
		out[c[1]][c[0]]++
	}
	return out
}
