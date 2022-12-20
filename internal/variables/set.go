package internal

import (
	"Candle/util"
	"regexp"
	"strings"
)

func Set(args []string, vars map[string]string) map[string]string {
	if res, _ := regexp.MatchString("[a-z]+[0-9]*", args[0]); !res {
		util.PrintError("\"" + args[0] + "\" does not match the RegEx \"[a-z]+[0-9]*\"")
	}
	if util.IsNumeric(args[2]) || strings.HasPrefix(args[2], "+") {
		vars[args[0]] = strings.Replace(strings.Join(args[2:], " "), "+", "", 1)
	} else if _, varExist := vars[args[2]]; varExist {
		vars[args[0]] = vars[args[2]]
	} else {
		util.PrintError("\"" + args[2] + "\" is not a Variable, nor numeric or introduced as a String")
	}
	return vars
}
