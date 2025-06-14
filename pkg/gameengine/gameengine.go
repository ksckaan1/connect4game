package gameengine

import "github.com/ksckaan1/connect4backend/internal/core/domain"

type GameEngine struct {
}

func New() *GameEngine {
	return &GameEngine{}
}

func (*GameEngine) GetTotalMove(board [][]domain.Stone) int {
	totalMove := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] != domain.E {
				totalMove++
			}
		}
	}
	return totalMove
}

func (g *GameEngine) Resolve(board [][]domain.Stone) (s domain.Status, coords [][]int, err error) {
	cols, rows, err := g.recognizeBoard(board)
	if err != nil {
		return 0, nil, err
	}

	if g.GetTotalMove(board) < 7 {
		return domain.Playing, nil, nil
	}

	// check horizontally
	s, coords = g.checkHorizontally(board, rows, cols)
	if s != domain.Playing {
		return s, coords, nil
	}

	// Check vertically
	s, coords = g.checkVertically(board, rows, cols)
	if s != domain.Playing {
		return s, coords, nil
	}

	// Check top-left to bottom-right
	s, coords = g.checkTopLeftToBottomRight(board, rows, cols)
	if s != domain.Playing {
		return s, coords, nil
	}

	// Check top-right to bottom-left
	s, coords = g.checkTopRightToBottomLeft(board, rows, cols)
	if s != domain.Playing {
		return s, coords, nil
	}

	// Is it playing now?
	if g.isItPlayingNow(board, rows, cols) {
		return domain.Playing, nil, nil
	}

	return domain.Draw, nil, nil
}

func (*GameEngine) getWinnerFromStone(s domain.Stone) domain.Status {
	switch s {
	case domain.X:
		return domain.XWin
	case domain.O:
		return domain.OWin
	default:
		return domain.Draw
	}
}

func (*GameEngine) recognizeBoard(board [][]domain.Stone) (cols, rows int, err error) {
	rows = len(board)
	if rows < 4 {
		return 0, 0, domain.ErrInvalidBoard
	}
	cols = len(board[0])
	for i := range board {
		if cols != len(board[i]) {
			return 0, 0, domain.ErrInvalidBoard
		}
	}

	for i := 0; i < rows; i++ {
		eFound := false // empty stone
		for j := 0; j < cols; j++ {
			if board[i][j] == domain.E {
				eFound = true
			}
			if eFound && board[i][j] != domain.E {
				return 0, 0, domain.ErrInvalidSquence
			}
		}
	}

	return cols, rows, nil
}

func (g *GameEngine) checkHorizontally(board [][]domain.Stone, rows, cols int) (domain.Status, [][]int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols-3; j++ {
			if board[i][j] != domain.E && board[i][j] == board[i][j+1] && board[i][j] == board[i][j+2] && board[i][j] == board[i][j+3] {
				return g.getWinnerFromStone(board[i][j]), [][]int{{i, j}, {i, j + 1}, {i, j + 2}, {i, j + 3}}
			}
		}
	}
	return domain.Playing, nil
}

func (g *GameEngine) checkVertically(board [][]domain.Stone, rows, cols int) (domain.Status, [][]int) {
	for i := 0; i < rows-3; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] != domain.E && board[i][j] == board[i+1][j] && board[i][j] == board[i+2][j] && board[i][j] == board[i+3][j] {
				return g.getWinnerFromStone(board[i][j]), [][]int{{i, j}, {i + 1, j}, {i + 2, j}, {i + 3, j}}
			}
		}
	}
	return domain.Playing, nil
}

func (g *GameEngine) checkTopLeftToBottomRight(board [][]domain.Stone, rows, cols int) (domain.Status, [][]int) {
	for i := 0; i < rows-3; i++ {
		for j := 0; j < cols-3; j++ {
			if board[i][j] != domain.E && board[i][j] == board[i+1][j+1] && board[i][j] == board[i+2][j+2] && board[i][j] == board[i+3][j+3] {
				return g.getWinnerFromStone(board[i][j]), [][]int{{i, j}, {i + 1, j + 1}, {i + 2, j + 2}, {i + 3, j + 3}}
			}
		}
	}
	return domain.Playing, nil
}

func (g *GameEngine) checkTopRightToBottomLeft(board [][]domain.Stone, rows, cols int) (domain.Status, [][]int) {
	for i := 0; i < rows-3; i++ {
		for j := 3; j < cols; j++ {
			if board[i][j] != domain.E && board[i][j] == board[i+1][j-1] && board[i][j] == board[i+2][j-2] && board[i][j] == board[i+3][j-3] {
				return g.getWinnerFromStone(board[i][j]), [][]int{{i, j}, {i + 1, j - 1}, {i + 2, j - 2}, {i + 3, j - 3}}
			}
		}
	}
	return domain.Playing, nil
}

func (g *GameEngine) isItPlayingNow(board [][]domain.Stone, rows, cols int) bool {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == domain.E {
				return true
			}
		}
	}
	return false
}
