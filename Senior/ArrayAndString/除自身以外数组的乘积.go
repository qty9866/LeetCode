package ArrayAndString

/*
给你一个整数数组nums，返回 数组answer，其中answer[i]等于nums中除nums[i]之外其余各元素的乘积。
题目数据 保证 数组nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。
请不要使用除法，且在O(n) 时间复杂度内完成此题。

示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]
*/
func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	// 初始化L数组,L[0] = 1 ,L[i] = L[i-1] * nums[i-1]
	L := make([]int, n)
	L[0] = 1
	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	// 初始化R数组,R[n-1] = 1 ,R[i] = R[i+1] * nums[i+1]
	R := make([]int, n)
	R[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = L[i] * R[i]
	}
	return ans
}

// 思路2: 将L和R用输出数组来计算
func productExceptSelf2(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)

	// ans[i] 表示索引 i 左侧所有元素的乘积
	// 因为索引为 '0' 的元素左侧没有元素， 所以 ans[0] = 1
	ans[0] = 1
	for i := 1; i < length; i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}

	// R 为右侧所有元素的乘积
	// 刚开始右边没有元素，所以 R = 1
	R := 1
	for i := length - 1; i >= 0; i-- {
		// 对于索引 i，左边的乘积为 ans[i]，右边的乘积为 R
		ans[i] = ans[i] * R
		// R 需要包含右边所有的乘积，所以计算下一个结果时需要将当前值乘到 R 上
		R *= nums[i]
	}
	return ans
}
