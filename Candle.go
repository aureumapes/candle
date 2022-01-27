// By AureumApes
package main

import (
	"bufio"
	"fmt"
	"os"
)

var programm = make(map[int]string)
var pos = 0

func main() {
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
	for pos < len(programm) {
		Execute(programm[pos])
		pos++
	}
}
