package main

import (
	"strconv"
	"strings"
)

func If(command string) {
	splitCmd := strings.Split(command, " ")
	variable := getFloat(variables[splitCmd[1]])
	number, _ := strconv.ParseFloat(splitCmd[3], 64)

	switch splitCmd[2] {
	case "==":
		if variable == number {
			executeCommand := strings.Join(splitCmd[4:], " ")
			Execute(executeCommand)
		}
		break
	case "!=":
		if variable != number {
			executeCommand := strings.Join(splitCmd[4:], " ")
			Execute(executeCommand)
		}
		break
	case ">=":
		if variable >= number {
			executeCommand := strings.Join(splitCmd[4:], " ")
			Execute(executeCommand)
		}
		break
	case "<=":
		if variable <= number {
			executeCommand := strings.Join(splitCmd[4:], " ")
			Execute(executeCommand)
		}
		break
	case "<":
		if variable < number {
			executeCommand := strings.Join(splitCmd[4:], " ")
			Execute(executeCommand)
		}
		break
	case ">":
		if variable > number {
			executeCommand := strings.Join(splitCmd[4:], " ")
			Execute(executeCommand)
		}
		break
	}
}
