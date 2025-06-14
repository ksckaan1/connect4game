package domain

import (
	"math/rand"
)

type Turn int

const (
	XTurn Turn = iota
	OTurn
	RandomTurn
)

func (t Turn) String() string {
	switch t {
	case XTurn:
		return "X"
	case OTurn:
		return "O"
	case RandomTurn:
		return "Random"
	}
	return "Unknown"
}

func GetRandomTurn() Turn {
	return Turn(rand.Intn(2))
}
