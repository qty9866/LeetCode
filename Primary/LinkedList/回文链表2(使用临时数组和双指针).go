package main

import "fmt"

func isPalindrome2(head *ListNode) bool {
	var vals []int
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	n4 := &ListNode{Val: 1}
	n3 := &ListNode{Val: 2, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	fmt.Println(isPalindrome2(n1))
}
