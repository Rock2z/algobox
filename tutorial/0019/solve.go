package _019

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	cur := head
	for ; n > 0; n-- {
		cur = cur.Next
	}
	prev := dummy
	for cur != nil {
		cur = cur.Next
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	return dummy.Next
}
