package main

import "fmt"

const (
	SIZE = 8
)

var (
	board [SIZE][SIZE]int
	count = SIZE
	res [][SIZE][SIZE]int
)


func main() {
	PlaceQueen(0)


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

	for i := 0; i < col; i++ {
		if board[row][i] == 1 {
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


func CheckColision() bool{
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if board[i][j] == 1 {
				if !IsValid(i, j) {
					fmt.Println("error---------------------------------------")
					return false
				}
			}
		}
	}
	return true
}

func CheckColisionWitharg(aa [SIZE][SIZE]int) bool{
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if aa[i][j] == 1 {
				if !IsValidWithargs(i, j,aa) {
					fmt.Println("error---------------------------------------")
					return false
				}
			}
		}
	}
	return true
}

func IsValidWithargs(row, col int,aa [SIZE][SIZE]int ) bool {
	for i := 0; i < row; i++ {
		if aa[i][col] == 1 {
			return false
		}
	}



	for i := 0; i < col; i++ {
		if aa[row][i] == 1 {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if aa[i][j] == 1 {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < count; i, j = i-1, j+1 {
		if aa[i][j] == 1 {
			return false
		}
	}
	return true
}
