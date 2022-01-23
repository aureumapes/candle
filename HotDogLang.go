package main

import (
	"bufio"
	"fmt"
	"os"
)
import _ "./commands"
import "./commands"

func main() {
	args := os.Args[1:]
	readFile, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		commands.Execute(fileScanner.Text())
	}
	readFile.Close()
}
