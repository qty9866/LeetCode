package BackTracking

/*
给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例 1：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

示例 2：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true

示例 3：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false
*/
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	var dfs func(row, column, i int) bool
	dfs = func(row, column, i int) bool {
		if i == len(word) {
			return true
		}
		if row < 0 || row >= m || column < 0 || column >= n {
			return false
		}
		if used[row][column] || board[row][column] != word[i] {
			return false
		}
		used[row][column] = true
		CanFindRest := dfs(row+1, column, i+1) || dfs(row-1, column, i+1) || dfs(row, column-1, i+1) || dfs(row, column+1, i+1)
		if CanFindRest {
			return true
		} else {
			used[row][column] = false
			return false
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
