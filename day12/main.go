package main

import "fmt"

func main() {
	testGroup0 := findGroupForID(0, testInput)
	fmt.Println(testGroup0, len(testGroup0))

	part1Group0 := findGroupForID(0, input)
	fmt.Println(part1Group0, len(part1Group0))

	// part2 ...
	allGroups := []programGroup{}
	allProgramsSeen := programGroup{}

	for anID := range input {
		if _, seen := allProgramsSeen[anID]; seen {
			continue
		}
		thisGroup := findGroupForID(anID, input)
		allGroups = append(allGroups, thisGroup)
		for i := range thisGroup {
			allProgramsSeen[i] = struct{}{}
		}
	}

	fmt.Println("Total groups is:", len(allGroups))

}

type programGroup map[int]struct{}

func findGroupForID(id int, pipes map[int][]int) programGroup {
	group := programGroup{}
	toProcess := map[int]struct{}{
		id: struct{}{},
	}
	current := 0
	for len(toProcess) > 0 {
		for current = range toProcess {
			break
		}
		delete(toProcess, current)
		group[current] = struct{}{}

		connections := pipes[current]
		for _, c := range connections {
			if _, done := group[c]; done {
				continue
			}

			toProcess[c] = struct{}{}
		}
	}

	return group
}
