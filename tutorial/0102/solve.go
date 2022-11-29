package _102

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
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
	return ret
}

type MyNode struct {
	t *TreeNode
	x int // level of tree, start 0
}

func levelOrder2(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	q := list.New()
	q.PushBack(&MyNode{t: root})

	for q.Len() > 0 {
		e := q.Remove(q.Front()).(*MyNode)
		for len(ret)-1 < e.x {
			ret = append(ret, make([]int, 0))
		}
		ret[e.x] = append(ret[e.x], e.t.Val)
		if e.t.Left != nil {
			q.PushBack(&MyNode{t: e.t.Left, x: e.x + 1})
		}
		if e.t.Right != nil {
			q.PushBack(&MyNode{t: e.t.Right, x: e.x + 1})
		}
	}
	return ret
}
