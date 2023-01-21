package gods

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head   *ListNode
	Length int
}

func (l *LinkedList) Prepend(n *ListNode) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.Length++
}

func (l LinkedList) PrintListVal() {
	toPrint := l.Head
	for l.Length != 0 {
		fmt.Printf("%d ", toPrint.Val)
		toPrint = toPrint.Next
		l.Length--
	}
	fmt.Println("\n")
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.Length == 0 {
		return
	}
	if l.Head.Val == value {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	previousToDelete := l.Head
	for previousToDelete.Next.Val != value {
		if previousToDelete.Next.Next == nil {
			return
		}
		previousToDelete = previousToDelete.Next
	}
	previousToDelete.Next = previousToDelete.Next.Next
	l.Length--
}

//

func (l *LinkedList) Append(n *ListNode) {
	if l.Length == 0 {
		l.Head = n
		l.Length++
		return
	}
	toAppend := l.Head
	for toAppend.Next != nil {
		toAppend = toAppend.Next
	}
	toAppend.Next = n
	l.Length++
}

func (l *LinkedList) Insert(n *ListNode, position int) bool {
	if position > l.Length {
		return false
	}
	if position == 1 {
		l.Prepend(n)
		return true
	}
	if position == l.Length {
		l.Append(n)
		return true
	}
	toInsert := l.Head
	for i := 1; i < position-1; i++ {
		toInsert = toInsert.Next
	}
	n.Next = toInsert.Next
	toInsert.Next = n
	l.Length++
	return true
}

func (l *LinkedList) Delete(position int) bool {
	if position > l.Length {
		return false
	}
	if position == 1 {
		l.Head = l.Head.Next
		l.Length--
		return true
	}
	toDelete := l.Head
	for i := 1; i < position-1; i++ {
		toDelete = toDelete.Next
	}
	toDelete.Next = toDelete.Next.Next
	l.Length--
	return true
}

func List2Slices(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func Slices2List(nums []int) *ListNode {
	var head *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		head = &ListNode{Val: nums[i], Next: head}
	}
	return head
}
