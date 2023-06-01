package main

// BinarySearch 二分查找的循环实现
func BinarySearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// BinaryDiGui 二分查找的递归实现
func BinaryDiGui(arr []int, target int) int {
	n := len(arr)
	return bs(arr, target, 0, n-1)
}

func bs(arr []int, target int, low int, high int) int {
	mid := low + (high-low)>>1
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return bs(arr, target, mid+1, high)
	} else {
		return bs(arr, target, low, mid-1)
	}
}

// BinarySearchFirstEqual 二分查找：查找第一个值等于给定值的元素
func BinarySearchFirstEqual(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			if mid == 0 || arr[mid-1] != target {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// BinarySearchLastEqual 查找最后一个值等于给定值的元素
func BinarySearchLastEqual(arr []int, x int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) >> 1
		if arr[mid] < x {
			low = mid + 1
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			if mid == n-1 || arr[mid+1] != x {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

// BinarySearchFirstBigger 二分查找：查找第一个值大于等于给定值的元素
func BinarySearchFirstBigger(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
