package cmd

import (
	"bufio"
	"os"
	"strings"
)

func Ask(args []string, vars map[string]string) map[string]string {
	print(strings.Join(args[1:], " ") + "\n")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	vars[args[0]] = input
	return vars
}
