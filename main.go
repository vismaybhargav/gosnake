package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram()
	fmt.Println("Hello World")
}

func InitModel() Model {
	return Game{}
}
