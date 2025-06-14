package gamesession

import (
	"slices"

	"github.com/ksckaan1/connect4backend/internal/core/domain"
)

type GameSession struct {
	board [][]domain.Stone
	turn  domain.Turn
}

func New() *GameSession {
	return &GameSession{
		board: newBoard(7, 6),
		turn:  domain.GetRandomTurn(),
	}
}

func (g *GameSession) WithSize(w, h int) *GameSession {
	g.board = newBoard(w, h)
	return g
}

func (g *GameSession) WithBoard(b [][]domain.Stone) *GameSession {
	g.board = b
	return g
}

func (g *GameSession) WithTurn(t domain.Turn) *GameSession {
	if t == domain.RandomTurn {
		t = domain.GetRandomTurn()
	}
	g.turn = t
	return g
}

func (g *GameSession) GetTurn() domain.Turn {
	return g.turn
}

func (g *GameSession) GetBoard() [][]domain.Stone {
	return g.board
}

func (g *GameSession) AddStone(col int) error {
	if len(g.board) <= col || col < 0 {
		return domain.ErrInvalidCol
	}

	if !slices.Contains(g.GetAvailableCols(), col) {
		return domain.ErrColNotAvailable
	}

	switch g.GetTurn() {
	case domain.XTurn:
		g.append(col, domain.X)
		g.turn = domain.OTurn
	default:
		g.append(col, domain.O)
		g.turn = domain.XTurn
	}

	return nil
}

func (g *GameSession) GetAvailableCols() []int {
	ac := make([]int, 0, len(g.board))
	for i := range g.board {
		if slices.Contains(g.board[i], domain.E) {
			ac = append(ac, i)
		}
	}

	return ac
}

func (g *GameSession) append(c int, s domain.Stone) {
	for i := range g.board[c] {
		if g.board[c][i] == domain.E {
			g.board[c][i] = s
			return
		}
	}
}

func newBoard(w, h int) [][]domain.Stone {
	b := make([][]domain.Stone, w)
	for i := range b {
		b[i] = make([]domain.Stone, h)
	}
	return b
}
