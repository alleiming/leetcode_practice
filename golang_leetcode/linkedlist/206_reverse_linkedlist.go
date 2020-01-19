package main

import (
	"golang_leetcode/datastruct"
)

//递归实现
func reverseList(head *datastruct.ListNode) *datastruct.ListNode {
	if head.Next == nil {
		return head
	}
	// 步骤 1: 先翻转 node 之后的链表
	newHead := reverseList(head.Next)

	// 步骤 2: 再把原 node 节点后继结点的后继结点指向 node，node 的后继节点设置为空(防止形成环)
	head.Next.Next = head
	head.Next = nil
	// 步骤 3: 返回翻转后的头结点
	return newHead
}

//迭代法实现
func reverseListByIteration(head *datastruct.ListNode) *datastruct.ListNode {
	var pre *datastruct.ListNode
	current := head
	for {
		if current == nil {
			break
		}
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}

	return pre
}

func main() {
	linkedList := datastruct.NewLinkedList()
	for i := 1; i < 100; i++ {
		linkedList.AddLast(&datastruct.ListNode{i, nil})
	}
	head := reverseListByIteration(linkedList.Head)
	//head := reverseList(linkedList.Head)
	linkedList.Head = head
	linkedList.Print()
}
