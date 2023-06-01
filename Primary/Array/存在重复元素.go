/*
存在重复元素
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。

示例 1：

输入：nums = [1,2,3,1]
输出：true
*/
package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}
func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	nums := []int{2, 31, 5, 6, 3}
	sort.Sort(IntSlice(nums))
	fmt.Println("intSort:", nums) // [2,3,5,6,31]
}
func containsDuplicate(nums []int) bool {
	sort.Sort(IntSlice(nums))
	length := len(nums) - 1
	for i := 0; i < length; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
