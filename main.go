package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

const WIDTH = 10
const HEIGHT = 10

func main() {
	fmt.Println("Hello World")
}

func InitModel() *Game {
	return &Game{
		grid:      Grid{}.New(WIDTH, HEIGHT),
		appleX:    2,
		appleY:    2,
		snakeX:    make([]int, WIDTH*HEIGHT),
		snakeY:    make([]int, WIDTH*HEIGHT),
		isRunning: false,
	}
}

type Game struct {
	grid      *Grid
	appleX    int
	appleY    int
	snakeX    []int
	snakeY    []int
	isRunning bool
}
