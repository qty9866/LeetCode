package main

/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。

示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：


输入：root = [1,2,2,null,3,null,3]
输出：false
*/

// 方法一： 递归
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

// 方法二：使用队列迭代

func isSymmetric1(root *TreeNode) bool {
	l, r := root, root
	var res []*TreeNode
	res = append(res, l)
	res = append(res, r)
	for len(res) > 0 {
		u, v := res[0], res[1]
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
		res = append(res, u.Right)
		res = append(res, v.Right)
		res = append(res, v.Left)
	}
	return true
}
