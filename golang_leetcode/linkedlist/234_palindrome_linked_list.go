package linkedlist

import "golang_leetcode/datastruct"

//https://leetcode-cn.com/problems/palindrome-linked-list/
//判断一个链表是否为回文链表
//快慢指针解答
func isPalindrome(head *datastruct.ListNode) bool {
	fast := head
	slow := head
	var pre *datastruct.ListNode

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}

	//链表长度为奇数情况
	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		if slow.Val != pre.Val {
			return false
		}

		slow = slow.Next
		pre = pre.Next
	}

	return true
}
