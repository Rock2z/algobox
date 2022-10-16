package _203RemoveLinkedListElements

import (
	"algobox/top_100_linked_questions/common"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *common.ListNode, val int) *common.ListNode {
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
