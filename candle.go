package main

import (
	"bufio"
	"github.com/aureumapes/candle/internal"
	"os"
	"strings"
)

var vars = make(map[string]string)
var program = []string{}
var funcs = make(map[string][]string)
var pos int

func main() {
	args := os.Args[1:]
	if args[0] == "version"{
		println(version)
		os.Exit(0);
	}
	if len(args) < 1 {
		panic("Please make sure to use a Filename as an argument!")
	}
	file, err := os.Open(args[0])
	if err != nil {
		panic(err.Error())
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		command := fileScanner.Text()
		if strings.HasPrefix(command, "//") || command == "" {
			continue
		} else {
			program = append(program, command)
		}
	}
	defer file.Close()

	pos = 0
	for pos < len(program) {
		newProgram := strings.Split(strings.Join(program, "\n"), program[pos])[1]
		vars, funcs = internal.Execute(program[pos], vars, funcs, strings.Split(newProgram, "\n"))
		pos++
	}
}
