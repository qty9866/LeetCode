package main

import "sort"

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。

示例1:
输入: s = "anagram", t = "nagaram"
输出: true

示例 2:
输入: s = "rat", t = "car"
输出: false
*/
func isAnagram(s string, t string) bool {
	hm1 := make(map[rune]int)
	hm2 := make(map[rune]int)
	for _, v := range s {
		hm1[v]++
	}
	for _, v := range t {
		hm2[v]++
	}
	for r := range hm1 {
		if hm1[r] != hm2[r] {
			return false
		}
	}
	return true
}

func isAnagram2(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}
