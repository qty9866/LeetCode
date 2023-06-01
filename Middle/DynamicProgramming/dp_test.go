package DynamicProgramming

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	jump := canJump1([]int{2, 3, 1, 1, 4})
	fmt.Println(jump)
}

func TestLongest(t *testing.T) {
	x := []int{10, 9, 2, 5, 3, 7, 101, 18}
	lis := lengthOfLIS(x)
	fmt.Println(lis)
}
