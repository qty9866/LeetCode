package BinaryTree

/*
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
假设二叉树中至少有一个节点。

示例 1:
输入: root = [2,1,3]
输出: 1

示例 2:
输入: [1,2,3,4,null,5,6,null,null,7]
输出: 7
*/
// 思路: 还是层序遍历 到最后一层 返回第一个数值
// 自己写的
func findBottomLeftValue(root *TreeNode) int {
	tmp := []*TreeNode{root}
	for len(tmp) > 0 {
		n := len(tmp)
		values := make([]int, n)
		for i := range values {
			node := tmp[0]
			tmp = tmp[1:]
			values[i] = node.Val
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		if len(tmp) == 0 {
			return values[0]
		}
	}
	return 0
}

// 层序遍历 从右边向左边遍历 返回最后一个节点的值
func findBottomLeftValue2(root *TreeNode) int {
	node := root
	q := []*TreeNode{root}
	for len(q) > 0 {
		node, q = q[0], q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	return node.Val
}

// 深度优先搜索
func findBottomLeftValue3(root *TreeNode) (curVal int) {
	curHeight := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		height++
		dfs(node.Left, height)
		dfs(node.Right, height)
		if height > curHeight {
			curHeight = height
			curVal = node.Val
		}
	}
	dfs(root, 0)
	return
}
