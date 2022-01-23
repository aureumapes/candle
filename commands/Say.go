package commands

import (
	"fmt"
	"log"
	"strings"
)

func Say(command string) {
	splitedCommand := strings.Split(command, " ")
	_, ok := variables[splitedCommand[1]]
	if ok {
		fmt.Println(variables[splitedCommand[1]])
	} else {
		log.Fatal(splitedCommand[1] + " is not declared!")
	}
}
