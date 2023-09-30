package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	mid := len(lists) / 2
	leftMerged := mergeKLists(lists[:mid])
	rightMerged := mergeKLists(lists[mid:])
	return mergeTwoLists(leftMerged, rightMerged)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res

	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			l2 = nil
			continue
		}
		if l2 == nil {
			cur.Next = l1
			l1 = nil
			continue
		}

		if l1.Val < l2.Val {
			cur.Next = l1
			l1, cur = l1.Next, cur.Next
			continue
		}

		cur.Next = l2
		l2, cur = l2.Next, cur.Next
	}

	return res.Next
}
