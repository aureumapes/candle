package commands

import (
	"strings"
)

func Execute(command string) {
	if strings.HasPrefix(command, "say") {
		Say(command)
	} else if strings.HasPrefix(command, "set") {
		Set(command)
	}
}
