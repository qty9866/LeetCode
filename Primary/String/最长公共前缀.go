package main

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串""。

示例 1：
输入：str = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：str = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[1]
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = findPrefix(prefix, strs[i])
	}
	return prefix
}

func findPrefix(a, b string) string {
	var prefix string
	m := min(len(a), len(b))
	for i := 0; i < m; i++ {
		if a[i] == b[i] {
			prefix += string(a[i])
		} else {
			break
		}
	}
	return prefix
}

func main() {
	strs := []string{"dog", "dogcecar", "dogcar", "dog1", "sd"}
	print(longestCommonPrefix(strs))
}

// 力扣题解 横向扫描
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

type TrieNode struct {
	data     string
	children []TrieNode
}
