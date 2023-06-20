package DynamicProgramming

import (
	"math"
)

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

示例 1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

示例 2：
输入：coins = [2], amount = 3
输出：-1

示例 3：
输入：coins = [1], amount = 0
输出：0
*/
// 方法1 使用回溯的思想（注意这里的i是可以重复选的）
// 状态转移方程为 dfs(i,c) = min(dfs(i-1,c),dfs(i,c-w[i])+v[i])

func CoinChange(coins []int, amount int) int {
	n := len(coins)
	var dfs func(i, c int) int
	dfs = func(i, c int) int {
		if i < 0 {
			if c == 0 {
				return 0
			} else {
				return math.MaxInt64 / 2
			}
		}
		if c < coins[i] {
			return dfs(i-1, c)
		}
		return min(dfs(i-1, c), dfs(i, c-coins[i])+1)
	}
	ans := dfs(n-1, amount)
	if ans < math.MaxInt64/2 {
		return ans
	} else {
		return -1
	}
}

// 方法2
// 递归搜索 + 保存计算结果 = 记忆化搜索
func coinChange1(coins []int, amount int) int {
	n := len(coins)
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, amount+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(i, c int) int
	dfs = func(i, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 0
			}
			return math.MaxInt64 >> 1 // 除以2防止下面的+1操作发生溢出
		}
		C := &cache[i][c]
		if *C != -1 {
			return *C
		}
		defer func() { *C = res }()
		if c < coins[i] {
			return dfs(i-1, c)
		}
		return min(dfs(i-1, c), dfs(i, c-coins[i]+1))
	}
	ans := dfs(n-1, amount)
	if ans < math.MaxInt64>>1 {
		return ans
	} else {
		return -1
	}
}

//1:1 翻译成递推
func coinChange2(coins []int, amount int) int {
	n := len(coins)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	// f[0][j] 代表最后一次操作
	for j := range f[0] {
		f[0][j] = math.MaxInt / 2
	}
	f[0][0] = 0
	for i, coin := range coins {
		for c := 0; c <= amount; c++ {
			if c < coin {
				f[i+1][c] = f[i][c]
			} else {
				f[i][c] = min(f[i][c], f[i][c-coin]+1)
			}
		}
	}
	ans := f[n][amount]
	if ans < math.MaxInt/2 {
		return ans
	} else {
		return -1
	}
}

// 空间优化：两个数组（滚动数组）
func coinChange3(coins []int, amount int) int {
	n := len(coins)
	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	for j := range f[0] {
		f[0][j] = math.MaxInt / 2
	}
	f[0][0] = 0
	for i, coin := range coins {
		for c := 0; c <= amount; c++ {
			if c < coin {
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = min(f[i%2][c], f[(i + 1)][c-coin]+1)
			}
		}
	}
	ans := f[n%2][amount]
	if ans < math.MaxInt/2 {
		return ans
	} else {
		return -1
	}
}

// 空间优化：一个数组
func coinChange4(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := range f {
		f[i] = math.MaxInt / 2 // 除 2 是防止下面 + 1 溢出
	}
	f[0] = 0
	for _, coin := range coins {
		for c := coin; c <= amount; c++ {
			f[c] = min(f[c], f[c-coin]+1)
		}
	}
	ans := f[amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
