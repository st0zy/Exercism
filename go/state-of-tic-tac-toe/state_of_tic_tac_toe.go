package stateoftictactoe

import (
	"fmt"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	err := validate(board)

	if err != nil {
		return "", err
	}

	numberOfTurnsForX := countSymbol(board, 'X')
	numberOfTurnsForO := countSymbol(board, 'O')

	// fmt.Println(numberOfTurnsForO, numberOfTurnsForX)

	if numberOfTurnsForO > numberOfTurnsForX {
		return "", fmt.Errorf("o started before x")
	}
	if numberOfTurnsForX-numberOfTurnsForO > 1 {
		return "", fmt.Errorf("incorrect orders of turns")
	}

	hasXWon, errX := hasPersonWon(board, 'X')
	if errX != nil {
		return "", errX
	}
	hasYWon, errO := hasPersonWon(board, 'O')
	if errO != nil {
		return "", errO
	}

	if hasXWon && hasYWon {
		return "", fmt.Errorf("continued after winning")
	}
	if hasXWon || hasYWon {
		return Win, nil
	}
	if numberOfTurnsForO+numberOfTurnsForX == 9 {
		return Draw, nil
	}
	return Ongoing, nil
}

func validate(board []string) error {
	rows := len(board)
	if rows != 3 {
		return fmt.Errorf("invalid number of rows")
	}

	for i, _ := range board {
		if len(board[i]) != 3 {
			return fmt.Errorf("invalid number of columns")
		}
	}

	for _, row := range board {
		for _, val := range row {
			if !(val == 'X' || val == 'O' || val == ' ') {
				return fmt.Errorf("unknown value in board")
			}
		}
	}

	return nil
}

func countSymbol(board []string, symbol byte) int {
	count := 0
	for i := range 3 {
		for j := range 3 {
			if board[i][j] == symbol {
				count++
			}
		}
	}

	return count
}

func hasPersonWon(board []string, symbol byte) (bool, error) {
	won := false
	for i := range 3 {
		if board[i][0] == symbol && board[i][1] == symbol && board[i][2] == symbol {
			if won {
				return false, fmt.Errorf("already continued after playing")
			}
			won = true
		}
	}
	if won {
		return won, nil
	}

	for j := range 3 {
		if board[0][j] == symbol && board[1][j] == symbol && board[2][j] == symbol {
			if won {
				return false, fmt.Errorf("already continued after playing")
			}
			won = true
		}
	}

	if board[0][0] == symbol && board[1][1] == symbol && board[2][2] == symbol {
		won = true
	}

	if board[0][2] == symbol && board[1][1] == symbol && board[2][0] == symbol {
		won = true
	}

	return won, nil
}
