package linkedlist

import (
	"golang_leetcode/datastruct"
)

func middleNode(head *datastruct.ListNode) *datastruct.ListNode {
	if head == nil {
		return nil
	}
	//map存储节点位置与节点的对应关系,也可以用数组实现
	nodeMap := make(map[int]*datastruct.ListNode)
	current := head
	listLen := 0
	for {
		if current == nil {
			break
		}
		listLen++
		nodeMap[listLen] = current
		current = current.Next
	}
	return nodeMap[listLen/2+1]
}

func middleNodeByFastSlowPointer(head *datastruct.ListNode) *datastruct.ListNode {
	fast := head
	slow := head
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
