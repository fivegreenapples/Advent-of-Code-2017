package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	testShouts := parseInput(testInput)
	testTower := createTower(testShouts)
	fmt.Println("Test1 root is ", testTower.p.name)
	fmt.Println("Test1 weight is ", testTower.Weight())
	fmt.Println("Test1 balancedness is ", testTower.IsBalanced())
	fmt.Println("Test1 balancing weight is ", testTower.FindBalancingWeight())

	part1Shouts := parseInput(input)
	part1Tower := createTower(part1Shouts)
	fmt.Println("Answer1 root is ", part1Tower.p.name)
	fmt.Println("Answer1 weight is ", part1Tower.Weight())
	fmt.Println("Answer1 balancedness is ", part1Tower.IsBalanced())
	fmt.Println("Answer1 balancing weight is ", part1Tower.FindBalancingWeight())
}

func parseInput(input string) []shout {

	out := []shout{}
	rawShouts := strings.Split(input, "\n")
	shoutRegex := regexp.MustCompile(`^([a-z]+)\s*\(([0-9]+)\)(\s*->\s*([a-z, ]+))?$`)
	subProgramRegex := regexp.MustCompile(`[a-z]+`)
	for _, s := range rawShouts {
		matches := shoutRegex.FindStringSubmatch(s)
		if matches == nil {
			panic("raw shout didn't match regex: " + s)
		}
		name := matches[1]
		weight, weightErr := strconv.Atoi(matches[2])
		if weightErr != nil {
			panic("error converting weight to int: " + weightErr.Error())
		}
		subPrograms := []programName{}
		if len(matches) >= 5 && matches[4] != "" {
			subProgramRaw := subProgramRegex.FindAllString(matches[4], -1)
			if len(subProgramRaw) == 0 {
				panic("unexpected zero subprograms from raw shout: " + s)
			}
			for _, p := range subProgramRaw {
				subPrograms = append(subPrograms, programName(p))
			}
		}

		newShout := shout{
			name:        programName(name),
			weight:      weight,
			subPrograms: subPrograms,
		}
		out = append(out, newShout)
	}
	return out
}

func createTower(shouts []shout) tower {

	// To create the tower we first make a map of all programs
	// While doing this, keep track of programs that are holding sub towers
	allPrograms := map[programName]tower{}
	towersWithSubTowers := []*tower{}

	for _, s := range shouts {
		newTower := tower{
			p: program{
				name:   s.name,
				weight: s.weight,
			},
			weight:    -1,
			subTowers: map[programName]tower{},
		}
		for _, subP := range s.subPrograms {
			newTower.subTowers[subP] = tower{}
		}
		if len(newTower.subTowers) > 0 {
			towersWithSubTowers = append(towersWithSubTowers, &newTower)
		}
		allPrograms[s.name] = newTower
	}

	// Now we go through the with-sub-tower list and move towers from the

	for _, t := range towersWithSubTowers {
		for program := range t.subTowers {
			st, found := allPrograms[program]
			if !found {
				panic(program + " not found in tower")
			}
			t.subTowers[program] = st
			delete(allPrograms, program)
		}
	}

	// Now we should have only one tower in our map
	if len(allPrograms) != 1 {
		panic("tower rearrangement failed to reduce to single tower")
	}

	for rootProgram := range allPrograms {
		return allPrograms[rootProgram]
	}

	panic("tower rearrangement failed")
}

type shout struct {
	name        programName
	weight      int
	subPrograms []programName
}
type programName string
type program struct {
	name   programName
	weight int
}
type tower struct {
	p         program
	weight    int
	subTowers map[programName]tower
}

func (t *tower) Weight() int {
	if t.weight >= 0 {
		return t.weight
	}
	t.weight = t.p.weight
	for _, st := range t.subTowers {
		t.weight += st.Weight()
	}
	return t.weight
}
func (t *tower) IsBalanced() bool {
	weight := -1
	for _, st := range t.subTowers {
		if weight == -1 {
			weight = st.Weight()
		} else if weight != st.Weight() {
			return false
		}
	}
	return true
}
func (t *tower) FindBalancingWeight() int {

	for _, st := range t.subTowers {
		if !st.IsBalanced() {
			return st.FindBalancingWeight()
		}
	}

	weightCounts := map[int]int{}
	programByWeight := map[int]programName{}
	for pn, st := range t.subTowers {
		thisWeight := st.Weight()
		programByWeight[thisWeight] = pn
		if _, found := weightCounts[thisWeight]; !found {
			weightCounts[thisWeight] = 0
		}
		weightCounts[thisWeight] = weightCounts[thisWeight] + 1
	}
	badWeight := -1
	okWeight := -1
	for w, count := range weightCounts {
		if count == 1 {
			badWeight = w
		} else {
			okWeight = w
		}
	}
	diffWeight := badWeight - okWeight
	badProgram := t.subTowers[programByWeight[badWeight]].p
	return badProgram.weight - diffWeight
}
