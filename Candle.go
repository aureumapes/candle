package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var variables = make(map[string]interface{})
var program = make(map[int]string)
var pos = 0

// main function
func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("Please add a file as an argument!")
	}
	readFile, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	position := 0
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if strings.HasPrefix(text, "#") || text == "" || strings.HasPrefix(text, "//") {
			continue
		} else {
			program[position] = fileScanner.Text()
			position++
		}
	}
	err = readFile.Close()
	if err != nil {
		return
	}
	for pos < len(program) {
		Execute(program[pos])
		pos++
	}
}
