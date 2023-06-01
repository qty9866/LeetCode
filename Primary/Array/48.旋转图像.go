package main

import "fmt"

/*
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]

这里我准备实现两个解法，一个原地一个非原地
*/

// 暴力求解 非原地算法
// 时间复杂度O(N^2) 空间复杂度O(N^2)
func rotateImage1(matrix [][]int) {
	length := len(matrix)
	tmp := make([][]int, length, length)
	for i := range tmp {
		tmp[i] = make([]int, length)
	}
	for i := 0; i < length; i++ {
		for k, v := range matrix[i] {
			tmp[k][length-1-i] = v
		}
	}
	copy(matrix, tmp)
	fmt.Println(matrix)
}

// 通过规律 原地算法
func rotateImage2(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] =
				matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
		}
	}
}
