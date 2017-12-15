package main

import (
	"fmt"
)

func main() {
	testJudge := judge{
		genA: makeGenerator(testInputGenA, multiplierGenA),
		genB: makeGenerator(testInputGenB, multiplierGenB),
	}
	fmt.Println("Test part1:", testJudge.part1FindMatchesAfterNRounds(40000000))
	testJudge.reset()
	fmt.Println("Test part2:", testJudge.part2FindMatchesAfterNRounds(5000000))

	part1Judge := judge{
		genA: makeGenerator(inputGenA, multiplierGenA),
		genB: makeGenerator(inputGenB, multiplierGenB),
	}
	fmt.Println("Part1:", part1Judge.part1FindMatchesAfterNRounds(40000000))
	part1Judge.reset()
	fmt.Println("Part2:", part1Judge.part2FindMatchesAfterNRounds(5000000))
}

type judge struct {
	genA generator
	genB generator
}
type generator struct {
	start      int
	multiplier int
	previous   int
}

func makeGenerator(start, multiplier int) generator {
	return generator{
		start:      start,
		multiplier: multiplier,
		previous:   start,
	}
}

func (g *generator) reset() {
	g.previous = g.start
}

func (g *generator) next() int {
	intermediate := (g.previous * g.multiplier)
	g.previous = intermediate - (2147483647 * (intermediate / 2147483647))
	return g.previous
}
func (g *generator) nextMultipleOf4() int {
	for {
		intermediate := (g.previous * g.multiplier)
		g.previous = intermediate - (2147483647 * (intermediate / 2147483647))
		if (g.previous & 0x3) == 0 {
			break
		}
	}

	return g.previous
}
func (g *generator) nextMultipleOf8() int {
	for {
		intermediate := (g.previous * g.multiplier)
		g.previous = intermediate - (2147483647 * (intermediate / 2147483647))
		if (g.previous & 0x7) == 0 {
			break
		}
	}

	return g.previous
}

func (j *judge) reset() {
	j.genA.reset()
	j.genB.reset()
}

func (j *judge) part1FindMatchesAfterNRounds(rounds int) int {
	matches, a, b := 0, 0, 0
	for rounds > 0 {
		a, b = j.genA.next(), j.genB.next()
		if a&0xffff == b&0xffff {
			matches++
		}
		rounds--
	}
	return matches
}

func (j *judge) part2FindMatchesAfterNRounds(rounds int) int {
	matches, a, b := 0, 0, 0
	for rounds > 0 {
		a, b = j.genA.nextMultipleOf4(), j.genB.nextMultipleOf8()
		if a&0xffff == b&0xffff {
			matches++
		}
		rounds--
	}
	return matches
}
