package ArrayAndString

import "sort"

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

示例 1:
输入: str = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

示例 2:
输入: str = [""]
输出: [[""]]

示例 3:
输入: str = ["a"]
输出: [["a"]]
*/
// 方法1： 使用排序
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	res := make([][]string, 0, len(mp))
	for _, strings := range mp {
		res = append(res, strings)
	}
	return res
}

// 方法二： 哈希表计数
func groupAnagrams1(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, v := range str {
			cnt[v-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, strings := range mp {
		ans = append(ans, strings)
	}
	return ans
}
