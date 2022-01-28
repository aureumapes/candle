package main

import (
	"strconv"
	"strings"
)

// Execute parses and executes an instruction
func Execute(command string) {
	if strings.HasPrefix(command, "goto") {
		newPos, _ := strconv.Atoi(strings.Split(command, " ")[1])
		pos = newPos - 1
	}
	if strings.HasPrefix(command, "say") {
		Say(command)
	} else if strings.HasPrefix(command, "set") {
		Set(command)
	} else if strings.HasPrefix(command, "add") {
		Add(command)
	} else if strings.HasPrefix(command, "sub") {
		Sub(command)
	} else if strings.HasPrefix(command, "mul") {
		Mul(command)
	} else if strings.HasPrefix(command, "div") {
		Div(command)
	} else if strings.HasPrefix(command, "input") {
		Input(command)
	} else if strings.HasPrefix(command, "sqrt") {
		Sqrt(command)
	} else if strings.HasPrefix(command, "if") {
		If(command)
	}
}
