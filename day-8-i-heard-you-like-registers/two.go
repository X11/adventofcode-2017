package main

import (
	"fmt"
	"strings"

	"github.com/X11/adventofcode-2017/util"
)

type Check struct {
	Name  string
	Type  string
	Value int
}

func (c Check) Exec(against int) bool {
	switch c.Type {
	case ">":
		return against > c.Value
	case "<":
		return against < c.Value
	case ">=":
		return against >= c.Value
	case "<=":
		return against <= c.Value
	case "==":
		return against == c.Value
	case "!=":
		return against != c.Value
	}
	return false
}

type Instruction struct {
	Name      string
	Increment bool
	Value     int
	Check     Check
}

func parseInstruction(s string) Instruction {
	parts := strings.Split(s, "if")
	instructionParts := strings.Split(parts[0], " ")
	checkParts := strings.Split(strings.TrimSpace(parts[1]), " ")

	return Instruction{
		Name:      instructionParts[0],
		Increment: instructionParts[1] == "inc",
		Value:     util.ParseInteger(instructionParts[2]),
		Check: Check{
			Name:  checkParts[0],
			Type:  checkParts[1],
			Value: util.ParseInteger(checkParts[2]),
		},
	}
}

func main() {
	input := util.ReadInputFromFile("./input.txt")

	fmt.Printf("input = %+v\n", input)

	register := make(map[string]int)

	getValue := func(s string) int {
		val, ok := register[s]
		if !ok {
			register[s] = 0
			return 0
		}
		return val
	}

	max := 0
	setValue := func(s string, v int) {
		getValue(s)
		register[s] = v
		if v > max {
			max = v
		}
	}

	for _, instruction := range strings.Split(input, "\n") {
		if instruction == "" {
			continue
		}

		inst := parseInstruction(instruction)
		if inst.Check.Exec(getValue(inst.Check.Name)) {
			if inst.Increment {
				setValue(inst.Name, getValue(inst.Name)+inst.Value)
			} else {
				setValue(inst.Name, getValue(inst.Name)-inst.Value)
			}
		}
	}

	fmt.Printf("max = %+v\n", max)
}
