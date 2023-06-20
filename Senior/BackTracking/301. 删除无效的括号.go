package BackTracking

/*
给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
返回所有可能的结果。答案可以按 任意顺序 返回。

示例 1：
输入：s = "()())()"
输出：["(())()","()()()"]

示例 2：
输入：s = "(a)())()"
输出：["(a())()","(a)()()"]

示例 3：
输入：s = ")("
输出：[""]
*/

// IsValid 定义函数 判断其中左右括号的个数是否有效
func IsValid(str string) bool {
	cnt := 0
	for _, ch := range str {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func dfs(ans *[]string, str string, start, lr, rr int) {
	if lr == 0 && rr == 0 {
		if IsValid(str) {
			*ans = append(*ans, str)
		}
		return
	}
	for i := start; i < len(str); i++ {
		// 如果是形似"((())"这样重复的
		if i != start && str[i] == str[i-1] {
			continue
		}
		// 如果剩余的字符无法满足去掉的数量要求，直接返回
		if lr+rr > len(str)-1 {
			return
		}
		// 尝试去掉一个左括号
		if lr > 0 && str[i] == '(' {
			dfs(ans, str[:i]+str[i+1:], i, lr-1, rr)
		}
		// 尝试去掉一个有括号
		if rr > 0 && str[i] == ')' {
			dfs(ans, str[:i]+str[i+1:], i, lr, rr-1)
		}
	}
}

func removeInvalidParentheses(s string) (ans []string) {
	lr, rr := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lr++
		} else if ch == ')' {
			if lr == 0 {
				rr++
			} else {
				lr--
			}
		}
	}
	dfs(&ans, s, 0, lr, rr)
	return
}
