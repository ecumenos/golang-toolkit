package numbertools

// Fibonacci generate fibonacci number
func Fibonacci(n, max int) int {
	if n <= 1 {
		return 1
	}

	var n2, n1 = 0, 1
	for i := 2; i < n; i++ {
		n2, n1 = n1, n1+n2
		if (n2 + n1) > max {
			return max
		}
	}

	return n2 + n1
}
