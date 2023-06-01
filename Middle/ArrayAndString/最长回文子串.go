package ArrayAndString

/*
给你一个字符串 s，找到 s 中最长的回文子串。
如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"
*/

// 方法一：动态规划
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	ans := ""
	for l := 0; l < n; l++ { // l代表了当前枚举的子串长度
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
			if dp[i][j] && l+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

func longestPalindrome1(s string) string {
	n := len(s)
	ret := string(s[0])
	for i := 0; i < n; i++ {
		// "*****XX*****"形式
		if i < n-1 && s[i] == s[i+1] {
			j := 1
			for i-j >= 0 && i+j+1 < n && s[i-j] == s[i+j+1] {
				j++
			}
			if len(ret) < j*2 {
				ret = s[i-j+1 : i+j+1]
			}
		}

		// "*****XYX*****"形式
		if i > 0 && i < n-1 && s[i-1] == s[i+1] {
			j := 1
			for i-j-1 >= 0 && i+j+1 < n && s[i-j-1] == s[i+j+1] {
				j++
			}
			if len(ret) < j*2+1 {
				ret = s[i-j : i+j+1]
			}
		}
	}
	return ret
}
