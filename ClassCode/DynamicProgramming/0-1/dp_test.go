package DynamicProgramming

import (
	"lettcode/ClassCode/DynamicProgramming/completely"
	"testing"
)

func TestSquare(t *testing.T) {
	DynamicProgramming.numSquares(5)
}

func TestOneAndZeros(t *testing.T) {
	str := []string{"10", "0001", "111001", "1", "0"}
	findMaxForm(str, 5, 3)
}
