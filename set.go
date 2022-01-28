package main

import (
	"strconv"
	"strings"
)

func Set(command string) {
	splitCmd := strings.Split(command, " ")
	if splitCmd[2] == "to" {
		if isNumeric(splitCmd[3]) {
			content, _ := strconv.ParseFloat(splitCmd[3], 64)
			variables[splitCmd[1]] = content
		} else {
			variables[splitCmd[1]] = splitCmd[3]
		}
	}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
