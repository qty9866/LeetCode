package SubSet

import "strconv"

/*
累加数 是一个字符串，组成它的数字可以形成累加序列。
一个有效的 累加序列 必须 至少 包含 3 个数。除了最开始的两个数以外，序列中的每个后续数字必须是它之前两个数字之和。
给你一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是 累加数 。如果是，返回 true ；否则，返回 false 。
说明：累加序列里的数，除数字 0 之外，不会 以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。

示例 1：
输入："112358"
输出：true
解释：累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8

示例 2：
输入："199100199"
输出：true
解释：累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199
*/
func isAdditiveNumber(num string) bool {
	var isCanAdded func(num string, first, second, sumIDx int) bool
	isCanAdded = func(num string, first, second, sumIdx int) bool {
		if sumIdx == len(num) {
			return true
		}

		sumStr := strconv.Itoa(first + second)
		if sumIdx+len(sumStr) > len(num) {
			return false
		}
		actualSum := num[sumIdx : sumIdx+len(sumStr)]
		return sumStr == actualSum && isCanAdded(num, second, first+second, sumIdx+len(sumStr))
	}

	first := 0
	for i := 0; i < len(num); i++ {
		if i > 0 && num[0] == '0' {
			return false
		}
		first = first*10 + int(num[i]-'0')
		second := 0
		for j := i + 1; j < len(num); j++ {
			second = second*10 + int(num[j]-'0')
			if j > i+1 && num[i+1] == '0' {
				break
			}
			if j+1 < len(num) && isCanAdded(num, first, second, j+1) {
				return true
			}
		}
	}
	return false
}

func isAdditiveNumber1(arr string) bool {
	n := len(arr)
	var dfs func(idx int, numCnt int, num1 int, num2 int) bool
	dfs = func(idx, numCnt, num1, num2 int) bool {
		if idx == n {
			return numCnt > 2
		}
		for i, cur := idx, 0; i < n && !(arr[idx] == '0' && i != idx); i++ {
			cur = cur*10 + int(arr[i]-'0')
			if numCnt < 2 || numCnt >= 2 && num1+num2 == cur {
				if dfs(i+1, numCnt+1, num2, cur) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, 0, 0, 0)
}
