package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev, curr := dummy, head

	for curr != nil && curr.Next != nil {
		// save pointers
		nextPair := curr.Next.Next
		second := curr.Next

		// swap
		second.Next = curr
		curr.Next = nextPair
		prev.Next = second

		// update pointers
		prev = curr
		curr = nextPair
	}

	return dummy.Next
}
