package internal

import (
	"github.com/Amiraxoba2/candle/internal/calc"
	"github.com/Amiraxoba2/candle/internal/variables"
	"os"
	"strings"
)

func Execute(rawCommand string, vars map[string]string, funcs map[string][]string, funcBody []string) (map[string]string, map[string][]string) {
	if rawCommand == "exit" {
		os.Exit(1)
	}
	command := strings.Split(rawCommand, " ")
	switch command[0] {
	case "exit":
		os.Exit(0)
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
		vars, funcs = If(command[1:], vars, funcs)
	case "function":
		funcs = CreateFunction(command[1:], funcs, funcBody)
	case "end":
		break
	case "read":
		vars = variables.Read(command[1:], vars)
	default:
		vars, funcs = ExecuteFunction(command, vars, funcs)
	}
	return vars, funcs
}
