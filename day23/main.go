package main

import (
	"fmt"

	multicoretablet "github.com/fivegreenapples/adventofcode2017/day18/multicore-tablet"
)

func main() {

	fmt.Printf("\nPart 1 (inspect 'mul' from the dumped stats):\n")
	part1Core := multicoretablet.MakeCore(nil, 0, input)
	part1Core.Run()

	fmt.Printf("\nPart 2 - value of register h is %d\n", doPart2(1))

}
