package LinkedList

/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

示例 1:
输入：head = [1,2,3,4]
输出：[2,1,4,3]

示例 2：
输入：head = []
输出：[]

示例 3：
输入：head = [1]
输出：[1]
*/
func swapPairs(head *ListNode) *ListNode {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	// 设立哨兵节点dummy
	dummy := &ListNode{Next: head}
	p0 := dummy
	var cur, pre *ListNode = p0.Next, nil
	for ; n >= 2; n -= 2 {

		for i := 0; i < 2; i++ {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		ntx := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = ntx
	}
	return dummy.Next
}
