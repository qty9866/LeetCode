package main

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	x := []int{2, 0, 2, 1, 1, 0}
	sortColors1(x)
	fmt.Println(x)
}
func TestSortColors2(t *testing.T) {
	x := []int{2, 0, 2, 1, 1, 0}
	sortColors2(x)
	fmt.Println(x)
}

func TestFrequency(t *testing.T) {
	ans := topKFrequent([]int{3, 0, 1, 0}, 1)
	fmt.Println(ans)
}

func TestFindKthLargest(t *testing.T) {
	largest := findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
	fmt.Println(largest)
}

func TestMerge(t *testing.T) {
	i := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	fmt.Println(i)
}
