package main

import "fmt"

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

输入：nums = [0,1]
输出：[[0,1],[1,0]]

输入：nums = [1]
输出：[[1]]
*/

// PrintAll k代表需要处理的子数组的长度
func PrintAll(a []int, x int, k int) {
	if k == 1 {
		for i := 0; i < x; i++ {
			fmt.Print(a[i], " ")
		}
		fmt.Println()
	}
	for i := 0; i < k; i++ {
		a[i], a[k-1] = a[k-1], a[i]

		PrintAll(a, x, k-1)
		a[i], a[k-1] = a[k-1], a[i]
	}
}

//func main() {
//	a := []int{1, 2, 3, 4}
//	PrintAll(a, 4, 4)
//	//x = xibao()
//}

func permute(nums []int) [][]int {
	x := order(nums, len(nums), len(nums))
	fmt.Println(x)
	return x
}

func order(a []int, n, k int) [][]int {
	var x [][]int
	if k == 1 {
		var tmp []int
		for i := 0; i < n; i++ {
			tmp = append(tmp, a[i])
		}
		x = append(x, tmp)
	}
	for i := 0; i < k; i++ {
		a[i], a[k-1] = a[k-1], a[i]

		order(a, n, k-1)
		a[i], a[k-1] = a[k-1], a[i]
	}
	return x
}

func xibao(n int) int {
	if n == 0 {
		return 1
	}
	if n < 3 {
		return 2 * xibao(n-1)
	}
	return 2*xibao(n-1) - xibao(n-3)
}
