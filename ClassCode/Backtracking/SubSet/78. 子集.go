package SubSet

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

示例 2：
输入：nums = [0]
输出：[[],[0]]
*/

// 输入的视角: 选还是不选
func subsets(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	path := make([]int, 0, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...)) // 固定答案
			return
		}
		// 不选 nums[i]
		dfs(i + 1)
		// 选 nums[i]
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1] // 恢复现场
	}
	dfs(0)
	return ans
}

// 答案的视角（选哪个数）
func subsets1(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	path := make([]int, 0, n)
	var dfs func(int)
	dfs = func(i int) {
		ans = append(ans, append([]int(nil), path...)) // 固定答案
		if i == n {
			return
		}
		for j := i; j < n; j++ { // 枚举选择的数字
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1] // 恢复现场
		}
	}
	dfs(0)
	return ans
}
