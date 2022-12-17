package cmd

import (
	"Candle/util"
	"strings"
)

func Set(args []string, vars map[string]string) map[string]string {
	if util.IsNumeric(args[2]) {
		vars[args[0]] = args[2]
	} else {
		vars[args[0]] = strings.Join(args[2:], " ")
	}
	return vars
}
