package main

import "fmt"
import "strings"
import "strconv"

func main() {
	testString := makeString(testLength)
	testString.applyTwists(testInput)
	testCurrentSeq := testString.getSequence()
	fmt.Println(testCurrentSeq, testCurrentSeq[0]*testCurrentSeq[1])

	part1String := makeString(inputLength)
	part1String.applyTwists(input)
	part1CurrentSeq := part1String.getSequence()
	fmt.Println(part1CurrentSeq, part1CurrentSeq[0]*part1CurrentSeq[1])

	for test, answer := range part2Tests {
		testAnswer := part2Hash(test)
		fmt.Printf("Test %s %v\n", test, testAnswer == answer)
	}
	fmt.Printf("Part2: %s\n", part2Hash(part2Input))
}
func makeString(length int) stringCircle {
	sc := stringCircle{
		marks: map[int]mark{},
	}
	for i := 0; i < length; i++ {
		sc.marks[i] = mark{
			prev: (i + length - 1) % length,
			self: i,
			next: (i + 1) % length,
		}
	}
	return sc
}

type stringCircle struct {
	marks        map[int]mark
	skip         int
	current      int
	currentIndex int
}

func (sc *stringCircle) applyTwists(twists []int) {
	for _, t := range twists {
		sc.twist(t)
	}
}
func (sc *stringCircle) twist(length int) {
	if length > len(sc.marks) {
		panic("invalid length value: " + strconv.Itoa(length))
	}

	if length > 0 {
		previousMark := sc.marks[sc.current].prev
		startMark := sc.current
		numberToToggle := length
		for numberToToggle > 0 {
			// swap next/previous
			current := sc.marks[sc.current]
			current.prev, current.next = current.next, current.prev
			sc.marks[sc.current] = current
			// move forward & decrement numberToToggle
			// use prev because we've just toggled
			sc.current = sc.marks[sc.current].prev
			numberToToggle--
		}

		if length < len(sc.marks) {
			// now sort out the connections at the extremities of the twist

			endMark := sc.marks[sc.current].prev
			afterEndMark := sc.current

			var aMark mark
			aMark = sc.marks[previousMark]
			aMark.next = endMark
			sc.marks[previousMark] = aMark

			aMark = sc.marks[startMark]
			aMark.next = afterEndMark
			sc.marks[startMark] = aMark

			aMark = sc.marks[endMark]
			aMark.prev = previousMark
			sc.marks[endMark] = aMark

			aMark = sc.marks[afterEndMark]
			aMark.prev = startMark
			sc.marks[afterEndMark] = aMark
		} else {
			// for the specila case of length = len(marks) we need to set sc.current to
			// current.prev as that is the one we're actually on
			sc.current = sc.marks[sc.current].next
		}
	}
	// and move on the right number of skips
	for s := 0; s < sc.skip; s++ {
		sc.current = sc.marks[sc.current].next
	}
	// and move the currentIndex to keep track of where we are in the string
	sc.currentIndex += (length + sc.skip)
	sc.currentIndex = sc.currentIndex % len(sc.marks)
	// and increment the skip value
	sc.skip++
}
func (sc *stringCircle) getSequence() []int {
	// wind back until we are at the start
	curMark := sc.current
	curIndex := sc.currentIndex
	for curIndex > 0 {
		curMark = sc.marks[curMark].prev
		curIndex--
	}
	// move forward collecting marks into a slice
	out := []int{}
	for i := 0; i < len(sc.marks); i++ {
		out = append(out, curMark)
		curMark = sc.marks[curMark].next
	}
	// join slice together
	return out
}
func (sc *stringCircle) String() string {
	seq := sc.getSequence()
	out := []string{}
	for i, s := range seq {
		mark := strconv.Itoa(s)
		if i == sc.currentIndex {
			mark = "[" + mark + "]"
		}
		out = append(out, mark)
	}
	// join slice together
	return strings.Join(out, ",") + " [" + strconv.Itoa(sc.currentIndex) + "] [" + strconv.Itoa(sc.skip) + "]"
}

type mark struct {
	prev int
	self int
	next int
}

func parseInput(in string) []int {
	bytes := []byte(in)
	out := []int{}
	for _, b := range bytes {
		out = append(out, int(b))
	}
	return out
}

func condense(in []int) []int {
	if len(in)%16 != 0 {
		panic("condense needs input slice length divisible by 16")
	}

	out := []int{}
	for i := 0; i < len(in); i += 16 {
		val := 0
		for j := 0; j < 16; j++ {
			val = val ^ in[i+j]
		}
		out = append(out, val)
	}
	return out
}
func part2Hash(in string) string {
	inputLengths := parseInput(in)
	inputLengths = append(inputLengths, part2LengthSuffix...)

	part2String := makeString(part2Length)
	for r := 0; r < 64; r++ {
		part2String.applyTwists(inputLengths)
	}

	denseHash := condense(part2String.getSequence())
	out := ""
	for _, d := range denseHash {
		out += fmt.Sprintf("%.2x", d)
	}
	return out
}
