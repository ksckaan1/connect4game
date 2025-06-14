package ports

import (
	"context"

	"github.com/ksckaan1/connect4backend/internal/core/domain"
)

type GamePort interface {
	CreateRoom(ctx context.Context) (string, error)
}

type GameEnginePort interface {
	Resolve(board [][]domain.Stone) (domain.Status, [][]int, error)
	GetTotalMove(board [][]domain.Stone) int
}

type GameSessionPort interface {
	GetTurn() domain.Turn
	GetBoard() [][]domain.Stone
	AddStone(col int) error
	GetAvailableCols() []int
}
