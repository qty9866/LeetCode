package DynamicProgramming

/*
状态转移方程
① dp[i][j] = dp[i−1][j−1]+1,text1[i−1] = text2[j−1]
② dp[i][j] = max(dp[i-1][j],dp[i][j-1]) ,text1[i−1] != text2[j−1]
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
