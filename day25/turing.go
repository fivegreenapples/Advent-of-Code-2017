package main

import (
	"fmt"
)

type machinestate int

const (
	A machinestate = iota
	B
	C
	D
	E
	F
)

type turingmachine struct {
	tape   map[int]bool
	cursor int
	state  machinestate
}

func newTuring() *turingmachine {
	return &turingmachine{
		tape:   map[int]bool{},
		cursor: 0,
		state:  A,
	}
}
func (t *turingmachine) advanceTest() {
	// 	In state A:
	if t.state == A {
		// 	If the current value is 0:
		if !t.tape[t.cursor] {
			// 	  - Write the value 1.
			t.tape[t.cursor] = true
			// 	  - Move one slot to the right.
			t.cursor++
			// 	  - Continue with state B.
			t.state = B
		} else { // 	If the current value is 1:
			// 	  - Write the value 0.
			t.tape[t.cursor] = false
			// 	  - Move one slot to the left.
			t.cursor--
			// 	  - Continue with state B.
			t.state = B
		}
	} else { //   In state B:
		// 	If the current value is 0:
		if !t.tape[t.cursor] {
			// 	  - Write the value 1.
			t.tape[t.cursor] = true
			// 	  - Move one slot to the left.
			t.cursor--
			// 	  - Continue with state A.
			t.state = A
		} else { // 	If the current value is 1:
			// 	  - Write the value 1.
			t.tape[t.cursor] = true
			// 	  - Move one slot to the right.
			t.cursor++
			// 	  - Continue with state A.
			t.state = A
		}
	}
}
func (t *turingmachine) advancePart1() {
	// 	In state A:
	if t.state == A {
		// 	If the current value is 0:
		if !t.tape[t.cursor] {
			// 	  - Write the value 1.
			t.tape[t.cursor] = true
			// 	  - Move one slot to the right.
			t.cursor++
			// 	  - Continue with state B.
			t.state = B
		} else { // 	If the current value is 1:
			// 	  - Write the value 0.
			t.tape[t.cursor] = false
			// 	  - Move one slot to the left.
			t.cursor--
			// 	  - Continue with state D.
			t.state = D
		}
	} else if t.state == B { // In state B:
		// 	If the current value is 0:
		if !t.tape[t.cursor] {
			// 	  - Write the value 1.
			t.tape[t.cursor] = true
			// 	  - Move one slot to the right.
			t.cursor++
			// 	  - Continue with state C.
			t.state = C
		} else { // 	If the current value is 1:
			// 	  - Write the value 0.
			t.tape[t.cursor] = false
			// 	  - Move one slot to the right.
			t.cursor++
			// 	  - Continue with state F.
			t.state = F
		}
	} else if t.state == C { // In state C:
		// 	If the current value is 0:
		if !t.tape[t.cursor] {
			// 	  - Write the value 1.
			t.tape[t.cursor] = true
			// 	  - Move one slot to the left.
			t.cursor--
			// 	  - Continue with state C.
			t.state = C
		} else { // 	If the current value is 1:
			// 	  - Write the value 1.
			t.tape[t.cursor] = true
			// 	  - Move one slot to the left.
			t.cursor--
			// 	  - Continue with state A.
			t.state = A
		}
	} else if t.state == D { // In state D:
		// 	If the current value is 0:
		if !t.tape[t.cursor] {
			// 	  - Write the value 0.
			t.tape[t.cursor] = false
			// 	  - Move one slot to the left.
			t.cursor--
			// 	  - Continue with state E.
			t.state = E
		} else { // 	If the current value is 1:
			// 	  - Write the value 1.
			t.tape[t.cursor] = true
			// 	  - Move one slot to the right.
			t.cursor++
			// 	  - Continue with state A.
			t.state = A
		}
	} else if t.state == E { // In state E:
		// 	If the current value is 0:
		if !t.tape[t.cursor] {
			// 	  - Write the value 1.
			t.tape[t.cursor] = true
			// 	  - Move one slot to the left.
			t.cursor--
			// 	  - Continue with state A.
			t.state = A
		} else { // 	If the current value is 1:
			// 	  - Write the value 0.
			t.tape[t.cursor] = false
			// 	  - Move one slot to the right.
			t.cursor++
			// 	  - Continue with state B.
			t.state = B
		}
	} else if t.state == F { // In state F:
		// 	If the current value is 0:
		if !t.tape[t.cursor] {
			// 	  - Write the value 0.
			t.tape[t.cursor] = false
			// 	  - Move one slot to the right.
			t.cursor++
			// 	  - Continue with state C.
			t.state = C
		} else { // 	If the current value is 1:
			// 	  - Write the value 0.
			t.tape[t.cursor] = false
			// 	  - Move one slot to the right.
			t.cursor++
			// 	  - Continue with state E.
			t.state = E
		}
	} else {
		panic("bad state")
	}
}

func (t *turingmachine) printTape(from int, length int) {
	for i := from; i < from+length; i++ {
		val := 0
		if t.tape[i] {
			val = 1
		}
		if i == t.cursor {
			fmt.Print("[")
		}
		fmt.Printf("%d", val)
		if i == t.cursor {
			fmt.Print("]")
		}
		fmt.Print(" ")
	}
	fmt.Print("\n")
}

func (t *turingmachine) checkSum() int {
	check := 0
	for _, v := range t.tape {
		if v {
			check++
		}
	}
	return check
}
