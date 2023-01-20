package internal

import (
	"github.com/Amiraxoba/candle/util"
	"strings"
)

func CreateFunction(args []string, funcs map[string][]string, funcBody []string) map[string][]string {
	funcBody = strings.Split(strings.Split(strings.Join(funcBody, "\n"), "end")[0], "\n")

	for line, command := range funcBody {
		if line > len(funcBody)-1 {
			break
		}
		if command == "" || strings.HasPrefix(command, "//") {
			funcBody = util.RemoveFromArray(funcBody, line)
		}
	}

	for _, funCmd := range funcBody {
		funcs[args[0]] = append(funcs[args[0]], funCmd)
	}
	return funcs
}

func ExecuteFunction(command []string, vars map[string]string, funcs map[string][]string) (map[string]string, map[string][]string) {
	funcCommands, funcExists := funcs[command[0]]
	if !funcExists {
		util.PrintError("Unknown function\"" + command[0] + "\"")
	}
	for _, funcCommand := range funcCommands {
		vars, funcs = Execute(funcCommand, vars, funcs, []string{})
	}
	return vars, funcs
}
