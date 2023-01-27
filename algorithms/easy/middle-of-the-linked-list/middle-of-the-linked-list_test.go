package algortihms

import (
	"fmt"
	"github.com/goldennovember/leetcode-go/gods"
	"github.com/stretchr/testify/assert"
	"testing"
)

// testcases
var tcs = []struct {
	head []int
	ans  []int
}{

	{
		[]int{1, 2, 3, 4, 5},
		[]int{3, 4, 5},
	},

	{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{4, 5, 6},
	},
}

func Test_middleNode(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		head := gods.Slices2List(tc.head)
		actual := gods.List2Slices(middleNode(head))
		ast.Equal(tc.ans, actual, "Case:%v", tc)
	}

	ast := assert.New(t)

	fmt.Printf("------------------------Leetcode Problem 206. Reverse Linked List------------------------\n")
	for _, tc := range tcs {
		fmt.Printf("【Input】: head = %v \n【Output】: %v \n", tc.head, gods.List2Slices(reverseList(gods.Slices2List(tc.head))))
		ast.Equal(tc.ans, gods.List2Slices(reverseList(gods.Slices2List(tc.head))), "Case: %v", tc)
	}
	fmt.Printf("\n\n\n")
}

func Benchmark_middleNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := gods.Slices2List(tc.head)
			middleNode(head)
		}
	}
}
