package multicoretablet

import (
	"fmt"
	"regexp"
	"strconv"
)

type singleCore struct {
	parent       *Tablet
	coreID       int
	instructions []instruction
	registers    map[string]int
	playedSounds []int
	stats        map[string]int
}

// MakeCore constructs a Tablet Core from a set of programming instructions
func MakeCore(parent *Tablet, id int, rawInstructions []string) singleCore {

	core := singleCore{
		parent:       parent,
		coreID:       id,
		instructions: []instruction{},
		registers:    map[string]int{},
		stats:        map[string]int{},
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
				core.registers[operandAReg] = 0
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
				core.registers[operandBReg] = 0
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

			core.instructions = append(core.instructions, newInstruction)
			continue
		}
		panic(fmt.Errorf("instruction %s not understood", i))
	}

	return core
}

func (c *singleCore) reset() {
	for r := range c.registers {
		c.registers[r] = 0
	}
	c.registers["p"] = c.coreID
}

func (c *singleCore) Run() {

	c.reset()

	programCounter := 0
	for programCounter < len(c.instructions) {

		thisInstruction := c.instructions[programCounter]
		programCounter++
		c.stats[thisInstruction.command]++

		switch thisInstruction.command {
		case "set":
			c.registers[thisInstruction.operandA.inputRegister] = c.valueOfOperand(thisInstruction.operandB)
		case "add":
			c.registers[thisInstruction.operandA.inputRegister] += c.valueOfOperand(thisInstruction.operandB)
		case "sub":
			c.registers[thisInstruction.operandA.inputRegister] -= c.valueOfOperand(thisInstruction.operandB)
		case "mul":
			c.registers[thisInstruction.operandA.inputRegister] *= c.valueOfOperand(thisInstruction.operandB)
		case "mod":
			c.registers[thisInstruction.operandA.inputRegister] = c.valueOfOperand(thisInstruction.operandA) % c.valueOfOperand(thisInstruction.operandB)
		case "snd":
			c.parent.send(c.coreID, c.valueOfOperand(thisInstruction.operandA))
		case "rcv":
			recvVal, err := c.parent.receive(c.coreID)
			if err != nil {
				fmt.Printf("Core %d: error receiving from queue: %s\n", c.coreID, err.Error())
				return
			}
			c.registers[thisInstruction.operandA.inputRegister] = recvVal
		case "jnz":
			if c.valueOfOperand(thisInstruction.operandA) != 0 {
				programCounter += (c.valueOfOperand(thisInstruction.operandB) - 1)
			}
		case "jgz":
			if c.valueOfOperand(thisInstruction.operandA) > 0 {
				programCounter += (c.valueOfOperand(thisInstruction.operandB) - 1)
			}
		default:
			panic("instruction type not understood: " + thisInstruction.command)
		}

	}

	// Dump stats if we end here (used in day 23)
	fmt.Println(c.stats)
	fmt.Printf("Core %d: execution finished. Reached end of program.\n", c.coreID)
}

func (c *singleCore) valueOfOperand(op operand) int {
	if op.typ == "val" {
		return op.inputVal
	}
	return c.registers[op.inputRegister]
}

func (c *singleCore) printRegisters(pc int) {
	fmt.Printf("%d", pc)
	for cc := 'a'; cc <= 'h'; cc++ {
		fmt.Printf("%6d ", c.registers[string(cc)])
	}
	fmt.Printf("\n")

}
