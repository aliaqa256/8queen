package main

import "fmt"

const (
	SIZE = 8
)

var (
	board [SIZE][SIZE]int
	count = SIZE
)

var res [][SIZE][SIZE]int

func main() {
	PlaceQueen(0)

	for i := 0; i < len(res); i++ {
		PrintBoard(res[i])
	}

	fmt.Print("len is ")
	fmt.Println(len(res))
}

func PlaceQueen(row int) {
	if row == count {
		res = append(res, board)
		return
	}
	for col := 0; col < count; col++ {
		if IsValid(row, col) {
			board[row][col] = 1
			PlaceQueen(row + 1)
			board[row][col] = 0
		}
	}
}

func IsValid(row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 1 {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < count; i, j = i-1, j+1 {
		if board[i][j] == 1 {
			return false
		}
	}
	return true
}

func PrintBoard(board [SIZE][SIZE]int) {
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
