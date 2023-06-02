package BackTracking

/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。

示例 1：
输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]

示例 2：
输入：s = "a"
输出：[["a"]]
*/

func partition(s string) (ans [][]string) {
	n := len(s)
	var path []string
	var dfs func(i, start int)
	dfs = func(i, start int) {
		if i == n {
			ans = append(ans, append([]string(nil), path...))
			return
		}
		// 如果不选当前字符
		if i+1 < n {
			dfs(i+1, start)
		}

		// 如果选择当前字符的话
		if isPalindrome(s, start, i) {
			path = append(path, s[start:i+1])
			dfs(i+1, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
