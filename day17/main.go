package main

import (
	"fmt"

	"github.com/fivegreenapples/adventofcode2017/day17/buffer"
)

func main() {

	testBuffer := new(buffer.Circular)
	testZeroPos := testBuffer.Insert(0)
	insertValuesWithInterval(testBuffer, 9, testSteps, 0)
	fmt.Printf("Test buffer after 9 insertions: %v\n", testBuffer)
	fmt.Printf("Test of part 2: value after zero is: %d\n", testBuffer.MoveTo(testZeroPos).Move(1).Current())
	insertValuesWithInterval(testBuffer, 2017-9, testSteps, 9)
	fmt.Printf("Test of part 1: value after current is: %d\n", testBuffer.Move(1).Current())

	part1Buffer := new(buffer.Circular)
	/*zeroPos := */ part1Buffer.Insert(0)
	insertValuesWithInterval(part1Buffer, 2017, inputSteps, 0)
	fmt.Printf("Part1: value after current is: %d\n", part1Buffer.Move(1).Current())

	// This does work but takes over 5 minutes on my laptop (compared to 13ms for part1)
	// So we do part2 a different way. See below.
	//
	// insertValuesWithInterval(part1Buffer, 50000000-2017, inputSteps, 2017)
	// fmt.Printf("Part2: value after zero is: %d\n", part1Buffer.MoveTo(zeroPos).Move(1).Current())

	// Part2 the sneaky way
	// We only care about the value after zero. Which means we only care if an insert
	// lands on zero in which case we update our knowledge of what is after zero. For all
	// cycles we keep track of the length of the buffer and the current cursor position

	length := 1
	cursor := 0
	afterZero := 0
	stepsPerInsert := inputSteps
	for i := 1; i <= 50000000; i++ {
		cursor = (cursor + stepsPerInsert) % length
		if cursor == 0 {
			afterZero = i
		}
		cursor++
		length++
	}
	fmt.Printf("Part2: value after zero is: %d\n", afterZero)
}

func insertValuesWithInterval(b *buffer.Circular, numInserts, numSteps, valueOffset int) {
	startInsertions := 1 + valueOffset
	endInsertions := numInserts + valueOffset
	for insertions := startInsertions; insertions <= endInsertions; insertions++ {
		if insertions%10000 == 0 {
			fmt.Printf("\r%d insertions", insertions)
		}
		b.Move(numSteps).Insert(insertions)
	}
	// Clear progress line
	fmt.Print("\r                                \r")
}
