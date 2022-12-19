package internal

import (
	"Candle/util"
	"strings"
)

func If(args []string, vars map[string]string) {
	variable1, varExists := vars[args[0]]
	if !varExists {
		util.PrintError("\"" + args[0] + "\" is not a variable")
	}
	_, varExists = vars[args[2]]
	if !(util.IsNumeric(args[2]) || varExists || strings.HasPrefix(args[2], "+")) {
		util.PrintError("\"" + args[2] + "\" is not numeric, nor an variable or introduced as a string")
	}
	switch args[1] {
	case "==":
		if variable1 == vars[args[2]] || variable1 == strings.Replace(args[2], "+", "", 1) {
			Execute(strings.Join(args[3:], " "), vars)
		}
	case "!=":
		if variable1 != vars[args[2]] || variable1 != strings.Replace(args[2], "+", "", 1) {
			Execute(strings.Join(args[3:], " "), vars)
		}
	case "<=":
		if variable1 <= vars[args[2]] || variable1 <= strings.Replace(args[2], "+", "", 1) {
			Execute(strings.Join(args[3:], " "), vars)
		}
	case ">=":
		if variable1 >= vars[args[2]] || variable1 >= strings.Replace(args[2], "+", "", 1) {
			Execute(strings.Join(args[3:], " "), vars)
		}
	case "<":
		if variable1 < vars[args[2]] || variable1 < strings.Replace(args[2], "+", "", 1) {
			Execute(strings.Join(args[3:], " "), vars)
		}
	case ">":
		if variable1 > vars[args[2]] || variable1 > strings.Replace(args[2], "+", "", 1) {
			Execute(strings.Join(args[3:], " "), vars)
		}
	}
}
