package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	loopsTest, sizeTest := calc(testInput)
	fmt.Println("Test1:", loopsTest, sizeTest)
	loops, size := calc(input)
	fmt.Println("Answer:", loops, size)
}

func calc(input []int) (int, int) {

	memory := make([]int, len(input))
	copy(memory, input)
	memorySize := len(memory)
	knownArrangements := map[string]int{}
	loops := 0

	for {
		currentArrangement := joinInts(memory, ",")
		if loopNum, found := knownArrangements[currentArrangement]; found {
			return loops, loops - loopNum
		}
		knownArrangements[currentArrangement] = loops
		loops++

		largestBank := 0
		for bank, blocks := range memory {
			if blocks > memory[largestBank] {
				largestBank = bank
			}
		}

		// Get number of blocks to redistribute
		redistBlocks := memory[largestBank]
		// zero the bank
		memory[largestBank] = 0
		// distribute the blocks over the memory
		for j := 1; j <= redistBlocks; j++ {
			memory[(largestBank+j)%memorySize]++
		}
	}

}

func joinInts(in []int, sep string) string {
	intermediate := make([]string, len(in))
	for i, v := range in {
		intermediate[i] = strconv.Itoa(v)
	}
	return strings.Join(intermediate, sep)
}
