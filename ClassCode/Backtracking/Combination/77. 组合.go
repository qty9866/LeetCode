package Combination

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
		// 确定还要选多少个数
		d := k - len(path)
		if d == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		// 不选i
		if i > d {
			dfs(i - 1)
		}
		// 选i
		path = append(path, i)
		dfs(i - 1)
		path = path[:len(path)-1]
	}
	dfs(n)
	return
}
