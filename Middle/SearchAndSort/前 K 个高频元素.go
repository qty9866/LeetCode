package main

import "sort"

/*
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]

示例 2:
输入: nums = [1], k = 1
输出: [1]
*/
func topKFrequent(nums []int, k int) (ans []int) {
	HashTable := map[int]int{}
	var tmp []int
	for _, num := range nums {
		HashTable[num]++
	}
	for _, frequency := range HashTable {
		tmp = append(tmp, frequency)
	}
	sort.Ints(tmp)
	frequency := tmp[len(tmp)-k]
	for i, fre := range HashTable {
		if fre >= frequency {
			ans = append(ans, i)
		}
	}
	return
}
