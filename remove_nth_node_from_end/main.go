package main

import (
	"fmt"
)

func main() {
	n5 := &ListNode{
		Val:  5,
		Next: nil,
	}

	n4 := &ListNode{
		Val:  4,
		Next: n5,
	}

	n3 := &ListNode{
		Val:  3,
		Next: n4,
	}

	n2 := &ListNode{
		Val:  2,
		Next: n3,
	}

	head := &ListNode{
		Val:  1,
		Next: n2,
	}

	input := head
	target := 2
	show(head)
	show(removeNthFromEnd(input, target))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	left := dummy
	right := head
	for i := 0; i < n; i++ {
		right = right.Next
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next
	return dummy.Next
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
	for head.Next != nil {
		fmt.Printf("{%d} -> ", head.Val)
		head = head.Next
	}
	fmt.Printf("{%d}\n", head.Val)
}
