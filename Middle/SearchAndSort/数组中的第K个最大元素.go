package main

import "sort"

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5

示例2:
输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4
*/
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	index := 1
	if k == 1 {
		return nums[len(nums)-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] != nums[i+1] {
			index++
			if index == k {
				return nums[i]
			}
		}
	}
	return nums[len(nums)-1]
}
