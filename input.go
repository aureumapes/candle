package main

import (
	"bufio"
	"os"
	"strings"
)

func Input(command string) {
	splitCmd := strings.Split(command, " ")

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	variables[splitCmd[1]] = text
}
