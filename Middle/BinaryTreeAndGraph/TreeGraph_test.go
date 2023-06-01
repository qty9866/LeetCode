package BinaryTreeAndGraph

import (
	"fmt"
	"testing"
)

func TestInorder1(t *testing.T) {
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}
	n7 := &TreeNode{Val: 7}
	n2 := &TreeNode{Val: 2, Left: n4, Right: n5}
	n3 := &TreeNode{Val: 3, Left: n6, Right: n7}
	n1 := &TreeNode{Val: 1, Left: n2, Right: n3}
	res := inorderTraversal1(n1)
	fmt.Println(res)
}

func TestInorder2(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 2, Left: n3}
	n1 := &TreeNode{Val: 1, Right: n2}
	res := inorderTraversal1(n1)
	fmt.Println(res)
}

func TestBuildTree(t *testing.T) {
	preOrder := []int{3, 9, 20, 15, 7}
	InOrder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preOrder, InOrder)
	fmt.Println(tree)
}

func TestKthSmallest(t *testing.T) {
	n2 := &TreeNode{Val: 2}
	n1 := &TreeNode{Val: 1, Right: n2}
	n4 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 2, Left: n1, Right: n4}
	res := inorderTraversal(n3)
	fmt.Println(res)
}
