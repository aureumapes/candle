package internal

import (
	"strings"
)

func Set(args []string, vars map[string]string) map[string]string {
	vars[args[0]] = strings.Join(args[2:], " ")
	return vars
}
