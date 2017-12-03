package main

import "fmt"

func main() {
	fmt.Println("Test: ", calcChecksum(testInput))
	fmt.Println("Answer1: ", calcChecksum(input))
	fmt.Println("Test2: ", calcChecksum2(testInput2))
	fmt.Println("Answer2: ", calcChecksum2(input))
}

func calcChecksum(input [][]uint) uint {
	var checkSum uint
	for _, row := range input {
		if len(row) <= 1 {
			continue
		}
		var max uint
		var min = ^uint(0)
		for _, cell := range row {
			if cell > max {
				max = cell
			}
			if cell < min {
				min = cell
			}
		}
		checkSum += max - min
	}
	return checkSum
}

func calcChecksum2(input [][]uint) uint {

	var checkSum uint
	for _, row := range input {
		if len(row) <= 1 {
			continue
		}

		for iIndex, i := range row[:len(row)-1] {
			for _, j := range row[iIndex+1:] {
				if isDivisible, val := isEvenlyDivisible(i, j); isDivisible {
					checkSum += val
					goto next
				}
			}
		}
	next:
	}
	return checkSum
}

func isEvenlyDivisible(a, b uint) (bool, uint) {
	if a > b {
		if a%b == 0 {
			return true, a / b
		}
		return false, 0
	}
	if b%a == 0 {
		return true, b / a
	}
	return false, 0
}
