package LinkedList

/*
给定一个已排序的链表的头 head ，删除所有重复的元素，使每个元素只出现一次。返回 已排序的链表。

示例 1：
输入：head = [1,1,2]
输出：[1,2]

示例 2：
输入：head = [1,1,2,3,3]
输出：[1,2,3]
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur != nil {
		if cur.Val == cur.Next.Val {
			cur = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
