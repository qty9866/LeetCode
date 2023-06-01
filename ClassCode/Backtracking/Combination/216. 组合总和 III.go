package Combination

/*
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。

示例 1:
输入: k = 3, n = 7
输出: [[1,2,4]]
解释:
1 + 2 + 4 = 7
没有其他符合的组合了。

示例 2:
输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
解释:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
没有其他符合的组合了。

示例 3:
输入: k = 4, n = 1
输出: []
解释: 不存在有效的组合。
在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。
*/
func combinationSum3(k int, n int) (ans [][]int) {
	var path []int
	var dfs func(i, sum int)
	dfs = func(i, sum int) {
		d := k - len(path)
		if d == 0 && sum == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		// 不选i
		if sum+i < n {
			dfs(i+1, sum)
		}
		// 选i
		path = append(path, i)
		dfs(i+1, sum+i)
		path = path[:len(path)-1]
		sum -= i
	}
	dfs(1, 0)
	return
}

func combinationSum2(k int, n int) (ans [][]int) {
	var path []int
	var dfs func(i, sum int)
	dfs = func(i, sum int) {
		// 递归结束条件
		if len(path) == k && sum == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		// 剪枝：1、省略了过多数字后存在数字不够的情况 2、sum>目标值
		if len(path)+10-i < k || sum > n {
			return
		}
		// 如果不选当前这个数字
		dfs(i+1, sum)
		// 如果选当前数字
		path = append(path, i)
		dfs(i+1, sum+i)
		path = path[:len(path)-1]
	}
	dfs(1, 0)
	return
}
