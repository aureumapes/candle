package main

import (
	"github.com/charmbracelet/lipgloss"
	"os"
	"slices"
	"strings"
)

func main() {
	file, _ := os.ReadFile(os.Args[1])
	keywordsStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#5f9cc5")).Bold(true)
	stringsStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#c45e5e"))

	keywords := []string{"set", "say", "sayln", "ask", "read", "add", "sub", "mul", "div", "if", "function", "import", "to", "end"}

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		for wi, word := range words {
			if slices.Contains(keywords, word) {
				print(keywordsStyle.Render(word) + " ")
			} else if strings.HasPrefix(word, "+") {
				for _, stringWord := range words[wi:] {
					print(stringsStyle.Render(stringWord) + " ")
				}
				break
			} else if word != "\n" {
				print(word + " ")
			} else {
				println("")
			}
		}
		print("\n")
	}
}
