package main

import (
	"fmt"
)


var count = 8

func main() {
	cols := []int {}
	

	SetQueen(0,cols)
}

func SetQueen(row int,cols []int) {
	if count == len(cols) {
		fmt.Println(cols)
		return
	}

	for col := 0; col < count; col++ {
		if isValid(cols, row, col) {
			SetQueen(row+1,append(cols, col))
		}
	}
}

func isValid(cols []int, row, col int) bool {
	for r, c := range cols {
		if c == col {
			return false
		}

		if r-c == row-col {
			return false
		}

		if r+c == row+col {
			return false
		}
	}

	return true
}
