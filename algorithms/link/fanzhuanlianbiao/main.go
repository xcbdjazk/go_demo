package main

import "fmt"

// 反转链表  3 > 2> 1  ------- 1<2<3

type ListNode struct {
	value int
	Next  *ListNode
}

func reversrList(head *ListNode) *ListNode {
	var preNode *ListNode
	preNode = nil
	//  后一个节点
	var nextNode *ListNode
	nextNode = nil
	for head != nil {
		//  保存头节点的下一个节点，
		nextNode = head.Next
		//  将头节点指向前一个节点
		head.Next = preNode
		//  更新前一个节点
		preNode = head
		//  更新头节点
		head = nextNode
	}
	return preNode
}

func init() {
	l := &ListNode{
		value: 1,
		Next: &ListNode{
			value: 2,
			Next: &ListNode{
				value: 3,
			},
		},
	}
	l1 := reversrList(l)
	a := new(ListNode)
	a = l1
	for a != nil {
		fmt.Printf("%#v \n", a)
		a = a.Next

	}
	fmt.Println(l1)

}

func main() {

}
