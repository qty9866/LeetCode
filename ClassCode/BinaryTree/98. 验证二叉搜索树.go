package BinaryTree

import "math"

/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

示例 1：
输入：root = [2,1,3]
输出：true

示例 2：
输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。
*/

// 前序遍历思路
func dfs(node *TreeNode, left, right int) bool {
	if node == nil {
		return true
	}
	x := node.Val
	return left < x && x < right && dfs(node.Left, left, x) && dfs(node.Right, x, right)
}

func isValidBST1(root *TreeNode) bool {
	return dfs(root, math.MinInt64, math.MaxInt64)
}

// 中序遍历思路: 中序遍历二叉树 如果有不是按照递增顺序的 那么就返回false
func isValidBST2(root *TreeNode) bool {
	res := InOrder(root)
	for i := 0; i < len(res)-1; i++ {
		if res[i+1] <= res[i] {
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

// 使用后续遍历

func postOrder(node *TreeNode) (int, int) {
	if node == nil {
		return math.MaxInt, math.MinInt
	}
	lMin, lMax := postOrder(node.Left)
	rMin, rMax := postOrder(node.Right)
	x := node.Val
	if x <= lMax || x >= rMin {
		return math.MinInt, math.MaxInt
	}
	return min(lMin, x), max(rMax, x)
}

func isValidBST3(root *TreeNode) bool {
	_, mx := postOrder(root)
	return mx != math.MaxInt
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
