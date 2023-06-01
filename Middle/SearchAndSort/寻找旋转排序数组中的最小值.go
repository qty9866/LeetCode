package main

/*
已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

示例 1：
输入：nums = [3,4,5,1,2]
输出：1
解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。

示例 2：
输入：nums = [4,5,6,7,0,1,2]
输出：0
解释：原数组为 [0,1,2,4,5,6,7] ，旋转 4 次得到输入数组。

示例 3：
输入：nums = [11,13,15,17]
输出：11
解释：原数组为 [11,13,15,17] ，旋转 4 次得到输入数组。
*/
func findMin1(nums []int) int {
	left, right := -1, len(nums)-1 // (-1,n-1)开区间
	for left+1 < right {           // 开区间不为空
		mid := left + (right-left)>>1
		if nums[mid] < nums[len(nums)-1] {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

func lowerBound(nums []int, left, right, target int) int {
	r0 := right
	for left+1 < right { // 开区间不为空
		/*
			循环不变量：
				nums[left] < target
				nums[right] >= left
		*/
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid // 缩小范围到(mid,right)
		} else {
			right = mid
		}
	}
	if right == r0 || nums[right] != target {
		return -1
	}
	return right
}

func search(nums []int, target int) int {
	i := findMin1(nums)
	if target > nums[len(nums)-1] {
		return lowerBound(nums, -1, i, target)
	}
	return lowerBound(nums, i-1, len(nums), target) // 右段
}
