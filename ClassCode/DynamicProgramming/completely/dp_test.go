package DynamicProgramming

import (
	"fmt"
	"testing"
)

/*
示例 1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
*/
func TestCoin(t *testing.T) {
	coins := []int{1, 2, 5}
	CoinChange(coins, 11)
}

func TestChange(t *testing.T) {
	coins := []int{1, 2, 5}
	ans := change2(5, coins)
	fmt.Println(ans)
}
