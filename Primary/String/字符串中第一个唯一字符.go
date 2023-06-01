package main

import (
	"fmt"
)

/*
字符串中的第一个唯一字符
给定一个字符串s，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1。

示例 1：
输入: s = "leetcode"
输出: 0

示例 2:
输入: s = "loveleetcode"
输出: 2

*/

func firstUniqChar(s string) int {
	//str := strings.Split(s, "")
	ct := [26]int{}
	for _, v := range s {
		ct[v-'a']++
	}
	for i, v := range s {
		if ct[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	char := firstUniqChar("leetcode")
	fmt.Println(char)
}
