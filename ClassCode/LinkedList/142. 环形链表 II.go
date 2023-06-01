package LinkedList

/*
给定一个链表的头节点 head，返回链表开始入环的第一个节点。如果链表无环，则返回null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内
部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：
pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。

示例2：
输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。
*/
// 思路1:利用一个Map记录是否遍历过这个节点 如果有遍历过了就返回这个节点
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	visited := map[*ListNode]bool{}
	for head != nil {
		if visited[head] {
			return head
		} else {
			visited[head] = true
		}
		head = head.Next
	}
	return nil
}

// 思路2: 根据快慢指针进行分析
// head到入口的距离为 a , 入口到相遇点的距离为 b , 相遇点再到入口的距离为c
// 可以推出 a-c = k(b+c) 也就是说slow从相遇点出发，同时head从头结点出发 slow到达入口时
// 此时head到达入口处的距离正好是环长的k倍，两个指针都继续走，相遇的时候便是入口指针
// 当快慢指针在相遇点相遇后，

func detectCycle1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}
