package DynamicProgramming

/*
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
示例 1：
输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。

示例 2：
输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。
*/
// 回溯的思想
func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum >> 1

	var dfs func(i, c int) bool

	dfs = func(i, c int) bool {
		if i == len(nums) || c > sum {
			return false
		}
		if c == sum {
			return true
		}
		return dfs(i+1, c+nums[i]) || dfs(i+1, c)
	}
	return dfs(0, 0)
}

// 思路：典型的0-1背包问题,从len(nums)个数中选取数字，每个数字只能选一次，使得和为sum/2
func canPartition1(nums []int) bool {
	n := len(nums)
	var sum int
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
		sum += num
	}
	if sum%2 != 0 || max > (sum>>1) {
		return false
	}
	sum = sum >> 1
	// 初始化dp数组
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	// 遍历dp数组
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j <= sum; j++ {
			if j < v {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			}
		}
	}
	return dp[n-1][sum]
}

// 空间优化：一个数组
func canPartition2(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum >> 1
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := sum; j > 0; j-- {
			if j-nums[i] >= 0 && dp[j-nums[i]] {
				dp[j] = true
				if j == sum {
					return true
				}
			}
		}
	}
	return false
}
