package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	var recursion func(n *ListNode)
	recursion = func(n *ListNode) {
		if n == nil {
			return
		}
		if recursion(n.Next); head == nil {
			return
		} else if head == n || head.Next == n {
			head.Next = nil
		}
		head.Next, head, n.Next = n, head.Next, head.Next
	}
	recursion(head)
}
