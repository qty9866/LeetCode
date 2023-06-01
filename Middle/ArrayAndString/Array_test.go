package ArrayAndString

import (
	"fmt"
	"testing"
)

func TestSetZero(t *testing.T) {
	x := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(x)
	fmt.Println(x)
}

func TestAnagrams(t *testing.T) {
	x := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagrams := groupAnagrams(x)
	fmt.Println(anagrams)
}

func TestLongestSubstring(t *testing.T) {
	x3 := "pwwkew"
	n3 := lengthOfLongestSubstring(x3)
	fmt.Println(n3)
}

func TestLongestPalindrome(t *testing.T) {
	palindrome := longestPalindrome("babad")
	fmt.Println(palindrome)
}

func TestLongestPalindrome1(t *testing.T) {
	palindrome := longestPalindrome1("babad")
	fmt.Println(palindrome)
}

func TestIncreasingTriplet(t *testing.T) {
	triplet1 := increasingTriplet1([]int{1, 2, 3, 4, 5})
	triplet2 := increasingTriplet1([]int{5, 4, 3, 2, 1})
	triplet3 := increasingTriplet1([]int{2, 1, 5, 0, 4, 6})
	fmt.Println(triplet1)
	fmt.Println(triplet2)
	fmt.Println(triplet3)
}
