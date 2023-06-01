package ArrayAndString

/*
矩阵置零
给定一个m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

示例 1：
输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]

示例 2：
输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
*/
func setZeroes(matrix [][]int) {
	res := [][]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				res = append(res, []int{i, j})
			}
		}
	}
	for i := 0; i < len(res); i++ {
		fuckMatrix(matrix, res[i][0], res[i][1])
	}
}

func fuckMatrix(matrix [][]int, i int, j int) {
	// 先把i行全部变成0
	for k := 0; k < len(matrix[i]); k++ {
		matrix[i][k] = 0
	}
	// 将j列全部变成0
	for k := 0; k < len(matrix); k++ {
		matrix[k][j] = 0
	}
}
