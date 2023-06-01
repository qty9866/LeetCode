package main

/*
给定一个包含红色、白色和蓝色、共n 个元素的数组nums，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数0、1和2分别表示红色、白色和蓝色。
必须在不使用库内置的 sort 函数的情况下解决这个问题。

示例 1：
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]

示例 2：
输入：nums = [2,0,1]
输出：[0,1,2]
*/

// 方法一 统计0、1、2的个数 然后返回一个新的数组
func sortColors1(nums []int) {
	zero, one, two := 0, 0, 0
	var ans []int
	for _, value := range nums {
		switch value {
		case 0:
			zero++
		case 1:
			one++
		case 2:
			two++
		}
	}
	for i := 0; i < zero; i++ {
		ans = append(ans, 0)
	}
	for i := 0; i < one; i++ {
		ans = append(ans, 1)
	}
	for i := 0; i < two; i++ {
		ans = append(ans, 2)
	}
	copy(nums, ans)
}

// 方法二：单指针
func sortColors2(nums []int) {
	target := SwapColors(nums, 0)
	SwapColors(nums[target:], 1)
}

func SwapColors(colors []int, target int) (countTarget int) {
	for index, value := range colors {
		if value == target {
			colors[index], colors[countTarget] = colors[countTarget], colors[index]
			countTarget++
		}
	}
	return
}
