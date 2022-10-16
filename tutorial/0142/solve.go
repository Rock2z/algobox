package _142

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

//func detectCycleXJB(head *ListNode) *ListNode {
//	for head != nil && head.Next != nil {
//		if uintptr(unsafe.Pointer(head.Next)) <= uintptr(unsafe.Pointer(head)) {
//			return head.Next
//		} else {
//			head = head.Next
//		}
//	}
//	return nil
//}
