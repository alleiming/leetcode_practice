package datastruct

import "testing"

func TestAddLast(t *testing.T) {
	linkedList := NewLinkedList()
	for i := 1; i < 10; i++ {
		linkedList.AddLast(&ListNode{i, nil})
	}
}
