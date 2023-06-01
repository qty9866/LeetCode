package main

import (
	"fmt"
	"math"
)

/*
整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。

示例 1：
输入：x = 123
输出：321

示例 2：
输入：x = -123
输出：-321
*/
func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		dig := x % 10
		x = x / 10
		rev = rev*10 + dig
	}
	return rev
}

func main() {
	rev1 := reverse(480)
	fmt.Println(rev1)
	rev2 := reverse(122)
	fmt.Println(rev2)
	rev3 := reverse(3)
	fmt.Println(rev3)
}
