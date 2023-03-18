package main

import "fmt"

// 链表时候有环

type ListNode struct {
	value int
	Next  *ListNode
}

func findMiddle(head *ListNode) (b bool) {
	var preNode *ListNode = head
	var Middle *ListNode = head
	for preNode != nil && preNode.Next != nil && preNode.Next.Next != nil {
		preNode = preNode.Next.Next
		Middle = Middle.Next
		if Middle == preNode {
			return true
		}
	}
	return false
}

func init() {
	ll := &ListNode{
		value: 7,
	}
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
								Next:  ll,
							},
						},
					},
				},
			},
		},
	}
	ll.Next = l
	l1 := findMiddle(l)
	fmt.Printf("%#v \n", l1)

}

func main() {

}
