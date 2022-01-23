// By AureumApes

package commands

import (
	"strings"
)

func Execute(command string) {
	if strings.HasPrefix(command, "say") {
		Say(command)
	} else if strings.HasPrefix(command, "set") {
		Set(command)
	} else if strings.HasPrefix(command, "add") {
		Add(command)
	} else if strings.HasPrefix(command, "sub") {
		Sub(command)
	} else if strings.HasPrefix(command, "mult") {
		Mult(command)
	} else if strings.HasPrefix(command, "div") {
		Div(command)
	} else if strings.HasPrefix(command, "input") {
		Input(command)
	} else if strings.HasPrefix(command, "sqrt") {
		Sqrt(command)
	} else if strings.HasPrefix(command, "if") {
		If(command)
	}
}
