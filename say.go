package main

import (
	"fmt"
	"strings"
)

// Say prints a variable to the console
func Say(command string) {
	splitCmd := strings.Split(command, " ")
	_, ok := variables[splitCmd[1]]
	if ok {
		fmt.Println(variables[splitCmd[1]])
	} else {
		fmt.Println(splitCmd[1] + ": Unknown Variable")
	}
}
