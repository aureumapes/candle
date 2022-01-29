package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ask gets input from the console and stores it in
func Ask(command string) {
	splitCmd := strings.Split(command, " ")
	var message = "Insert your value:"

	if len(splitCmd) > 2 {
		// merge splitCmd[1:] into one string
		message = ""
		for i := 2; i < len(splitCmd); i++ {
			message += splitCmd[i] + " "
		}
	}

	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	//if !isNumeric(text) {
	//	panic(fmt.Sprintf("%s is not a number", text))
	//}
	variables[splitCmd[1]] = text
}
