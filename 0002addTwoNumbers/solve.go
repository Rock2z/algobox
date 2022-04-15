package _002addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return solve(l1, l2)
}

func solve(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	c := 0
	for {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		cur.Val = (a + b + c) % 10
		c = (a + b + c) / 10
		if l1 == nil && l2 == nil {
			break
		}
		cur.Next = &ListNode{}
		cur = cur.Next
	}
	if c != 0 {
		cur.Next = &ListNode{Val: c}
	}
	return head
}
