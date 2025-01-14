package variables

import (
	"github.com/aureumapes/candle/util"
	"regexp"
	"strings"
)

func Set(args []string, vars map[string]string) map[string]string {
	if res, _ := regexp.MatchString("[a-z][a-z|0-9]*", args[0]); !res {
		util.PrintError("\"" + args[0] + "\" does not match the RegEx \"[a-z]+[0-9]*\"")
	}
	if util.IsNumeric(args[2]) || strings.HasPrefix(args[2], "+") {
		vars[args[0]] = strings.Replace(strings.Join(args[2:], " "), "+", "", 1)
	} else if args[2] == ":" {
		if len(args[3:])%2 != 0 {
			util.PrintError("Maps must have a value for every key!")
		}
		tmparg := args[3:]
		for index, _ := range tmparg {
			if index%2 == 0 {
				vars[args[0]+"."+tmparg[index]] = tmparg[index+1]
			}
		}
		return vars
	} else if _, varExist := vars[args[2]]; varExist {
		vars[args[0]] = vars[args[2]]
	} else {
		util.PrintError("\"" + args[2] + "\" is not a Variable, nor numeric or introduced as a String")
	}
	return vars
}
