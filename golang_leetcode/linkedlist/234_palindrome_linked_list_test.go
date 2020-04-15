package linkedlist

import (
	"fmt"
	"golang_leetcode/datastruct"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	array := [...]int{1}

	var head *datastruct.ListNode
	var current *datastruct.ListNode

	for _, v := range array {
		if head == nil {
			head = &datastruct.ListNode{v, nil}
			current = head
		} else {
			temp := &datastruct.ListNode{v, nil}
			current.Next = temp
			current = temp
		}
	}
	fmt.Print(isPalindrome(head))
}
