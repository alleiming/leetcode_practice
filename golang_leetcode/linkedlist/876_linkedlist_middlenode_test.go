package linkedlist

import (
	"fmt"
	"golang_leetcode/datastruct"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	linkedList := datastruct.NewLinkedList()
	for i := 1; i < 6; i++ {
		linkedList.AddLast(&datastruct.ListNode{i, nil})
	}
	linkedList.Print()
	middleNode := middleNode(linkedList.Head)
	fmt.Println(middleNode.Val)
}

func TestMiddleNodeByFastSlowPointer(t *testing.T) {
	linkedList := datastruct.NewLinkedList()
	for i := 1; i < 6; i++ {
		linkedList.AddLast(&datastruct.ListNode{i, nil})
	}
	linkedList.Print()
	middleNode := middleNodeByFastSlowPointer(linkedList.Head)
	fmt.Println(middleNode.Val)
}
