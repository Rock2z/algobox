package _199

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {

	ret := make([][]int, 0)
	if root == nil {
		return []int{}
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

	res := make([]int, 0, len(ret))
	for _, x := range ret {
		res = append(res, x[len(x)-1])
	}
	return res
}
