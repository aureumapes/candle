// By AureumApes
package main

import (
	"HotdogLang/commands"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var programm = make(map[int]string)
	args := os.Args[1:]
	readFile, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	position := 0
	for fileScanner.Scan() {
		programm[position] = fileScanner.Text()
		position++
	}
	readFile.Close()
	pos := 0
	for pos < len(programm) {
		if strings.HasPrefix(programm[pos], "goto") {
			newPos, _ := strconv.Atoi(strings.Split(programm[pos], " ")[1])
			pos = newPos - 1
		} else {
			commands.Execute(programm[pos])
			pos++
		}
	}
}
