package commands

import (
	"strconv"
	"strings"
)

func If(command string) {
	splitedCommand := strings.Split(command, " ")
	variable := getFloat(variables[splitedCommand[1]])
	number, _ := strconv.ParseFloat(splitedCommand[3], 64)
	if splitedCommand[2] == "==" {
		if variable == number {
			executeCommand := strings.Join(splitedCommand[4:], " ")
			Execute(executeCommand)
		}
	}
	if splitedCommand[2] == "!=" {
		if variable != number {
			executeCommand := strings.Join(splitedCommand[4:], " ")
			Execute(executeCommand)
		}
	}
	if splitedCommand[2] == ">=" {
		if variable >= number {
			executeCommand := strings.Join(splitedCommand[4:], " ")
			Execute(executeCommand)
		}
	}
	if splitedCommand[2] == "<=" {
		if variable <= number {
			executeCommand := strings.Join(splitedCommand[4:], " ")
			Execute(executeCommand)
		}
	}
}
