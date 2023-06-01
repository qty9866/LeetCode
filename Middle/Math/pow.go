package main

import "math"

/*
实现pow(x, n)，即计算 x 的整数n 次幂函数（即，xn ）。
示例 1：
输入：x = 2.00000, n = 10
输出：1024.00000

示例 2：
输入：x = 2.10000, n = 3
输出：9.26100

示例 3：
输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25
*/
func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
