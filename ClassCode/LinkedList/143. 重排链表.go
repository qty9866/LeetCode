package LinkedList

/*
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1：
输入：head = [1,2,3,4]
输出：[1,4,2,3]

示例 2：
输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]
*/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	// Find Middle Node
	middle := FindMiddle(head)
	// Reverse the List
	reverse := Reverse(middle)

	for reverse.Next != nil {
		ntx := head.Next
		ntx2 := reverse.Next
		head.Next = reverse
		reverse.Next = ntx
		head = ntx
		reverse = ntx2
	}
}

func FindMiddle(head *ListNode) *ListNode {
	slow, quick := head, head
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	return slow
}

func Reverse(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		ntx := cur.Next
		cur.Next = pre
		pre = cur
		cur = ntx
	}
	return pre
}
