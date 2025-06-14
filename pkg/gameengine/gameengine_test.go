package gameengine

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/ksckaan1/connect4backend/internal/core/domain"
)

func TestGameEngineGetTotalMove(t *testing.T) {
	ge := New()

	board := [][]Stone{
		{X, O, X},
		{O, X, O},
		{O, X, O},
		{},
		{},
		{},
		{},
	}

	totalMove := ge.GetTotalMove(board)

	require.Equal(t, 9, totalMove)
}

func TestGameEngineResolve(t *testing.T) {
	ge := New()

	tests := []struct {
		name             string
		board            [][]Stone
		expectedStatus   Status
		expectedSequence [][]int
		expectedError    error
	}{
		{
			name: "XWinVertical",
			board: [][]Stone{
				{O, O, X, X, X},
				{O, X, X, E, E},
				{X, X, X, E, E},
				{O, X, X, E, E},
				{O, O, O, E, E},
			},
			expectedStatus:   XWin,
			expectedSequence: [][]int{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
			expectedError:    nil,
		},
		{
			name: "XWinHorizontal",
			board: [][]Stone{
				{O, O, X, X, X},
				{O, X, X, X, X},
				{X, X, O, E, E},
				{O, X, X, E, E},
			},
			expectedStatus:   XWin,
			expectedSequence: [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}},
			expectedError:    nil,
		},
		{
			name: "XWinTopLeftToBottomRight",
			board: [][]Stone{
				{O, X, X, X, O},
				{O, X, X, X, O},
				{X, O, O, X, E},
				{O, X, X, O, X},
			},
			expectedStatus:   XWin,
			expectedSequence: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			expectedError:    nil,
		},
		{
			name: "XWinTopRightToBottomLeft",
			board: [][]Stone{
				{O, X, X, X, O},
				{O, X, X, X, O},
				{X, X, O, O, E},
				{X, O, X, E, E},
			},
			expectedStatus:   XWin,
			expectedSequence: [][]int{{0, 3}, {1, 2}, {2, 1}, {3, 0}},
			expectedError:    nil,
		},
		{
			name: "OWinVertical",
			board: [][]Stone{
				{O, O, O, X, X},
				{O, X, O, X, X},
				{X, X, O, E, E},
				{O, X, O, E, E},
			},
			expectedStatus:   OWin,
			expectedSequence: [][]int{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
			expectedError:    nil,
		},
		{
			name: "OWinHorizontal",
			board: [][]Stone{
				{O, O, O, X, X},
				{O, X, O, X, X},
				{X, O, O, O, O},
				{O, X, X, E, E},
			},
			expectedStatus:   OWin,
			expectedSequence: [][]int{{2, 1}, {2, 2}, {2, 3}, {2, 4}},
			expectedError:    nil,
		},
		{
			name: "OWinTopRightToBottomLeft",
			board: [][]Stone{
				{O, O, O, X, O},
				{O, X, O, O, X},
				{X, O, O, X, O},
				{O, O, X, E, E},
			},
			expectedStatus:   OWin,
			expectedSequence: [][]int{{0, 4}, {1, 3}, {2, 2}, {3, 1}},
			expectedError:    nil,
		},
		{
			name: "OWinTopLeftToBottomRight",
			board: [][]Stone{
				{O, O, O, E, E},
				{O, X, O, O, E},
				{X, O, X, O, E},
				{O, X, X, X, O},
			},
			expectedStatus:   OWin,
			expectedSequence: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			expectedError:    nil,
		},
		{
			name: "Draw",
			board: [][]Stone{
				{O, O, O, X, X},
				{O, X, O, O, X},
				{X, O, X, X, X},
				{O, X, O, X, O},
			},
			expectedStatus:   Draw,
			expectedSequence: nil,
			expectedError:    nil,
		},
		{
			name: "Playing",
			board: [][]Stone{
				{O, O, O, X, X},
				{O, X, O, O, X},
				{X, O, X, X, E},
				{O, X, O, X, O},
			},
			expectedStatus:   Playing,
			expectedSequence: nil,
			expectedError:    nil,
		},
		{
			name: "InvalidSequence",
			board: [][]Stone{
				{O, O, O, X, X},
				{O, X, E, O, X},
				{X, O, X, X, E},
				{O, X, O, X, O},
			},
			expectedStatus:   0,
			expectedSequence: nil,
			expectedError:    ErrInvalidSquence,
		},
		{
			name: "InvalidBoard",
			board: [][]Stone{
				{O, O, O, X, X},
				{O, X, X, O, X},
				{X, O, X, X},
				{O, X, O, X, O},
			},
			expectedStatus:   0,
			expectedSequence: nil,
			expectedError:    ErrInvalidBoard,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, gotSequence, gotErr := ge.Resolve(tt.board)
			require.Equal(t, tt.expectedStatus, gotStatus)
			require.Equal(t, tt.expectedSequence, gotSequence)
			require.ErrorIs(t, gotErr, tt.expectedError)
		})
	}
}
