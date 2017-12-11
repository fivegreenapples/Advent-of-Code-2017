package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	testInstructions := parseInput(testInput)
	testRegisters := registerBank{}
	for _, i := range testInstructions {
		testRegisters.applyInstruction(i)
	}
	fmt.Println("Test registers max val is: ", testRegisters.maxRegValue())
	fmt.Println("Test registers max ever val is: ", testRegisters.maxEver)

	part1Instructions := parseInput(input)
	part1Registers := registerBank{}
	for _, i := range part1Instructions {
		part1Registers.applyInstruction(i)
	}
	fmt.Println("Part1 registers max val is: ", part1Registers.maxRegValue())
	fmt.Println("Part1 registers max ever val is: ", part1Registers.maxEver)
}

func parseInput(input string) []instruction {
	rawInstructions := strings.Split(input, "\n")
	instructions := []instruction{}
	instructionRegex := regexp.MustCompile(`^([a-z]+) (inc|dec) (-?[0-9]+) if ([a-z]+) ([!><=]+) (-?[0-9]+)$`)
	for _, i := range rawInstructions {
		matches := instructionRegex.FindStringSubmatch(i)
		if matches == nil {
			panic("raw instruction didn't match regex: " + i)
		}
		newInstruction := makeInstruction(matches[1], matches[2], matches[3], matches[4], matches[5], matches[6])
		instructions = append(instructions, newInstruction)
	}
	return instructions
}

func makeInstruction(reg, deltaType, delta, conditionReg, conditionType, conditionVal string) instruction {
	deltaVal, err := strconv.Atoi(delta)
	if err != nil {
		panic("failed converting deltaval (" + delta + ") to int: " + err.Error())
	}
	if deltaType == "dec" {
		deltaVal *= -1
	} else if deltaType != "inc" {
		panic("unrecognised delta type: " + deltaType)
	}

	return instruction{
		register: reg,
		delta:    deltaVal,
		cond:     makeCondition(conditionReg, conditionType, conditionVal),
	}
}
func makeCondition(conditionReg, conditionType, conditionVal string) instructionCondition {
	var test conditionTest
	switch conditionType {
	case ">":
		test = greater
	case ">=":
		test = greaterEq
	case "<":
		test = less
	case "<=":
		test = lessEq
	case "==":
		test = eq
	case "!=":
		test = notEq
	default:
		panic("unrecognised conditiontype: " + conditionType)
	}

	conditionValue, err := strconv.Atoi(conditionVal)
	if err != nil {
		panic("failed converting conditionVal (" + conditionVal + ") to int: " + err.Error())
	}

	return instructionCondition{
		register: conditionReg,
		val:      conditionValue,
		test:     test,
	}
}

type instruction struct {
	register string
	delta    int
	cond     instructionCondition
}
type instructionCondition struct {
	register string
	val      int
	test     conditionTest
}
type conditionTest int

const (
	greater conditionTest = iota
	less
	greaterEq
	lessEq
	eq
	notEq
)

type registerBank struct {
	registers map[string]int
	maxEver   int
}

func (r *registerBank) applyInstruction(ins instruction) {

	if r.registers == nil {
		r.registers = map[string]int{}
	}

	if !r.doesTestPass(ins.cond) {
		return
	}

	if _, found := r.registers[ins.register]; !found {
		r.registers[ins.register] = 0
	}

	r.registers[ins.register] += ins.delta
	if r.registers[ins.register] > r.maxEver {
		r.maxEver = r.registers[ins.register]
	}
}

func (r *registerBank) doesTestPass(cond instructionCondition) bool {
	if r.registers == nil {
		return false
	}

	if _, found := r.registers[cond.register]; !found {
		r.registers[cond.register] = 0
	}

	switch cond.test {
	case greater:
		return r.registers[cond.register] > cond.val
	case greaterEq:
		return r.registers[cond.register] >= cond.val
	case less:
		return r.registers[cond.register] < cond.val
	case lessEq:
		return r.registers[cond.register] <= cond.val
	case eq:
		return r.registers[cond.register] == cond.val
	case notEq:
		return r.registers[cond.register] != cond.val
	default:
		panic("unhandled condition test type")
	}

}
func (r *registerBank) maxRegValue() int {
	if r.registers == nil {
		return 0
	}

	var maxVal int
	for _, val := range r.registers {
		maxVal = val
		break
	}
	for _, val := range r.registers {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
