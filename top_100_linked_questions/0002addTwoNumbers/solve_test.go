package _002addTwoNumbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//https://leetcode.com/problems/two-sum/
func Test_addTwoNumbers(t *testing.T) {
	t.Run("example1", func(t *testing.T) {
		l1 := makeNodeFromSlice([]int{2, 4, 3})
		l2 := makeNodeFromSlice([]int{5, 6, 4})
		assert.EqualValues(t, []int{7, 0, 8}, makeSliceFromNode(addTwoNumbers(l1, l2)))
	})
	t.Run("example2", func(t *testing.T) {
		l1 := makeNodeFromSlice([]int{0})
		l2 := makeNodeFromSlice([]int{0})
		assert.EqualValues(t, []int{0}, makeSliceFromNode(addTwoNumbers(l1, l2)))
	})
	t.Run("example3", func(t *testing.T) {
		l1 := makeNodeFromSlice([]int{9, 9, 9, 9, 9, 9, 9})
		l2 := makeNodeFromSlice([]int{9, 9, 9, 9})
		assert.EqualValues(t, []int{8, 9, 9, 9, 0, 0, 0, 1}, makeSliceFromNode(addTwoNumbers(l1, l2)))
	})
}

func makeNodeFromSlice(s []int) *ListNode {
	head := &ListNode{}
	n := head
	for i, e := range s {
		n.Val = e
		if i < len(s)-1 {
			n.Next = &ListNode{}
			n = n.Next
		}
	}
	return head
}

func makeSliceFromNode(node *ListNode) []int {
	cur := node
	var ret []int
	for cur != nil {
		ret = append(ret, cur.Val)
		cur = cur.Next
	}
	return ret
}
