package DoublePointer

/*
困難題

给定n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9
*/

// 方法1 使用两个数组记录左右的最大值
func trap(height []int) (ans int) {
	n := len(height)
	if n <= 2 {
		return
	}
	LeftM := make([]int, n)
	RightM := make([]int, n)
	LeftM[0], RightM[n-1] = 0, 0
	preMax := 0
	for i := 1; i < n; i++ {
		preMax = max(preMax, height[i-1])
		LeftM[i] = preMax
	}
	behindMax := 0
	for i := n - 2; i > 0; i-- {
		behindMax = max(behindMax, height[i+1])
		RightM[i] = behindMax
	}
	for i := 1; i < n-1; i++ {
		if height[i] < min(LeftM[i], RightM[i]) {
			ans += min(LeftM[i], RightM[i]) - height[i]
		}
	}
	return
}
