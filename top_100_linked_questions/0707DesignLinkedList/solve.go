package _707DesignLinkedList

import (
	"algobox/top_100_linked_questions/common"
)

type MyLinkedList struct {
	head *common.ListNode
	tail *common.ListNode
	l    int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.l {
		return -1
	}
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	node := &common.ListNode{
		Val:  val,
		Next: l.head,
	}
	l.head = node
	l.l++

	if l.tail == nil {
		l.tail = node
	}
}

func (l *MyLinkedList) AddAtTail(val int) {
	node := &common.ListNode{
		Val: val,
	}
	if l.head == nil {
		l.head = node
	}
	if l.tail != nil {
		l.tail.Next = node
	}
	l.tail = node
	l.l++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		l.AddAtHead(val)
		return
	}
	if index == l.l {
		l.AddAtTail(val)
		return
	}
	if index > l.l {
		return
	}

	l.l++
	cur := l.head
	last := l.head
	for i := 0; i < index; i++ {
		last = cur
		cur = cur.Next
	}
	node := &common.ListNode{
		Val:  val,
		Next: l.head,
	}
	last.Next = node
	node.Next = cur
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.l {
		return
	}
	l.l--
	if l.l == 0 {
		l.head = nil
		l.tail = nil
		return
	}
	if index == 0 {
		l.head = l.head.Next
		return
	}

	cur := l.head
	last := l.head
	for i := 0; i < index; i++ {
		last = cur
		cur = cur.Next
	}
	last.Next = cur.Next
	if cur.Next == nil {
		l.tail = last
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
