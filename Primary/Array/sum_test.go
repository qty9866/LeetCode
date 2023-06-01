package main

import (
	"fmt"
	"testing"
)

func TestSum2(t *testing.T) {
	x := []int{3, 2, 1, 4, 8}
	sum2 := TwoSum2(x, 7)
	fmt.Println(sum2)
	fmt.Println(x[1:])
}
