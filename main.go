package main

import "fmt"
var res [][]int

var count = 4
func main() {
	cols := []int {}
	res =[][]int{}
	

	SetQueen(0,cols)

	fmt.Println(res)
	fmt.Println(len(res))

	CreateTableFrom(res)
}

func SetQueen(row int,cols []int) {
	if count == len(cols) {
		res = append(res, cols)
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


func CreateTableFrom(res [][]int){
	table := [][]int {
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},

	}
	for _,resualt := range res {

		for col ,row := range resualt {
			table[col][row] = 1
		}
	


		for _,row := range table {
			fmt.Println(row)
		}

		fmt.Println("------------------------")
		table = [][]int {
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0},

	}
	}
}