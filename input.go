package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Input gets input from the console and stores it in
func Input(command string) {
	splitCmd := strings.Split(command, " ")

	fmt.Println("Insert your value:")
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	variables[splitCmd[1]] = text
}
