package variables

import (
	"errors"
	"github.com/Amiraxoba/candle/util"
	"os"
)

func Read(args []string, vars map[string]string) map[string]string {
	fileBytes, err := os.ReadFile(args[0])
	if errors.Is(err, os.ErrNotExist) {
		util.PrintError("File " + args[0] + "does not exist")
	}
	fileContent := string(fileBytes)
	vars[args[1]] = fileContent
	return vars
}
