package main

func main() {
	println(Fibonacci(10))
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
