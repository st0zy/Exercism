package minesweeper

import (
	"bytes"
	"strconv"
)

// Annotate returns an annotated board
func Annotate(board []string) []string {
	if len(board) == 0 {
		return board
	}
	if len(board[0]) == 0 {
		return board
	}
	TopDownAnnotate(board)
	BottomUpAnnotate(board)

	return board
}

func TopDownAnnotate(board []string) []string {

	rows := len(board)
	cols := len(board[0])

	var value int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			value = 0
			if board[i][j] == '*' {
				continue
			}
			if i-1 >= 0 && isMine(board[i-1][j]) {
				value++
			}
			if j-1 >= 0 && isMine(board[i][j-1]) {
				value++
			}
			if i-1 >= 0 && j-1 >= 0 && isMine(board[i-1][j-1]) {
				value++
			}
			if i+1 < rows && j-1 >= 0 && isMine(board[i+1][j-1]) {
				value++
			}
			if value != 0 {
				updateBoard(board, i, j, value)
			}
		}

	}

	return board
}

func updateBoard(board []string, i int, j int, value int) {

	var buffer bytes.Buffer
	buffer.WriteString(board[i][:j])
	currentValue, _ := strconv.Atoi(string(board[i][j]))
	currentValue += value
	buffer.WriteString(strconv.Itoa(currentValue))
	if j != len(board[0])-1 {
		buffer.WriteString(board[i][j+1:])
	}
	board[i] = buffer.String()
}

func BottomUpAnnotate(board []string) []string {
	rows := len(board)
	cols := len(board[0])

	var value int

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			value = 0
			if board[i][j] == '*' {
				continue
			}
			if i+1 < rows && isMine(board[i+1][j]) {
				value++
			}
			if j+1 < cols && isMine(board[i][j+1]) {
				value++
			}
			if i+1 < rows && j+1 < cols && isMine(board[i+1][j+1]) {
				value++
			}
			if i-1 >= 0 && j+1 < cols && isMine(board[i-1][j+1]) {
				value++
			}
			if value != 0 {
				updateBoard(board, i, j, value)
			}
		}
	}
	return board
}

func isMine(item byte) bool {
	return item == byte('*')
}
