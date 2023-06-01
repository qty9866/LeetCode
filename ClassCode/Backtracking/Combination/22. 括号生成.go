package Combination

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
输入：n = 1
输出：["()"]
*/
func generateParenthesis(n int) (ans []string) {
	if n == 0 {
		return nil
	}
	m := 2 * n
	path := make([]byte, m)
	var dfs func(i, open int)
	dfs = func(i, open int) {
		// 结束条件
		if i == m {
			ans = append(ans, string(path))
			return
		}
		// '('的个数小于n
		if i < m {
			path = append(path, '(')
			dfs(i+1, open+1)
		}
		// ')'的个数小于'('
		if i-open < open {
			path = append(path, ')')
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return
}
