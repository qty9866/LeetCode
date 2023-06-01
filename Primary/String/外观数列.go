package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给定一个正整数 n ，输出外观数列的第 n 项。

「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。

你可以将其视作是由递归公式定义的数字字符串序列：

countAndSay(1) = "1"
countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。
前五项如下：
1.     1
2.     11
3.     21
4.     1211
5.     111221
第一项是数字 1
描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"

示例 1：

输入：n = 1
输出："1"
解释：这是一个基本样例。

示例 2：
输入：n = 4
输出："1211"
解释：
countAndSay(1) = "1"
countAndSay(2) = 读 "1" = 一 个 1 = "11"
countAndSay(3) = 读 "11" = 二 个 1 = "21"
countAndSay(4) = 读 "21" = 一 个 2 + 一 个 1 = "12" + "11" = "1211"
*/

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	count := 0
	str := countAndSay(n - 1)
	var x string
	for i := 0; i < len(str)-1; i++ {
		var j int
		j = i + 1
		if i == len(str)-2 {
			if str[i] == str[j] {
				count += 2
				x += strconv.Itoa(count)
				x += strconv.Itoa(int(str[i] - '0'))
				break
			} else {
				count++
				x += strconv.Itoa(count)
				x += strconv.Itoa(int(str[i] - '0'))
				x += strconv.Itoa(1)
				x += strconv.Itoa(int(str[j] - '0'))
				break
			}
		}
		if str[i] == str[j] {
			count++

		}
		if str[i] != str[j] {
			count++
			x += strconv.Itoa(count)
			x += strconv.Itoa(int(str[i] - '0'))
			count = 0
		}
	}
	return x
}

// LeetCode题解
func countAndSay1(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			cur.WriteString(strconv.Itoa(j - start))
			cur.WriteByte(prev[start])
		}
		prev = cur.String()
	}
	return prev
}

func main() {
	fmt.Println(countAndSay1(2))
}
