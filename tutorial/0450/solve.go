package _450

import (
	"fmt"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		ans := inorder(root, make([]int, 0))
		ans = delElem(ans, key)
		fmt.Println(ans)
		root = buildBST(ans)
	}
	if root == nil {
		return root
	}
	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}

func inorder(root *TreeNode, ans []int) []int {
	var fn func(*TreeNode)
	fn = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			fn(node.Left)
		}
		ans = append(ans, node.Val)
		if node.Right != nil {
			fn(node.Right)
		}
	}
	fn(root)
	return ans
}

func delElem(arr []int, tar int) []int {
	for i, n := range arr {
		if n == tar {
			if i == len(arr)-1 {
				return arr[:i]
			}
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func buildBST(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	sort.Ints(arr)
	root := &TreeNode{
		Val: arr[0],
	}
	cur := root
	for i := 1; i < len(arr); i++ {
		cur.Right = &TreeNode{
			Val: arr[i],
		}
		cur = cur.Right
	}
	return root
}

func deleteNode02(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}

	cur := root.Right
	for cur.Left != nil {
		cur = cur.Left
	}
	cur.Left = root.Left

	root.Val = root.Right.Val
	root.Left = root.Right.Left
	root.Right = root.Right.Right

	return root
}
