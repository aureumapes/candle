package internal

import (
	"github.com/Amiraxoba2/candle/util"
	"strings"
)

func If(args []string, vars map[string]string, funcs map[string][]string) (map[string]string, map[string][]string) {
	variable1, varExists := vars[args[0]]
	if !varExists {
		util.PrintError("\"" + args[0] + "\" is not a variable")
	}
	_, varExists = vars[args[2]]
	if !(util.IsNumeric(args[2]) || varExists || strings.HasPrefix(args[2], "+")) {
		vars, funcs = Execute(strings.Join(args[3:], " "), vars, make(map[string][]string), []string{})
	}
	switch args[1] {
	case "==":
		if variable1 == vars[args[2]] || variable1 == strings.Replace(args[2], "+", "", 1) {
			vars, funcs = Execute(strings.Join(args[3:], " "), vars, make(map[string][]string), []string{})
		}
	case "!=":
		if variable1 != vars[args[2]] || variable1 != strings.Replace(args[2], "+", "", 1) {
			vars, funcs = Execute(strings.Join(args[3:], " "), vars, make(map[string][]string), []string{})
		}
	case "<=":
		if variable1 <= vars[args[2]] || variable1 <= strings.Replace(args[2], "+", "", 1) {
			vars, funcs = Execute(strings.Join(args[3:], " "), vars, make(map[string][]string), []string{})
		}
	case ">=":
		if variable1 >= vars[args[2]] || variable1 >= strings.Replace(args[2], "+", "", 1) {
			vars, funcs = Execute(strings.Join(args[3:], " "), vars, make(map[string][]string), []string{})
		}
	case "<":
		if variable1 < vars[args[2]] || variable1 < strings.Replace(args[2], "+", "", 1) {
			vars, funcs = Execute(strings.Join(args[3:], " "), vars, make(map[string][]string), []string{})
		}
	case ">":
		if variable1 > vars[args[2]] || variable1 > strings.Replace(args[2], "+", "", 1) {
			vars, funcs = Execute(strings.Join(args[3:], " "), vars, make(map[string][]string), []string{})
		}
	}
	return vars, funcs
}
