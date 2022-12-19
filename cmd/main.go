package main

import (
	"Candle/internal"

	"bufio"
	"os"
	"strings"
)

var vars = make(map[string]string)
var program = []string{}
var funcs = make(map[string][]string)
var pos int

func main() {
	args := os.Args[1:]
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
		Execute(program[pos])
		pos++
	}
}

func Execute(rawCommand string) {
	if rawCommand == "exit" {
		os.Exit(1)
	}
	command := strings.Split(rawCommand, " ")
	switch command[0] {
	case "say":
		internal.Say(command[1:], vars)
	case "sayln":
		internal.Sayln(command[1:], vars)
	case "set":
		vars = internal.Set(command[1:], vars)
	case "ask":
		vars = internal.Ask(command[1:], vars)
	case "add":
		vars = internal.Add(command[1:], vars)
	case "sub":
		vars = internal.Sub(command[1:], vars)
	case "mul":
		vars = internal.Mul(command[1:], vars)
	case "div":
		vars = internal.Div(command[1:], vars)
	}
}
