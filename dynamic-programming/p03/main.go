package main

func main() {
	println(Fibonacci(10))
	println(FibonacciAdvance(10, make([]int, 10+1)))
}

func Fibonacci(i int) int {
	if i == 0 || i == 1 {
		return i
	}

	if i > 1 {
		return Fibonacci(i-1) + Fibonacci(i-2)
	}

	return 0
}

func FibonacciAdvance(i int, memo []int) int {
	if i == 0 || i == 1 {
		return i
	}

	if i > 1 {
		memo[i] = Fibonacci(i-1) + Fibonacci(i-2)
		return memo[i]
	}

	return 0
}
