package BinaryTree

/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。
示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true

示例 2：
输入：root = [1,2,2,null,3,null,3]
输出：false
*/
// 方法1:使用递归
func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	return Symmetric(root.Left, root.Right)
}

func Symmetric(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	return left.Val == right.Val && Symmetric(left.Left, right.Right) && Symmetric(left.Right, right.Left)
}

// 方法2: 使用迭代
func isSymmetric2(root *TreeNode) bool {
	u, v := root, root
	var res []*TreeNode
	res = append(res, u)
	res = append(res, v)
	for len(res) > 0 {
		u, v = res[0], res[1]
		res = res[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		res = append(res, u.Left)
		res = append(res, v.Right)
		res = append(res, u.Right)
		res = append(res, v.Left)
	}
	return true
}
