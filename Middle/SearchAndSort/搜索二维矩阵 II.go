package main

import "sort"

/*
编写一个高效的算法来搜索mxn矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

示例 1:
输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true

示例 2：
输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
输出：false
*/

// 方法一: 直接遍历矩阵
func searchMatrix1(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, v := range row {
			if v == target {
				return true
			}
		}
	}
	return false
}

// 方法二：二分查找
func searchMatrix2(matrix [][]int, target int) bool {
	for _, row := range matrix {
		i := sort.SearchInts(row, target)
		if i < len(row) && row[i] == target {
			return true
		}
	}
	return false
}

// 方法三：Z字型查找
func searchMatrix3(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if nums := BinarySearchFirstBiggerInNums(matrix, target); nums == -1 {
		return false
	} else {
		find := BinarySearchingN(matrix[nums], target)
		if !find {
			if nums == 0 {
				return false
			} else {
				return BinarySearchingN(matrix[nums-1], target)
			}
		}
	}
	return false
}

func BinarySearchFirstBiggerInNums(arr [][]int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid][0] < target {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1][0] < target {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func BinarySearchingN(nums []int, target int) bool {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
