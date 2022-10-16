package _024

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := head.Next

	cur := head
	after := cur.Next
	cur.Next = swapPairs(after.Next)
	after.Next = cur

	return ret
}
