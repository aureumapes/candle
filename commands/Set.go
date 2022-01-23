// By AureumApes

package commands

import (
	"strconv"
	"strings"
)

func Set(command string) {
	splitedCmd := strings.Split(command, " ")
	if splitedCmd[2] == "to" {
		if isNumeric(splitedCmd[3]) {
			content, _ := strconv.ParseFloat(splitedCmd[3], 64)
			variables[splitedCmd[1]] = content
		} else {
			variables[splitedCmd[1]] = splitedCmd[3]
		}
	}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
