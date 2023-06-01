package DynamicProgramming

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	pre, cur := 1, 1
	for i := 3; i <= n; i++ {
		sum := pre + cur
		pre = cur
		cur = sum
	}
	return cur
}
