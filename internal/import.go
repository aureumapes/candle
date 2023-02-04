package internal

import (
	"errors"
	"github.com/Amiraxoba2/candle/util"
	"os"
	"strings"
)

func Import(args []string, funcs map[string][]string) map[string][]string {
	funFile, err := os.ReadFile("./" + args[0] + ".cndl")
	if errors.Is(err, os.ErrNotExist) {
		util.PrintError("Can't find function " + args[0])
	}
	funcs[args[0]] = strings.Split(string(funFile), "\n")
	return funcs
}
