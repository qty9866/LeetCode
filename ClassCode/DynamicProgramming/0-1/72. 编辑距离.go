package DynamicProgramming

/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

示例 1：
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
*/

// 方法1 ：使用回溯思想
// dfs(i,j) = dfs(i-1,j-1) , word1[i] == word2[j]
// dfs(i,j) = min(dfs(i,j-1),dfs(i-1,j),dfs(i-1,j-1))+1 ,word1[i] != word2[j]
func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		if word1[i] == word2[j] {
			return dfs(i-1, j-1)
		}
		return Min(dfs(i, j-1), dfs(i-1, j), dfs(i-1, j-1)) + 1
	}
	return dfs(n-1, m-1)
}

func Min(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		} else {
			return b
		}
	} else {
		if a > c {
			return c
		} else {
			return a
		}
	}
}

// 回溯+记忆化搜索
func minDistance1(s, t string) int {
	n, m := len(s), len(t)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示还没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		p := memo[i][j]
		if p != -1 {
			return p
		}
		defer func() { p = res }()
		if s[i] == t[j] {
			return dfs(i-1, j-1)
		}
		return min(min(dfs(i-1, j), dfs(i, j-1)), dfs(i-1, j-1)) + 1
	}
	return dfs(n-1, m-1)
}

// 翻译成递归
func minDistance2(s, t string) int {
	n, m := len(s), len(t)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := 1; j <= m; j++ {
		f[0][j] = j
	}
	for i, x := range s {
		f[i+1][0] = i + 1
		for j, y := range t {
			if x == y {
				f[i+1][j+1] = f[i][j]
			} else {
				f[i+1][j+1] = min(min(f[i][j+1], f[i+1][j]), f[i][j]) + 1
			}
		}
	}
	return f[n][m]
}
