package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestName(t *testing.T) {
	x := 6666
	a := strconv.Itoa(x)
	fmt.Println(a[0] - '0')
}

func TestTitleToNumber(t *testing.T) {
	ans := titleToNumber("ZY")
	fmt.Println(ans)
}

func TestMySqrt(t *testing.T) {
	mySqrt(5)
}

func TestDivide(t *testing.T) {
	ans := divide(7, -3)
	fmt.Println(ans)
}
