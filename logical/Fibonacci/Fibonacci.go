package Fibonacci

func Fibonacci(n int) int {
	if (n == 1 || n == 2) {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
