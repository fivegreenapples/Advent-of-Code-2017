package main

import (
	"fmt"
)

func main() {
	testTuring := newTuring()
	for i := 1; i <= 6; i++ {
		testTuring.advanceTest()
		testTuring.printTape(-3, 6)
	}
	fmt.Printf("Test - checksum after 6 steps is %d\n", testTuring.checkSum())

	part1Turing := newTuring()
	for i := 1; i <= 12302209; i++ {
		part1Turing.advancePart1()
	}
	fmt.Printf("Checksum after 12302209 steps is %d\n", part1Turing.checkSum())
}
