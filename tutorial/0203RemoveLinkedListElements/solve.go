package _203RemoveLinkedListElements

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements02(head *ListNode, val int) *ListNode {
	ret := head
	for {
		if ret == nil {
			return nil
		}
		if ret.Val == val {
			ret = ret.Next
		} else {
			break
		}
	}
	cur := ret.Next
	last := ret
	for {
		if cur == nil {
			break
		}
		if cur.Val == val {
			last.Next = cur.Next
			cur = cur.Next
		} else {
			last = cur
			cur = cur.Next
		}
	}
	return ret
}

func removeElements(head *ListNode, val int) *ListNode {
	cur := head
	for cur != nil {
		if cur.Val == val {
			cur = cur.Next
		} else {
			break
		}
	}
	ret := cur
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return ret
}

func removeElementsUsingDummyHead(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
