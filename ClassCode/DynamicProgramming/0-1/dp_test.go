package DynamicProgramming

import "testing"

func TestSquare(t *testing.T) {
	numSquares(5)
}

func TestOneAndZeros(t *testing.T) {
	str := []string{"10", "0001", "111001", "1", "0"}
	findMaxForm(str, 5, 3)
}
