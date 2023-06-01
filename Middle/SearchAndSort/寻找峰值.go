package main

/*
峰值元素是指其值严格大于左右相邻值的元素。
给你一个整数数组nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

你可以假设nums[-1] = nums[n] = -∞ 。
你必须实现时间复杂度为 O(log n) 的算法来解决此问题。

思路：使用红蓝染色法
红色：false 表示峰值的左侧
蓝色：true 表示峰值或者峰值的右侧
M为当前元素 比较M和M+1
  - 如果 M < M+1 表示 M 在峰值的左侧
  - 如果 M > M+1 表示 M 要么是峰值 要么在峰值的
*/
func findPeakElement(nums []int) int {
	left, right := -1, len(nums)-1 // (-1,n-1) 开区间
	for left+1 < right {           // 开区间不为空
		mid := left + (right-left)>>1
		if nums[mid] > nums[mid+1] { //是峰值或者在峰值的右边
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
