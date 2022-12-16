package main

import (
	"strconv"
	"strings"
)

// If checks the if conditions
func If(command string) {
	splitCmd := strings.Split(command, " ")
	variable := getFloat(variables[splitCmd[1]])
	number := getFloat(splitCmd[3])

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
