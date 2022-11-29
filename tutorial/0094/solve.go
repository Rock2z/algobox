package _094

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		e := st.Remove(st.Back()).(*TreeNode)
		if e.Left == nil && e.Right == nil {
			ret = append(ret, e.Val)
			continue
		}
		st.PushBack(e.Right)
		st.PushBack(&TreeNode{Val: e.Val})
		st.PushBack(e.Left)
	}
	return ret
}
