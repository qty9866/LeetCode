package main

import (
	"sort"
)

/*
两个数组的交集 II
给你两个整数数组
nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致
（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]

示例 2:
输入：nums1 = v, nums2 = [9,4,9,8,4]
输出：[4,9]
*/

type IntSlice1 []int

func (p IntSlice1) Len() int {
	return len(p)
}
func (p IntSlice1) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p IntSlice1) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func intersect(nums1 []int, nums2 []int) []int {
	var result []int
	sort.Sort(IntSlice1(nums1))
	sort.Sort(IntSlice1(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return result
}

//func main() {
//	n1 := []int{1, 2, 2, 1}
//	n2 := []int{2}
//	re := intersect(n1, n2)
//	fmt.Println(re)
//}
