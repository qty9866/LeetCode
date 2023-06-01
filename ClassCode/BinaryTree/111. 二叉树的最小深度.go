package BinaryTree

import "math"

/*
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：2

示例 2：
输入：root = [2,null,3,null,4,null,5,null,6]
输出：5
*/

// 思路1: 递归
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//var left, right int
	minD := math.MaxInt64
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD
}

// 思路2: 使用队列进行迭代
func minDepth1(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			node := queue[0]
			if node.Left == nil && node.Right == nil {
				return ans + 1
			}
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			length--
		}
		ans++
	}
	return
}
