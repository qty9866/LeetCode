package DoublePointer

import (
	"fmt"
	"testing"
)

func TestSubarrayProductLessThanK(t *testing.T) {
	x := []int{10, 5, 2, 6}
	ans := numSubarrayProductLessThanK(x, 100)
	ans1 := numSubarrayProductLessThanK1(x, 100)
	fmt.Println(ans)
	fmt.Println(ans1)
}

func TestThreeSum(t *testing.T) {
	ans := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(ans)
}

func TestMaxArea(t *testing.T) {
	ans := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(ans)
}

func TestTrap(t *testing.T) {
	ans := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Println(ans)
}
