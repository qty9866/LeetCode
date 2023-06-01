package main

import "fmt"

func main() {
	fmt.Println("ss")
	var b [][]int
	fmt.Scan(&b)
	i := maximumWealth(b)
	fmt.Printf("i: %v\n", i)
}
func maximumWealth(accounts [][]int) int {
	max := 0
	sum := 0
	for _, a := range accounts {
		for _, c := range a {
			sum += c
		}
		if sum >= max {
			max = sum
		}
	}

	return max
}
