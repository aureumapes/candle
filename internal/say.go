package internal

import (
	"github.com/aureumapes/candle/util"
	"strings"
)

func Say(args []string, vars map[string]string) {
	_, variableExists := vars[args[0]]
	if strings.HasPrefix(args[0], "+") {
		print(strings.Replace(strings.Join(args, " "), "+", "", 1))
		return
	} else if variableExists {
		print(vars[args[0]])
		return
	}
	util.PrintError("Unknown variable \"" + args[0] + "\"")
}

func Sayln(args []string, vars map[string]string) {
	_, variableExists := vars[args[0]]
	if strings.HasPrefix(args[0], "+") {
		println(strings.Replace(strings.Join(args, " "), "+", "", 1))
		return
	} else if variableExists {
		println(vars[args[0]])
		return
	}
	util.PrintError("Unknown variable \"" + args[0] + "\"")
}
