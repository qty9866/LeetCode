package DynamicProgramming

/*
给你一个整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

示例 1：
输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3

示例 2：
输入：nums = [1], target = 1
输出：1
*/
func findTargetSumWays(nums []int, target int) (count int) {
	var dfs func(i, sum int)
	dfs = func(i, sum int) {
		if i == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		dfs(i+1, sum+nums[i])
		dfs(i+1, sum-nums[i])
	}
	dfs(0, 0)
	return
}

// 递归搜索 + 保存计算结果 = 记忆化搜索
func findTargetSumWays1(nums []int, target int) int {
	// 添加"+"的数字和为p，所有元素和为s
	// 添加"-"的数字和为(s-p),那么target = p-(s-p) ==> p = (s+target)/2
	// 所以(s+p) > 0 && (s+p)%2 == 0
	for _, x := range nums {
		target += x
	}
	if target < 0 || target%2 != 0 {
		return 0
	}
	target /= 2 // p
	n := len(nums)
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, target+1)
		for j := range cache[i] {
			cache[i][j] = -1 // -1 表示没用访问过
		}
	}
	var dfs func(i, c int) int
	dfs = func(i, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		C := &cache[i][c]
		if *C != -1 {
			return *C
		}
		defer func() { *C = res }()
		if c < nums[i] {
			return dfs(i-1, c)
		}
		return dfs(i-1, c) + dfs(i-1, c-nums[i])
	}
	return dfs(n-1, target)
}

//1:1 翻译成递推
func findTargetSumWays2(nums []int, target int) int {
	for _, x := range nums {
		target += x
	}
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2

	n := len(nums)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, target+1)
	}
	f[0][0] = 1
	for i, x := range nums {
		for c := 0; c <= target; c++ {
			if c < x {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = f[i][c] + f[i][c-x]
			}
		}
	}
	return f[n][target]
}

// 空间优化：两个数组（滚动数组）
func findTargetSumWays3(nums []int, target int) int {
	for _, x := range nums {
		target += x
	}
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2

	n := len(nums)
	f := [2][]int{make([]int, target+1), make([]int, target+1)}
	f[0][0] = 1
	for i, x := range nums {
		for c := 0; c <= target; c++ {
			if c < x {
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = f[i%2][c] + f[i%2][c-x]
			}
		}
	}
	return f[n%2][target]
}

// 空间优化：一个数组
func findTargetSumWays4(nums []int, target int) int {
	for _, x := range nums {
		target += x
	}
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2

	f := make([]int, target+1)
	f[0] = 1
	for _, x := range nums {
		for c := target; c >= x; c-- {
			f[c] += f[c-x]
		}
	}
	return f[target]
}
