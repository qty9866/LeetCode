package BackTracking

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。

示例 1：
输入：n = 4, k = 2
输出：
[

	[2,4],
	[3,4],
	[2,3],
	[1,2],
	[1,3],
	[1,4],

]

示例 2：
输入：n = 1, k = 1
输出：[[1]]
*/
func combine(n, k int) (ans [][]int) {
	var path []int
	var dfs func(int)
	dfs = func(i int) {
		// d表示还需要选多少个数字
		d := k - len(path)
		if d == 0 {
			ans = append(ans, path)
			return
		}
		for j := i; j >= d; j-- {
			path = append(path, j)
			dfs(j - 1)
			path = path[:len(path)-1]
		}
	}
	dfs(n)
	return
}
