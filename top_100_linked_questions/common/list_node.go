package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeNodeFromSlice(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
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

func MakeSliceFromNode(node *ListNode) []int {
	cur := node
	var ret []int
	for cur != nil {
		ret = append(ret, cur.Val)
		cur = cur.Next
	}
	return ret
}
