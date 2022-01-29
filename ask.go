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

	fmt.Println("Insert your value:")
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	variables[splitCmd[1]] = text
}
