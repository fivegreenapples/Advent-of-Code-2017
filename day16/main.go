package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	testTroupe := makeTroupe(testTroupe)
	testMoves := parseInput(testInput)
	for _, m := range testMoves {
		m.applyToTroupe(&testTroupe)
	}
	fmt.Println("Test part1:", testTroupe.String())

	part1Troupe := makeTroupe(part1Troupe)
	part1Moves := parseInput(strings.Split(part1Input, ","))
	for _, m := range part1Moves {
		m.applyToTroupe(&part1Troupe)
	}
	fmt.Println("Part1:", part1Troupe.String())

	testTroupe.applyMoves(testMoves, 1000000000-1)
	fmt.Println("Test part2:", testTroupe.String())
	part1Troupe.applyMoves(part1Moves, 1000000000-1)
	fmt.Println("Part2:", part1Troupe.String())

}

type spin int

func (s spin) applyToTroupe(t *troupe) {
	t.spin(int(s))
}

type exchange struct {
	a, b int
}

func (e exchange) applyToTroupe(t *troupe) {
	t.exchange(e.a, e.b)
}

type partner struct {
	a, b rune
}

func (p partner) applyToTroupe(t *troupe) {
	t.partner(p.a, p.b)
}

type danceMove interface {
	applyToTroupe(t *troupe)
}

type troupe struct {
	programs           []rune
	positionsByProgram map[rune]int
	startOffset        int
}

func (t *troupe) programAt(pos int) rune {
	return t.programs[(t.startOffset+pos)%len(t.programs)]
}
func (t *troupe) positionFor(program rune) int {
	i, found := t.positionsByProgram[program]
	if !found {
		panic("Invalid program")
	}
	return (len(t.programs) + i - t.startOffset) % len(t.programs)
}
func (t *troupe) setProgramAt(pos int, p rune) {
	t.programs[(t.startOffset+pos)%len(t.programs)] = p
	t.positionsByProgram[p] = (t.startOffset + pos) % len(t.programs)
}

func (t *troupe) spin(amount int) {
	if amount < 1 || amount > len(t.programs) {
		panic("Invalid spin amount")
	}
	t.startOffset = (t.startOffset + len(t.programs) - amount) % len(t.programs)
}
func (t *troupe) exchange(a, b int) {
	if a < 0 || a >= len(t.programs) || b < 0 || b >= len(t.programs) {
		panic("Invalid exchange indices")
	}
	pA := t.programAt(a)
	pB := t.programAt(b)
	t.setProgramAt(a, pB)
	t.setProgramAt(b, pA)
}
func (t *troupe) partner(a, b rune) {
	iA := t.positionFor(a)
	iB := t.positionFor(b)
	t.setProgramAt(iA, b)
	t.setProgramAt(iB, a)
}

func (t *troupe) String() string {
	out := []rune{}
	for i := 0; i < len(t.programs); i++ {
		out = append(out, t.programAt(i))
	}
	return string(out)
}

type idx struct {
	pattern   string
	moveIndex int
}

func (t *troupe) applyMoves(moves []danceMove, cycles int) {
	seen := map[idx]int{}
	total := cycles * len(moves)
	rounds := 0
	for rounds < total {
		thisMove := moves[rounds%len(moves)]
		thisMove.applyToTroupe(t)

		thisIdx := idx{t.String(), (rounds % len(moves))}
		if where, found := seen[thisIdx]; found {
			loopFactor := rounds - where
			remaining := total - rounds - 1
			skip := loopFactor * (remaining / loopFactor)
			rounds += skip
		} else {
			seen[thisIdx] = rounds
		}
		rounds++
	}

}

func parseInput(in []string) []danceMove {

	reSpin := regexp.MustCompile(`^s([0-9]+)$`)
	reExchange := regexp.MustCompile(`^x([0-9]+)/([0-9]+)$`)
	rePartner := regexp.MustCompile(`^p([a-z])/([a-z])$`)

	out := []danceMove{}
	for _, m := range in {

		if matches := reSpin.FindStringSubmatch(m); matches != nil {
			spinVal, err := strconv.Atoi(matches[1])
			if err != nil {
				panic("Input spin faild strconv: " + m)
			}
			out = append(out, spin(spinVal))
			continue
		}

		if matches := reExchange.FindStringSubmatch(m); matches != nil {
			aVal, errA := strconv.Atoi(matches[1])
			bVal, errB := strconv.Atoi(matches[2])
			if errA != nil || errB != nil {
				panic("Input spin faild strconv: " + m)
			}
			out = append(out, exchange{aVal, bVal})
			continue
		}

		if matches := rePartner.FindStringSubmatch(m); matches != nil {
			aVal := rune(matches[1][0])
			bVal := rune(matches[2][0])
			out = append(out, partner{aVal, bVal})
			continue
		}

		panic("Input move didn't match any regex: " + m)
	}

	return out
}

func makeTroupe(in string) troupe {
	t := troupe{
		programs:           []rune(in),
		positionsByProgram: map[rune]int{},
		startOffset:        0,
	}
	for i, r := range t.programs {
		t.positionsByProgram[r] = i
	}
	return t
}
