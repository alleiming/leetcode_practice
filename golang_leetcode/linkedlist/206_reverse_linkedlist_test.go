package linkedlist

import (
	"golang_leetcode/datastruct"
	"testing"
)

//迭代实现
func TestReverseListByIteration(t *testing.T) {
	linkedList := datastruct.NewLinkedList()
	for i := 1; i < 100; i++ {
		linkedList.AddLast(&datastruct.ListNode{i, nil})
	}
	head := reverseListByIteration(linkedList.Head)
	linkedList.Head = head
	linkedList.Print()
}

//递归
func TestReverseList(t *testing.T) {
	linkedList := datastruct.NewLinkedList()
	for i := 1; i < 100; i++ {
		linkedList.AddLast(&datastruct.ListNode{i, nil})
	}
	head := reverseList(linkedList.Head)
	linkedList.Head = head
	linkedList.Print()
}
