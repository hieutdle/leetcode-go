package gods

import "fmt"

type ListNode struct {
	Val int
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