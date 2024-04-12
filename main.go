package main

import (
	"fmt"
)

type field struct {
	value   int
	isGiven bool
}

var numOfRows int

func checkVerticalAndHroizontally(sudoku []field, index int, num int) bool { //działa

	for x := range numOfRows {
		if sudoku[(numOfRows*x)+(index%numOfRows)].value == num || sudoku[(index/numOfRows)*numOfRows+x].value == num {
			return false
		}
	}

	return true
}

func checkBox(sudoku []field, index int, num int) bool {
	idx := index / numOfRows
	idx = (idx / 3) * 3
	position := ((index % numOfRows) / 3) * 3
	for x := range numOfRows / 3 {
		for y := range numOfRows / 3 {
			if sudoku[((idx+x)*numOfRows)+position+y].value == num {
				return false
			}
		}
	}
	return true
}

func checkCorectnes(sudoku []field, index int, num int) bool {
	v1 := checkVerticalAndHroizontally(sudoku, index, num)
	v2 := checkBox(sudoku, index, num)
	return v1 && v2

}

func readData(sudoku []field) {
	for m := range numOfRows {
		var row string
		fmt.Scanln(&row)
		for x, l := range row {
			if l == '.' {
				sudoku[numOfRows*m+x].isGiven = false
				sudoku[numOfRows*m+x].value = 0
			} else {
				sudoku[numOfRows*m+x].isGiven = true
				sudoku[numOfRows*m+x].value = int(l) - '0'
			}
		}
	}
}

func printSudoku(sudoku []field) {
	for m := range numOfRows {
		for y := range numOfRows {
			fmt.Printf("%d", sudoku[numOfRows*m+y].value)
		}
		fmt.Println()
	}
	println()
}

func isGit(sudoku []field) bool {
	for x := range numOfRows * numOfRows {
		value := sudoku[x].value
		sudoku[x].value = 0
		if !checkCorectnes(sudoku, x, value) {
			return false
		}
		sudoku[x].value = value
	}
	return true
}

func main() {
	fmt.Scanf("%d", &numOfRows)
	sudoku := make([]field, (numOfRows)*(numOfRows))
	readData(sudoku)
	i := 0
	forward := true
	for i >= 0 && i < numOfRows*numOfRows {
		//fmt.Printf("i := %d\n", i)
		if sudoku[i].isGiven {
			if forward {
				i++
			} else {
				i--
			}
			continue
		}
		candidate := 1
		if forward {
			candidate = 1
		} else {
			candidate = sudoku[i].value + 1
		}
		for ; candidate <= 10; candidate++ {
			if candidate == 10 {
				sudoku[i].value = 0
				i -= 1
				forward = false
				break
			}
			if checkCorectnes(sudoku, i, candidate) { // rozwiązac zagadkę
				sudoku[i].value = candidate
				forward = true
				i += 1
				printSudoku(sudoku)
				break
			}

		}
	}
	if i == -1 {
		fmt.Printf("nie udało się :(\n %v", sudoku)
	} else {
		if isGit(sudoku) {
			fmt.Printf("udało sie fajnie co nie\n")
			printSudoku(sudoku)
		} else {
			fmt.Printf("nie udało się :(")
		}
	}
}
