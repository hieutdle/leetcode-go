package merge_two_sorted_lists

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

	var ans *LinkedList

	if list1.Val > list2.Val {
		ans = &LinkedList{Head: list2}
		list2 = list2.Next
	} else {
		ans = &LinkedList{Head: list1}
		list1 = list1.Next
	}

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			ans.Prepend(list2)
			list2 = list2.Next
		} else {
			ans.Prepend(list1)
			list1 = list1.Next
		}
	}

	return ans.Head
}
