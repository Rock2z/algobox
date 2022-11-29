package _144

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// stack
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		pop := st.Remove(st.Back()).(*TreeNode)
		if pop != nil {
			ret = append(ret, pop.Val)
			st.PushBack(pop.Right)
			st.PushBack(pop.Left)
		}
	}
	return ret
}

// recursive
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{root.Val}
	if root.Left != nil {
		ret = append(ret, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		ret = append(ret, preorderTraversal(root.Right)...)
	}
	return ret
}
