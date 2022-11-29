package _145

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		e := st.Remove(st.Back()).(*TreeNode)
		if e == nil {
			continue
		}
		if e.Left == nil && e.Right == nil {
			ret = append(ret, e.Val)
			continue
		}
		st.PushBack(&TreeNode{Val: e.Val})
		st.PushBack(e.Right)
		st.PushBack(e.Left)
	}
	return ret
}
