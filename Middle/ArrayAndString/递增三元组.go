package ArrayAndString

/*
给你一个整数数组nums ，判断这个数组中是否存在长度为 3 的递增子序列。

如果存在这样的三元组下标 (i, j, k)且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。

示例 1：
输入：nums = [1,2,3,4,5]
输出：true
解释：任何 i < j < k 的三元组都满足题意

示例 2：
输入：nums = [5,4,3,2,1]
输出：false
解释：不存在满足题意的三元组

示例 3：
输入：nums = [2,1,5,0,4,6]
输出：true
解释：三元组 (3, 4, 5) 满足题意，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6
*/

// 暴力求解 时间复杂度O(n^3) 无法通过
func increasingTriplet(nums []int) bool {
	flag := false
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			if nums[j] <= nums[i] {
				continue
			} else {
				for k := j + 1; k < n; k++ {
					if nums[k] > nums[j] {
						flag = true
					}
				}
			}
		}
	}
	return flag
}

// 方法二：双向遍历
func increasingTriplet1(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	leftMin := make([]int, n)
	rightMax := make([]int, n)
	leftMin[0] = nums[0]
	rightMax[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		leftMin[i] = Min(leftMin[i-1], nums[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = Max(rightMax[i+1], nums[i])
	}
	for i := 0; i < n-1; i++ {
		if nums[i] > leftMin[i] && nums[i] < rightMax[i] {
			return true
		}
	}
	return false
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
