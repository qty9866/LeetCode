package DynamicProgramming

/*
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
假设每一种面额的硬币有无限个。
题目数据保证结果符合 32 位带符号整数。

示例 1：
输入：amount = 5, coins = [1, 2, 5]
输出：4
解释：有四种方式可以凑成总金额：
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

示例 2：
输入：amount = 3, coins = [2]
输出：0
解释：只用面额 2 的硬币不能凑成总金额 3 。

示例 3：
输入：amount = 10, coins = [10]
输出：1
*/

// 方法1 回溯 但是时间复杂度是指数级别 应该会超时
// 这里用的是倒序 因为coins长度为1的时候不好处理
func change(amount int, coins []int) (ans int) {
	n := len(coins)
	var dfs func(i, c int)
	dfs = func(i, c int) {
		if i < 0 {
			if c == 0 {
				ans++
			}
			return
		}
		// 不选当前的零钱，直接跳过
		dfs(i-1, c)
		// 剩下的钱比当前的这个币值大 选当前币值
		if c-coins[i] >= 0 {
			dfs(i, c-coins[i])
		}
	}
	dfs(n-1, amount)
	return
}

// 回溯+记忆化搜索
func change1(amount int, coins []int) int {
	n := len(coins)
	cache := make([][]int, len(coins)+1)
	// 初始化cache数组
	for i := range cache {
		cache[i] = make([]int, amount+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(i, c int) int
	dfs = func(i, c int) int {
		if c == 0 {
			return 1
		}
		if i < 0 || c < 0 {
			return 0
		}
		if cache[i][c] != -1 {
			return cache[i][c]
		}
		var res int
		// 不选当前的零钱
		res += dfs(i-1, c)
		// 选当前的零钱
		res += dfs(i, c-coins[i])
		cache[i][c] = res
		return res
	}
	return dfs(n-1, amount)
}

// 动态规划 想状态转移方程 dp[i] = dp[]
func change2(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
