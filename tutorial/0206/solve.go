package _206

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	cur := head
	var before, after *ListNode
	for cur != nil {
		after = cur.Next
		cur.Next = before
		before = cur
		cur = after
	}
	return before
}
