package main

/*
给定单链表的头节点head，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
第一个节点的索引被认为是 奇数 ， 第二个节点的索引为偶数 ，以此类推。
请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
你必须在O(1)的额外空间复杂度和O(n)的时间复杂度下解决这个问题。

示例 1:
输入: head = [1,2,3,4,5]
输出:[1,3,5,2,4]

示例 2:
输入: head = [2,1,3,5,6,4,7]
输出: [2,3,6,7,1,5,4]
*/

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	singleTail := head
	evenHead := head.Next
	evenTail := evenHead
	for singleTail.Next != nil || evenTail.Next != nil {
		singleTail.Next = evenTail.Next
		if singleTail.Next != nil {
			singleTail = singleTail.Next
		}
		evenTail.Next = singleTail.Next
		if evenTail.Next != nil {
			evenTail = evenTail.Next
		}
	}
	singleTail.Next = evenHead
	return head
}
