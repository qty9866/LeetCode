package BackTracking

import (
	"fmt"
	"testing"
)

func TestLetterCombination(t *testing.T) {
	x := letterCombinations("23")
	fmt.Println(x)
}

func TestCombine(t *testing.T) {
	ans := combine(3, 2)
	fmt.Println(ans)
}

func TestGenerate(t *testing.T) {
	ans := generateParenthesis(3)
	fmt.Println(ans)
}

func TestPermute(t *testing.T) {
	ans := permute([]int{1, 2, 3})
	fmt.Println(ans)
}

func TestSubsets(t *testing.T) {
	i := subsets([]int{1, 2})
	fmt.Println(i)
}
