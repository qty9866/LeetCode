package DynamicProgramming

/*
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

示例 1：
输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。

示例 2：
输入：strs = ["10", "0", "1"], m = 1, n = 1
输出：2
解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
*/
func findMaxForm(strs []string, m int, n int) int {
	length := len(strs)
	zero, one := make([]int, length), make([]int, length)
	for i := range strs {
		for _, j := range strs[i] {
			if j == '0' {
				zero[i]++
			} else {
				one[i]++
			}
		}
	}
	// 初始化dp数组
	dp := make([][][]int, length+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i := range strs {
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i+1][j][k] = dp[i][j][k]
				if j >= zero[i] && k >= one[i] {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-zero[i]][k-one[i]]+1)
				}
			}
		}
	}
	return dp[length][m][n]
}
