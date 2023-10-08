package main

import (
	"fmt"
)

func main() {
	l3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}

	show(l1)
	show(reverseList(l1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	revHead := &ListNode{
		Val:  head.Val,
		Next: nil,
	}

	it := head
	for it.Next != nil {
		it = it.Next

		revHead = &ListNode{
			Val:  it.Val,
			Next: revHead,
		}
	}

	return revHead
}

func show(head *ListNode) {
	fmt.Print("[ ")
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println("]")
}
