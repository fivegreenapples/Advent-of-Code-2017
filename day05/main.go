package main

import "fmt"

func main() {
	testInput1 := make([]int, len(testInput))
	input1 := make([]int, len(input))
	copy(testInput1, testInput)
	copy(input1, input)

	fmt.Println("Test:", calc(testInput1))
	fmt.Println("Answer:", calc(input1))
	fmt.Println("Test2:", calc2(testInput))
	fmt.Println("Answer2:", calc2(input))
}

func calc(input []int) int {
	len := len(input)
	pointer := 0
	steps := 0
	for pointer < len {
		steps++
		jump := input[pointer]
		input[pointer]++
		pointer += jump
	}
	return steps
}
func calc2(input []int) int {
	len := len(input)
	pointer := 0
	steps := 0
	for pointer < len {
		steps++
		jump := input[pointer]
		if jump >= 3 {
			input[pointer]--
		} else {
			input[pointer]++
		}
		pointer += jump
	}
	return steps
}
