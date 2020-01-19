package datastruct

import "fmt"

type ListNode struct {
	//节点值
	Val int
	//下一个节点
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) Print() {
	node := list.Head
	for {
		if node != nil {
			fmt.Println(node.Val)
		}
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}

//添加元素到链表尾部
func (list *LinkedList) AddLast(node *ListNode) {
	last := list.Head
	if last == nil {
		list.Head = node
		return
	}
	for {
		if last.Next == nil {
			break
		}
		last = last.Next
	}
	last.Next = node
}
