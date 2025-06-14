package gamesession

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/ksckaan1/connect4backend/internal/core/domain"
)

func TestAddStone(t *testing.T) {
	tests := []struct {
		name          string
		gs            *GameSession
		col           int
		times         int
		check         func(*GameSession)
		expectedError error
	}{
		{
			name: "EmptyCol",
			gs:   New().WithSize(1, 6).WithTurn(OTurn),
			col:  0,
			check: func(gs *GameSession) {
				require.Equal(t, O, gs.GetBoard()[0][0])
				require.Equal(t, XTurn, gs.GetTurn())
			},
			expectedError: nil,
		},
		{
			name: "AvailableCol",
			gs: New().
				WithTurn(OTurn).
				WithBoard([][]Stone{
					{X, O, X, O, X, E},
				}),
			col: 0,
			check: func(gs *GameSession) {
				require.Equal(t, O, gs.GetBoard()[0][5])
				require.Equal(t, XTurn, gs.GetTurn())
			},
			expectedError: nil,
		},
		{
			name: "FullCol",
			gs: New().
				WithTurn(XTurn).
				WithBoard([][]Stone{
					{X, O, X, O, X, O},
				}),
			check: func(gs *GameSession) {
				require.Len(t, gs.GetBoard()[0], 6)
				require.Equal(t, XTurn, gs.GetTurn())
			},
			col:           0,
			expectedError: ErrColNotAvailable,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.times; i++ {
				err := tt.gs.AddStone(tt.col)
				require.Equal(t, tt.expectedError, err)
				if tt.check != nil {
					tt.check(tt.gs)
				}
			}
		})
	}
}

func TestGetAvailableCols(t *testing.T) {
	tests := []struct {
		name         string
		gs           *GameSession
		expectedCols []int
	}{
		{
			name:         "Example1",
			gs:           New().WithSize(7, 6),
			expectedCols: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "Example2",
			gs: New().
				WithBoard([][]Stone{
					{X, O, X, O, X, O},
					make([]Stone, 6),
				}),
			expectedCols: []int{1},
		},
		{
			name: "Example3",
			gs: New().
				WithBoard([][]Stone{
					{X, O, X, O, X, O},
					{X, O, X, O, X, O},
					{X, O, X, O, X, O},
					{X, O, X, O, X, O},
				}),
			expectedCols: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectedCols, tt.gs.GetAvailableCols())
		})
	}
}

func TestGetTurn(t *testing.T) {
	tests := []struct {
		name         string
		gs           *GameSession
		expectedTurn Turn
	}{
		{
			name:         "XTurn",
			gs:           New().WithTurn(XTurn),
			expectedTurn: XTurn,
		},
		{
			name:         "OTurn",
			gs:           New().WithTurn(OTurn),
			expectedTurn: OTurn,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectedTurn, tt.gs.GetTurn())
		})
	}
}
