package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	var (
		m    = make(map[int]int)
		flag = len(nums) / 2
	)
	for index, _ := range nums {
		v := nums[index]
		m[v] += 1
		if m[v] > flag {
			return v
		}
	}
	return -1

}

func majorityElement1(nums []int) int {
	// 摩尔投票法
	var major = 0
	var num = 0

	for i := 0; i < len(nums); i++ {
		if num == 0 {
			major = nums[i]
		}
		if major == nums[i] {
			num++
		} else {
			num--
		}
	}

	var count = 0
	for j := 0; j < len(nums); j++ {
		if nums[j] == major {
			count++
		}
	}
	if count <= len(nums)/2 {
		major = -1
	}
	return major
}

func main() {
	fmt.Println(majorityElement1([]int{1, 2, 3, 1, 1}))
}
