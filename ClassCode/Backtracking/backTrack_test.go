package Backtracking

import (
	"fmt"
	"lettcode/ClassCode/Backtracking/SubSet"
	"testing"
)

func TestBackTrack(t *testing.T) {
	mat := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}
	SubSet.maximumRows(mat, 2)
}

func TestIsAdd(t *testing.T) {
	number := SubSet.isAdditiveNumber("112358")
	fmt.Println(number)
}
