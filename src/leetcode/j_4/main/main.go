package main

import "fmt"

/**
*
Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var b []int

	for head != nil {
		val := head.Val
		b = append([]int{val}, b...)
		head = head.Next
	}
	return b
}

func main() {
	fmt.Println(reversePrint(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}))
}
