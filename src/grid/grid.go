package grid

type Grid struct {
	width  int
	height int
	board  [][]int
}

func (g Grid) New(width, height int) Grid {
	return Grid{
		width:  width,
		height: height,
		board:  CreateGridBoard(width, height),
	}
}

func CreateGridBoard(width, height int) (board [][]int) {
	board = make([][]int, height)

	for i := 0; i < height; i++ {
		board[i] = make([]int, width)
	}

	return
}
