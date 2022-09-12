package main

import "fmt"

func findRepeatNumber(nums []int) int {

	var arr = make([]int, len(nums))

	for index, _ := range nums {
		val := nums[index]
		arr[val] += 1
		if arr[val] > 1 {
			return val
		}
	}
	return -1
}

func main() {
	fmt.Println(findRepeatNumber([]int{1, 2, 3, 1}))
}
