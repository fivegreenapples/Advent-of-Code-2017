package tablet

import (
	"fmt"
	"regexp"
	"strconv"
)

// Tablet describes a programmable sound making device
type Tablet struct {
	instructions []instruction
	registers    map[string]int
	playedSounds []int
}

// Make constructs a Tablet from a set of programming instructions
func Make(rawInstructions []string) Tablet {

	tab := Tablet{
		instructions: []instruction{},
		registers:    map[string]int{},
	}

	instructionRegex := regexp.MustCompile(`^([a-z]+) (([a-z])|([0-9-]+))( (([a-z])|([0-9-]+)))?$`)
	for _, i := range rawInstructions {
		if matches := instructionRegex.FindStringSubmatch(i); matches != nil {

			command := matches[1]
			operandAReg := matches[3]
			operandAVal := matches[4]
			operandBReg := matches[7]
			operandBVal := matches[8]

			var opA operand
			if operandAReg != "" {
				tab.registers[operandAReg] = 0
				opA = operand{
					typ:           "reg",
					inputRegister: operandAReg,
				}
			} else if operandAVal != "" {
				bVal, err := strconv.Atoi(operandAVal)
				if err != nil {
					panic(fmt.Errorf("failed converting opA value: %s", operandAVal))
				}
				opA = operand{
					typ:      "val",
					inputVal: bVal,
				}
			}
			var opB operand
			if operandBReg != "" {
				tab.registers[operandBReg] = 0
				opB = operand{
					typ:           "reg",
					inputRegister: operandBReg,
				}
			} else if operandBVal != "" {
				bVal, err := strconv.Atoi(operandBVal)
				if err != nil {
					panic(fmt.Errorf("failed converting opB value: %s", operandBVal))
				}
				opB = operand{
					typ:      "val",
					inputVal: bVal,
				}
			}

			newInstruction := instruction{
				command:  command,
				operandA: opA,
				operandB: opB,
			}

			tab.instructions = append(tab.instructions, newInstruction)
			continue
		}
		panic(fmt.Errorf("instruction %s not understood", i))
	}

	return tab
}

// Reset initialises the Tablet to a starting position where
// all internal registers are set to zero.
func (t *Tablet) Reset() {
	for r := range t.registers {
		t.registers[r] = 0
	}
}

// Run executes the Tablet's instructions
func (t *Tablet) Run() int {

	programCounter := 0
	for programCounter < len(t.instructions) {

		thisInstruction := t.instructions[programCounter]
		programCounter++

		switch thisInstruction.command {
		case "set":
			t.registers[thisInstruction.operandA.inputRegister] = t.valueOfOperand(thisInstruction.operandB)
		case "add":
			t.registers[thisInstruction.operandA.inputRegister] += t.valueOfOperand(thisInstruction.operandB)
		case "mul":
			t.registers[thisInstruction.operandA.inputRegister] *= t.valueOfOperand(thisInstruction.operandB)
		case "mod":
			t.registers[thisInstruction.operandA.inputRegister] = t.valueOfOperand(thisInstruction.operandA) % t.valueOfOperand(thisInstruction.operandB)
		case "snd":
			t.playedSounds = append(t.playedSounds, t.valueOfOperand(thisInstruction.operandA))
		case "rcv":
			if t.valueOfOperand(thisInstruction.operandA) > 0 {
				return t.playedSounds[len(t.playedSounds)-1]
			}
		case "jgz":
			if t.valueOfOperand(thisInstruction.operandA) > 0 {
				programCounter += (t.valueOfOperand(thisInstruction.operandB) - 1)
			}
		default:
			panic("instruction type not understood: " + thisInstruction.command)
		}

		// fmt.Println(t.registers, programCounter)
	}

	panic("program exited with no rcv instruction")
}

func (t *Tablet) valueOfOperand(op operand) int {
	if op.typ == "val" {
		return op.inputVal
	}
	return t.registers[op.inputRegister]
}

type instruction struct {
	command  string
	operandA operand
	operandB operand
}

type operand struct {
	typ           string
	inputRegister string
	inputVal      int
}
