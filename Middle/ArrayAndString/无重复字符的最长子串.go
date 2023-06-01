package ArrayAndString

/*
给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。
示例1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
请注意，你的答案必须是子串长度，"pwke"是一个子序列，不是子串。
*/
func lengthOfLongestSubstring(s string) (ans int) {
	left := 0
	cnt := [128]int{}
	for i, c := range s {
		cnt[c]++
		for cnt[c] > 1 { // 不满足要求
			cnt[s[left]]--
			left++
		}
		ans = Max(ans, i-left+1)
	}
	return ans
}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
