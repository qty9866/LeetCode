package BackTracking

import "strings"

/*
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。


示例 1：
输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。

示例 2：
输入：n = 1
输出：[["Q"]]
*/
// 基于集合的回溯
// 创建三个数组 cols、diagonals1、diagonals2 分别来记录这一列、方向1(左上-->右下)、方向2(右上-->左下)上有没有别的皇后
// i代表在第i行进行排列

var ans [][]string

func solveNQueens(n int) [][]string {
	ans = [][]string{}
	// 初始化Queens数组用于记录皇后的位置
	Queens := make([]int, n)
	for i := range Queens {
		Queens[i] = -1
	}
	// cols 数组用于记录皇后摆放在了哪一列
	cols := make([]bool, n)
	// diagonals1,diagonals2用于记录两个方向的斜线上是否有皇后
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backTracking(0, n, Queens, cols, diagonals1, diagonals2)
	return ans
}

func backTracking(row, n int, Queues []int, cols []bool, diagonals1, diagonals2 map[int]bool) {
	// 递归的结束条件
	if row == n {
		board := generate(Queues, n)
		ans = append(ans, board)
		return
	}
	for i := 0; i < n; i++ {
		if cols[i] {
			continue
		}
		if diagonals1[row-i] {
			continue
		}
		if diagonals2[row+i] {
			continue
		}
		Queues[row] = i
		cols[i], diagonals1[row-i], diagonals2[row+i] = true, true, true
		backTracking(row+1, n, Queues, cols, diagonals1, diagonals2)
		// 恢复现场
		Queues[row] = -1
		cols[i], diagonals1[row-i], diagonals2[row+i] = false, false, false
	}
}

func generate(Queens []int, n int) (board []string) {
	for i := 0; i < n; i++ {
		path := make([]byte, n)
		for j := 0; j < n; j++ {
			path[j] = '.'
		}
		path[Queens[i]] = 'Q'
		board = append(board, string(path))
	}
	return
}

// 灵神写法
func solveNQueens1(n int) (ans [][]string) {
	col := make([]int, n)
	onPath := make([]bool, n)
	diag1 := make([]bool, n*2-1)
	diag2 := make([]bool, n*2-1)
	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			board := make([]string, n)
			for i, c := range col {
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
			}
			ans = append(ans, board)
			return
		}
		for c, on := range onPath {
			rc := r - c + n - 1
			if !on && !diag1[r+c] && !diag2[rc] {
				col[r] = c
				onPath[c], diag1[r+c], diag2[rc] = true, true, true
				dfs(r + 1)
				onPath[c], diag1[r+c], diag2[rc] = false, false, false // 恢复现场
			}
		}
	}
	dfs(0)
	return
}
