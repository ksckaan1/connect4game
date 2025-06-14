package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ksckaan1/connect4backend/internal/core/domain"
	"github.com/ksckaan1/connect4backend/pkg/gameengine"
	"github.com/ksckaan1/connect4backend/pkg/gamesession"
)

func main() {
	ge := gameengine.New()
	gs := gamesession.New().
		WithSize(7, 6).
		WithTurn(domain.XTurn)

	for {
		fmt.Print("\033[H\033[2J")
		printBoard(gs.GetBoard())
		targetCol := -1
		fmt.Printf("Player %s choose column: ", gs.GetTurn())
		_, err := fmt.Scanf("%d", &targetCol)
		fmt.Println("")
		if err != nil {
			fmt.Println("invalid input")
			continue
		}

		err = gs.AddStone(targetCol)
		if err != nil {
			fmt.Println(err)
			continue
		}

		st, _, err := ge.Resolve(gs.GetBoard())
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch st {
		case domain.OWin:
			fmt.Println("Player O Win!")
			os.Exit(0)
		case domain.XWin:
			fmt.Println("Player X Win!")
			os.Exit(0)
		case domain.Draw:
			fmt.Println("Draw!")
			os.Exit(0)
		}
	}
}

func printBoard(board [][]domain.Stone) {
	rows := make([]string, 0, len(board[0]))

	for i := 0; i < len(board[0]); i++ {
		var row string
		for j := 0; j < len(board); j++ {
			if board[j][i] == domain.E {
				row += " "
				continue
			}
			row += string(board[j][i])
		}
		rows = append(rows, "|"+strings.Join(strings.Split(row, ""), "|")+"|")
	}

	slices.Reverse[[]string](rows)

	var numbers string
	for i := 0; i < len(board); i++ {
		numbers += fmt.Sprintf(" %d", i)
	}
	fmt.Println(numbers)

	dv := divider(len(board))
	fmt.Println(dv)
	for i := range rows {
		fmt.Println(rows[i])
		fmt.Println(dv)
	}

	fmt.Println("----------------------------")
}

func divider(cols int) string {
	return strings.Repeat("+-", cols) + "+"
}
