package main

import "log"

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}

	res := mergeTwoLists(l1, l2)
	log.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	res := head

	for {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				res.Next = &ListNode{
					Val: list1.Val,
				}
				list1 = list1.Next
			} else {
				res.Next = &ListNode{
					Val: list2.Val,
				}
				list2 = list2.Next
			}
			res = res.Next
		} else if list1 == nil && list2 != nil {
			res.Next = &ListNode{
				Val: list2.Val,
			}
			list2 = list2.Next
			res = res.Next
		} else if list1 != nil && list2 == nil {
			res.Next = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
			res = res.Next
		} else {
			break
		}
	}

	return head.Next
}
