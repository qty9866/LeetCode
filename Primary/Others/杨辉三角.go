package main

/*
给定一个非负整数numRows，生成「杨辉三角」的前numRows行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。

示例 1:
输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

示例2:
输入: numRows = 1
输出: [[1]]
*/
func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	var res [][]int
	res = append(res, []int{1})
	res = append(res, []int{1, 1})
	for i := 2; i < numRows; i++ {
		var tmp []int
		tmp = append(tmp, 1)
		for j := 0; j < i-1; j++ {
			x := res[i-1][j] + res[i-1][j+1]
			tmp = append(tmp, x)
		}
		tmp = append(tmp, 1)
		res = append(res, tmp)
	}
	return res
}
