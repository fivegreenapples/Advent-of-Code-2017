package main

import "fmt"

func main() {
	fmt.Println("Test (1122):", calc([]byte("1122")))
	fmt.Println("Test (1111):", calc([]byte("1111")))
	fmt.Println("Test (1234):", calc([]byte("1234")))
	fmt.Println("Test (91212129):", calc([]byte("91212129")))
	fmt.Println("Answer1:", calc([]byte(input)))

	fmt.Println("Test (1212):", calc2([]byte("1212")))
	fmt.Println("Test (1221):", calc2([]byte("1221")))
	fmt.Println("Test (123425):", calc2([]byte("123425")))
	fmt.Println("Test (123123):", calc2([]byte("123123")))
	fmt.Println("Test (12131415):", calc2([]byte("12131415")))
	fmt.Println("Answer2:", calc2([]byte(input)))

}

func calc(input []byte) int {

	if len(input) <= 1 {
		return 0
	}

	var sum = 0
	var prev = input[len(input)-1]
	for _, b := range input {
		if prev == b {
			sum += int(b - 48)
		}
		prev = b
	}
	return sum
}

func calc2(input []byte) int {

	if len(input) <= 1 {
		return 0
	}
	if len(input)%2 != 0 {
		return 0
	}

	var diff = len(input) / 2

	var sum = 0
	for i, b := range input[0:diff] {
		var cmp = input[i+diff]
		if b == cmp {
			sum += int(b - 48)
		}
	}
	return sum * 2
}
