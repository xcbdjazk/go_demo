package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	var (
		lenA = len(matrix)
		lenB = len(matrix[0])
		row  = 0
		clo  = lenB - 1
	)
	for row < lenA && clo >= 0 {
		num := matrix[row][clo]
		if num == target {
			return true
		} else if num > target {
			clo--
		} else {
			row++
		}
	}

	return false
}

func main() {
	fmt.Println(findNumberIn2DArray([][]int{
		[]int{1, 2, 3},
		[]int{2, 3, 4},
	}, 11))
}
