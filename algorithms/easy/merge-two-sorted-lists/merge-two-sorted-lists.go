package algorithms

import (
	"github.com/goldennovember/leetcode-go/gods"
)

type LinkedList = gods.LinkedList
type ListNode = gods.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	ans := LinkedList{}
	var node *ListNode

	if list1.Val > list2.Val {
		ans.Head = list2
		node = list2
		list2 = list2.Next
	} else {
		ans.Head = list1
		node = list1
		list1 = list1.Next
	}

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			node.Next = list2
			list2 = list2.Next
		} else {
			node.Next = list1
			list1 = list1.Next
		}
		node = node.Next
	}

	if list1 != nil {
		node.Next = list1
	}

	if list2 != nil {
		node.Next = list2
	}

	return ans.Head
}
