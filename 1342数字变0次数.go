/* 给你一个非负整数num，请你返回将它变成 0 所需要的步数。 如果当前数字是偶数，你需要把它除以 2 ；否则，减去 1 。

示例 1：

输入：num = 14
输出：6
解释：
步骤 1) 14 是偶数，除以 2 得到 7 。
步骤 2） 7 是奇数，减 1 得到 6 。
步骤 3） 6 是偶数，除以 2 得到 3 。
步骤 4） 3 是奇数，减 1 得到 2 。
步骤 5） 2 是偶数，除以 2 得到 1 。
步骤 6） 1 是奇数，减 1 得到 0 。
*/
package main

import "fmt"

func numberOfSteps(num int) int {
	count := 0
	for num != 0 {
		if num%2 == 0 {
			num = num / 2
			count++
		} else {
			num = num - 1
			count++
		}

	}
	return count
}

func main() {
	var a int
	fmt.Println("输入测试数字")
	fmt.Scan(&a)
	fmt.Printf("迭代次数： %v\n", numberOfSteps(a))
}
