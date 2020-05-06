package linkedlist

import (
	"golang_leetcode/datastruct"
)

//https://leetcode-cn.com/problems/linked-list-cycle-ii/
func detectCycle(head *datastruct.ListNode) *datastruct.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			cur := head
			for cur != fast {
				fast = fast.Next
				cur = cur.Next
			}
			return cur
		}
	}
	return nil
}
