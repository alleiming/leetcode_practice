package linkedlist

import (
	"fmt"
	"golang_leetcode/datastruct"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	linkedList := datastruct.NewLinkedList()
	for i := 1; i < 5; i++ {
		linkedList.AddLast(&datastruct.ListNode{i, nil})
	}

	linkedList.Print()

	firstNode := detectCycle(linkedList.Head)
	fmt.Print(firstNode)
}
