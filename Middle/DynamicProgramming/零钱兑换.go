package DynamicProgramming

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。
你可以认为每种硬币的数量是无限的。

示例1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

示例 2：
输入：coins = [2], amount = 3
输出：-1

示例 3：
输入：coins = [1], amount = 0
输出：0
*/
func coinChange(coins []int, amount int) int {
	n := len(coins)
	// 找到第i个问题与前面的子问题的关系
	// dp[i]定义为组成金额为i的最少硬币数量
	// dp[i] = dp[i-Cj]+1 其中Cj为数组coins中的一个硬币面额
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0 // 金額等于0显然不需要硬币
	for i := 1; i <= amount; i++ {
		for j := 0; j < n; j++ {
			// 如果当前硬币面额coins[j]小于当前面额i
			if coins[j] <= i {
				if dp[i-coins[j]]+1 < dp[i] {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
