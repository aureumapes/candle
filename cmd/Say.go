package cmd

import "strings"

func Say(args []string, vars map[string]string) {
	_, variableExists := vars[args[0]]
	if variableExists {
		print(vars[args[0]])
	} else {
		print(strings.Join(args, " "))
	}
}

func Sayln(args []string, vars map[string]string) {
	_, variableExists := vars[args[0]]
	if variableExists {
		println(vars[args[0]])
	} else {
		println(strings.Join(args, " "))
	}
}
