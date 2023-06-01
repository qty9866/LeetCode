package SubSet

import "math/bits"

/*
给你一个下标从 0 开始的 m x n 二进制矩阵 mat 和一个整数 cols ，表示你需要选出的列数。
如果一行中，所有的 1 都被你选中的列所覆盖，那么我们称这一行 被覆盖 了。
请你返回在选择 cols 列的情况下，被覆盖 的行数 最大 为多少。

示例 1：
输入：mat = [[0,0,0],[1,0,1],[0,1,1],[0,0,1]], cols = 2
输出：3
解释：
如上图所示，覆盖 3 行的一种可行办法是选择第 0 和第 2 列。
可以看出，不存在大于 3 行被覆盖的方案，所以我们返回 3 。

示例 2：
输入：mat = [[1],[0]], cols = 1
输出：2
解释：
选择唯一的一列，两行都被覆盖了，原因是整个矩阵都被覆盖了。
所以我们返回 2 。
*/
func maximumRows(mat [][]int, cols int) (ans int) {
	m, n := len(mat), len(mat[0])
	mask := make([]int, m)
	for i, row := range mat {
		for j, v := range row {
			mask[i] |= v << j
		}
	}
	for set := 1<<cols - 1; set < 1<<n; {
		cnt := 0
		for _, row := range mask {
			if row&set == row { // row 是 set 的子集，所有 1 都被覆盖
				cnt++
			}
		}
		ans = max(ans, cnt)
		lb := set & -set
		x := set + lb
		// 下式等价于 set = (set^x)/lb>>2 | x
		set = (set^x)>>bits.TrailingZeros(uint(lb))>>2 | x
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
