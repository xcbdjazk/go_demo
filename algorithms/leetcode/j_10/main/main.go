package main

import (
	"fmt"
)

func minArray(n []int) int {
	var (
		i = 0
		j = len(n) - 1
	)

	for i < j {
		m := (i + j) / 2
		if n[m] > n[j] {
			i = m + 1
		} else if n[m] > n[i] {
			j = m
		} else {
			j--
		}
	}
	return n[i]
}

func main() {
	fmt.Println(minArray([]int{4, 5, 6, 3}))
}
