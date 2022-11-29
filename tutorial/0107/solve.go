package _107

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		l := q.Len()
		s := make([]int, 0)
		for i := 0; i < l; i++ {
			e := q.Remove(q.Front()).(*TreeNode)
			if e.Left != nil {
				q.PushBack(e.Left)
			}
			if e.Right != nil {
				q.PushBack(e.Right)
			}
			s = append(s, e.Val)
		}
		ret = append(ret, s)
	}

	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
	}

	return ret
}
