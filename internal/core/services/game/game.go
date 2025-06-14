package game

import "github.com/ksckaan1/connect4backend/internal/core/ports"

type Game struct {
	ge ports.GameEnginePort
	gs ports.GameSessionPort
}

func New(ge ports.GameEnginePort, gs ports.GameSessionPort) *Game {
	return &Game{
		ge: ge,
		gs: gs,
	}
}
