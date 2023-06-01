package ArrayAndString

/*
给你一个 m 行 n 列的矩阵matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
*/
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	visited := make([][]bool, m)
	for index := range visited {
		visited[index] = make([]bool, n)
	}
	Traversal()
}

func Traversal(matrix [][]int) {

}
