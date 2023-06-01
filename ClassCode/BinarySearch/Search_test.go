package main

import (
	"fmt"
	"testing"
)

func TestSearchNums(t *testing.T) {
	x := []int{5, 7, 7, 8, 8, 10}
	answer := searchRange(x, 8)
	fmt.Println(answer)
}
