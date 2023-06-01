package main

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{
		Val: -1,
	}
	point := head
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			point.Next = list2
			list2 = list2.Next
		} else {
			point.Next = list1
			list1 = list1.Next
		}
		point = point.Next
	}
	if list1 != nil {
		point.Next = list1
		list1 = list1.Next
	}
	if list2 != nil {
		point.Next = list2
		list2 = list2.Next
	}
	return head.Next
}

func main() {
	n5 := &ListNode{Val: 9}
	n4 := &ListNode{Val: 7, Next: n5}
	n3 := &ListNode{Val: 5, Next: n4}
	n2 := &ListNode{Val: 3, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}

	l4 := &ListNode{Val: 6}
	l3 := &ListNode{Val: 4, Next: l4}
	l2 := &ListNode{Val: 2, Next: l3}
	l1 := &ListNode{Val: 1, Next: l2}
	mergeTwoLists(l1, n1)
}
