package BackTracking

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
*/
func generateParenthesis(n int) (ans []string) {
	if n == 0 {
		return
	}
	m := 2 * n
	path := make([]byte, m)
	var dfs func(i, open int)
	dfs = func(i, open int) {
		if i == m {
			ans = append(ans, string(path))
			return
		}
		// '('个数小于n
		if open < n {
			path[i] = '('
			dfs(i+1, open+1)
		}
		// ')'小于'('个数
		if i-open < open {
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return
}
