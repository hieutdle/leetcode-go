package algorithms

import "github.com/goldennovember/leetcode-go/gods"

type Stack = gods.Stack

func backspaceCompare(s string, t string) bool {

	sStack := Stack{}
	tStack := Stack{}

	for _, c := range s {
		if c == '#' {
			sStack.Pop()
		} else {
			sStack.Push(c)
		}
	}

	for _, c := range t {
		if c == '#' {
			tStack.Pop()
		} else {
			tStack.Push(c)
		}
	}

	if sStack.Size() != tStack.Size() {
		return false
	}

	for i := 0; i < sStack.Size(); i++ {
		if sStack.Items[i] != tStack.Items[i] {
			return false
		}
	}

	return true
}
