package BinaryTree

/*
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
示例：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

// 方法1: 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left, right int
	// 递归出口
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil {
		left = maxDepth(root.Left)
	}
	if root.Right != nil {
		right = maxDepth(root.Right)
	}
	return max(left+1, right+1)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 利用队列 进行迭代
func maxDepth2(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			node := queue[0]
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
