package DoublePointer

/*
给定一个含有n个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[nums l, nums l+1, ..., nums r-1, nums r] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组[4,3]是该条件下的长度最小的子数组。

示例 2：
输入：target = 4, nums = [1,4,4]
输出：1

示例 3：
输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0
*/
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans, s, left := n+1, 0, 0
	for right, x := range nums {
		s += x
		for s >= target {
			ans = min(ans, right-left+1)
			s -= nums[left]
			left++
		}
	}
	if ans <= n {
		return ans
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
