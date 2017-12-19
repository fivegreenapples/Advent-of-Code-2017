package multicoretablet

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
