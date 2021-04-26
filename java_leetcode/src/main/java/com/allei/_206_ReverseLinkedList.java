package com.allei;

/**
 * 链表翻转
 * https://leetcode-cn.com/problems/reverse-linked-list/
 */
public class _206_ReverseLinkedList {

    public static class ListNode {
        int val;
        ListNode next;

        ListNode() {
        }

        ListNode(int val) {
            this.val = val;
        }

        ListNode(int val, ListNode next) {
            this.val = val;
            this.next = next;
        }

        @Override
        public String toString() {
            return "ListNode{" +
                    "val=" + val +
                    ", next=" + next +
                    '}';
        }

        public void println() {
            ListNode listNode = this;
            while (listNode != null) {
                System.out.print(listNode.val);
                listNode = listNode.next;
            }
            System.out.println();
        }
    }

    /**
     * 迭代法
     *
     * @param head
     * @return
     */
    public static ListNode iterateReverseList(ListNode head) {
        ListNode current = head, next, pre = null;
        while (current != null) {
            next = current.next;
            current.next = pre;
            pre = current;
            current = next;
        }
        return pre;
    }

    /**
     * 递归法
     *
     * @param head
     * @return
     */
    public static ListNode recursionReverseList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode newHead = recursionReverseList(head.next);
        head.next.next = head;
        head.next = null;
        return newHead;
    }


    public static void main(String[] args) {
        ListNode e = new ListNode(5, null);
        ListNode d = new ListNode(4, e);
        ListNode c = new ListNode(3, d);
        ListNode b = new ListNode(2, c);
        ListNode a = new ListNode(1, b);
        ListNode result = iterateReverseList(a);
        result.println();
        ListNode newResult = recursionReverseList(result);
        newResult.println();
    }
}
