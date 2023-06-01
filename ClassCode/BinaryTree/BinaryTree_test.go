package BinaryTree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	n5 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 4}
	n2 := &TreeNode{
		Val:   2,
		Left:  n4,
		Right: n5,
	}
	n3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root := &TreeNode{
		Val:   1,
		Left:  n2,
		Right: n3,
	}
	ans := Inorder(root)
	fmt.Println(ans)
}
