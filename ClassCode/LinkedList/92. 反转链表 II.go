package LinkedList

/*
给你单链表的头指针 head 和两个整数left 和 right ，其中left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

示例 1：
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

示例 2：
输入：head = [5], left = 1, right = 1
输出：[5]
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	p0 := dummy
	for i := 0; i < left-1; i++ {
		// 移动到left的前一个节点
		p0 = p0.Next
	}
	var pre, cur *ListNode = nil, p0.Next
	for i := 0; i < right-left+1; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	p0.Next.Next = cur
	p0.Next = pre
	return dummy.Next
}
