package BinaryTree

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：true

示例 2：
输入：root = [1,2,2,3,3,null,null,4,4]
输出：false

示例 3：
输入：root = []
输出：true
*/
func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	LeftH := getHeight(node.Left)
	if LeftH == -1 {
		return -1 // 提前退出 不再递归
	}
	rightH := getHeight(node.Right)
	if rightH == -1 || abs(LeftH-rightH) > 1 {
		return -1
	}
	return max(LeftH, rightH) + 1
}

func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}
