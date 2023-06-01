package main

/*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
*/
func rotate(nums []int, k int) {
	length := len(nums)
	if k >= length {
		k = k % length
	}
	if k == 0 {
		return
	} else {
		Inverse(nums, 0, length-1)
		Inverse(nums, k, length-1)
		Inverse(nums, 0, k-1)
	}
}

func Inverse(array []int, start, end int) {
	mid := (start + end) / 2
	for i := start; i <= mid; i++ {
		j := start + end - i
		array[i], array[j] = array[j], array[i]
	}
}
