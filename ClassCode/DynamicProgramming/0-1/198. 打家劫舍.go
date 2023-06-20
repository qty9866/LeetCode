package DynamicProgramming

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装
有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

示例 1：
输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。

	偷窃到的最高金额 = 1 + 3 = 4 。

示例 2：
输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。

	偷窃到的最高金额 = 2 + 9 + 1 = 12 。
*/

// 动态规划
// 状态转移方程： dp[i] = max(dp[i-1],dp[i-2]+nums[i])
func rob(nums []int) (ans int) {
	n := len(nums)
	// dp数组代表在当前节点能够获取到的最高金额
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 递归搜索 + 保存计算结果 = 记忆化搜索
func rob1(nums []int) int {
	n := len(nums)
	cache := make([]int, n)
	for i := range cache {
		cache[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if cache[i] != -1 {
			return cache[i]
		}
		res := max(dfs(i-1), dfs(i-2)+nums[i])
		cache[i] = res
		return res
	}
	return dfs(n - 1)
}

// 优化空间复杂度到O(1)
// 时间 0 ms 击败 100%
// 内存 1.9 MB 击败 99.86%
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	first, second := 0, 0
	for i := 0; i < n; i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}
