package main

import "fmt"

// 快慢指针

type ListNode struct {
	value int
	Next  *ListNode
}

func findMiddle(head *ListNode) *ListNode {
	var preNode *ListNode = head
	var Middle *ListNode = head
	for preNode != nil && preNode.Next != nil && preNode.Next.Next != nil {
		preNode = preNode.Next.Next
		Middle = Middle.Next
	}
	return Middle
}

func init() {
	l := &ListNode{
		value: 1,
		Next: &ListNode{
			value: 2,
			Next: &ListNode{
				value: 3,
				Next: &ListNode{
					value: 4,
					Next: &ListNode{
						value: 5,
						Next: &ListNode{
							value: 6,
							Next: &ListNode{
								value: 7,
							},
						},
					},
				},
			},
		},
	}
	l1 := findMiddle(l)
	fmt.Printf("%#v \n", l1)

}

func main() {

}
