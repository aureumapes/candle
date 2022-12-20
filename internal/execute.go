package internal

import (
	"Candle/internal/calc"
	"Candle/internal/variables"
	"os"
	"strings"
)

func Execute(rawCommand string, vars map[string]string) map[string]string {
	if rawCommand == "exit" {
		os.Exit(1)
	}
	command := strings.Split(rawCommand, " ")
	switch command[0] {
	case "say":
		Say(command[1:], vars)
	case "sayln":
		Sayln(command[1:], vars)
	case "set":
		vars = variables.Set(command[1:], vars)
	case "ask":
		vars = variables.Ask(command[1:], vars)
	case "add":
		vars = calc.Add(command[1:], vars)
	case "sub":
		vars = calc.Sub(command[1:], vars)
	case "mul":
		vars = calc.Mul(command[1:], vars)
	case "div":
		vars = calc.Div(command[1:], vars)
	case "if":
		If(command[1:], vars)
	}
	return vars
}
