package commands

import (
	"fmt"
	"strings"
)

func Say(command string) {
	splitedCommand := strings.Split(command, " ")
	_, ok := variables[splitedCommand[1]]
	if ok {
		fmt.Println(variables[splitedCommand[1]])
	} else {
		fmt.Println(splitedCommand[1] + ": Unknown Variable")
	}
}
