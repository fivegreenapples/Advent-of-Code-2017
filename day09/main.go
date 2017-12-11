package main

import (
	"errors"
	"fmt"
)

func main() {
	for _, t := range testInput {
		s := groupStream{
			raw: t,
		}
		err := s.parse()
		if err != nil {
			fmt.Printf("%s: %s\n", t, err)
			continue
		}
		fmt.Printf("%s: score is %d, cancelled garbage is %d\n", t, s.total, s.cancelledGarbage)
	}

	part1Stream := groupStream{
		raw: input,
	}
	err := part1Stream.parse()
	if err != nil {
		fmt.Println("Error with part1 input:", err)
	}
	fmt.Println("Part1 score:", part1Stream.total)
	fmt.Println("Part2 cancelled garbage:", part1Stream.cancelledGarbage)

}

type groupStream struct {
	raw              string
	pointer          int
	state            string
	nestLevel        int
	total            int
	cancelledGarbage int
}

func (g *groupStream) parse() error {

	g.pointer = 0
	g.state = "expectItem"

	if len(g.raw) == 0 {
		return errors.New("no input")
	}

	for {
		done, err := g.parseOne()
		if done {
			return err
		}
		g.pointer++
	}
}

func (g *groupStream) parseOne() (bool, error) {

	if g.pointer >= len(g.raw) {
		if g.state == "expectEnd" {
			return true, nil
		}
		return true, errors.New("unexpected end of input when in state: " + g.state)
	}

	char := string(g.raw[g.pointer])

	switch g.state {
	case "expectItem":
		if char == "{" {
			g.nestLevel++
			g.total += g.nestLevel
			// stay expectingitem
		} else if char == "}" {
			g.nestLevel--
			if g.nestLevel == 0 {
				g.state = "expectEnd"
			} else {
				g.state = "endedItem"
			}
		} else if char == "<" {
			g.state = "inGarbage"
		} else {
			return true, errors.New("unexpected char, " + char + ", in state " + g.state)
		}
	case "endedItem":
		if char == "," {
			g.state = "expectItem"
		} else if char == "}" {
			g.nestLevel--
			// stay endedItem unless finished
			if g.nestLevel == 0 {
				g.state = "expectEnd"
			}
		} else {
			return true, errors.New("unexpected char, " + char + ", in state " + g.state)
		}
	case "inGarbage":
		if char == "!" {
			g.state = "escapeNextInGarbage"
		} else if char == ">" {
			g.state = "endedItem"
		} else {
			g.cancelledGarbage++
			// stay inGarbage
		}
	case "escapeNextInGarbage":
		g.state = "inGarbage"
	default:
		return true, errors.New("unexpected state " + g.state + " with char " + char)
	}

	return false, nil
}
