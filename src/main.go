package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vismaybhargav/gosnake/src/grid"
	"log"
)

const WIDTH = 10
const HEIGHT = 10

func main() {
	p := tea.NewProgram(InitModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitModel() Game {
	return Game{
		grid:      grid.Grid{}.New(WIDTH, HEIGHT),
		appleX:    2,
		appleY:    2,
		snake:     Snake{}.New(WIDTH, HEIGHT),
		isRunning: false,
	}
}

type Game struct {
	grid      grid.Grid
	appleX    int
	appleY    int
	snake     Snake
	isRunning bool
}

type Snake struct {
	snakeX []int
	snakeY []int
	dir    string
}

func (s Snake) UpdateLocations() {
	dx, dy := 0, 0

	switch s.dir {
	case "up":
		dy = -1
	case "down":
		dy = 1
	case "left":
		dx = -1
	case "right":
		dx = 1
	}

	for i := range s.snakeX {
		s.snakeX[i] += dx
	}

	for i := range s.snakeY {
		s.snakeY[i] += dy
	}
}

func (s Snake) New(width, height int) Snake {
	return Snake{
		snakeX: make([]int, width*height),
		snakeY: make([]int, width*height),
		dir:    "",
	}
}

func (g Game) Init() tea.Cmd {
	return nil
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return g, tea.Quit
		case "up", "k":
			if g.snake.dir != "down" {
				g.snake.dir = "up"
			}
		case "down", "j":
			if g.snake.dir != "up" {
				g.snake.dir = "down"
			}
		case "left", "h":
			if g.snake.dir != "right" {
				g.snake.dir = "left"
			}
		case "right", "l":
			if g.snake.dir != "left" {
				g.snake.dir = "right"
			}
		}
	}

	g.snake.UpdateLocations()

	return g, nil
}

func (g Game) View() string {
	return g.snake.dir
}
