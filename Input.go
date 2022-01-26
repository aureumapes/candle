// By TUBS1001

package main

import (
	"bufio"
	"os"
	"strings"
)

func Input(command string) {
	splitedCmd := strings.Split(command, " ")

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	variables[splitedCmd[1]] = text
}
