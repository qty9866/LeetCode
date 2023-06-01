package main

import "math"

/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/

// 递归方法
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(root *TreeNode, low, up int) bool {
	if root == nil {
		return true
	}
	if root.Val <= low || root.Val >= up {
		return false
	}
	return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, up)
}

// 中序遍历之后检测
func isValidBST2(root *TreeNode) bool {
	res := InOrder(root)
	for i := 1; i < len(res); i++ {
		if res[i] < res[i-1] {
			return false
		}
	}
	return true
}

func InOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := InOrder(root.Left)
	res = append(res, root.Val)
	res = append(res, InOrder(root.Right)...)
	return res
}
