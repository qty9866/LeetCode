package BackTracking

import "unicode"

/*
给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。
返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。

示例 1：
输入：s = "a1b2"
输出：["a1b2", "a1B2", "A1b2", "A1B2"]

示例 2:
输入: s = "3z4"
输出: ["3z4","3Z4"]
*/

// 深度优先遍历
/*
采用回溯的思想，从左往右依次遍历字符，当在进行搜索时，搜索到字符串 sss 的第 iii 个字符 ccc 时:

如果 c 为一个数字，则我们继续检测下一个字符；
如果 c 为一个字母，我们将字符中的第 i 个字符 c 改变大小写形式后，往后继续搜索，完成改写形式的子状态搜索后，
我们将 c 进行恢复，继续往后搜索；
如果当前完成字符串搜索后，则表示当前的子状态已经搜索完成，该序列为全排列中的一个；


*/
func letterCasePermutation(s string) (ans []string) {
	var dfs func([]byte, int)
	dfs = func(s []byte, pos int) {
		// 当前枚举位置小于字符串长度且当前枚举的字符是一个数字 那么向前
		for pos < len(s) && unicode.IsDigit(rune(s[pos])) {
			pos++
		}
		if pos == len(s) {
			// 枚举到头了
			ans = append(ans, string(s))
			return
		}
		// 如果当前枚举的字符是一个字母
		// 不对字符做更改
		dfs(s, pos+1)
		// 更改
		s[pos] ^= 32
		dfs(s, pos+1)
		s[pos] ^= 32 // 恢复现场
	}
	dfs([]byte(s), 0)
	return
}
