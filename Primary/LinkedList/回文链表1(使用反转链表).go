package main

import "fmt"

/*
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

示例 1：
输入：head = [1,2,2,1]
输出：true

示例 2：
输入：head = [1,2]
输出：false

思路：
找到这个链表的中间节点 将链表的前半部分进行反转 然后双指针比较
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	point := findMiddleNode(head)
	point = point.Next
	start := Reverse(point)
	for start != nil {
		if start.Val != head.Val {
			return false
		}
		start = start.Next
		head = head.Next
	}
	return true
}

func findMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func Reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {
	n2 := &ListNode{Val: 1}
	n1 := &ListNode{Val: 1, Next: n2}
	fmt.Println(isPalindrome(n1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}
