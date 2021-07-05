package main

import "fmt"

func GetMinCoinCount(total int, values []int, valueCount int) int {
	rest := total
	count := 0

	for i := 0; i < valueCount; i++ {
		currentCount := rest / values[i]
		rest -= currentCount * values[i]
		count += currentCount
	}

	if rest == 0 {
		return count
	}

	return -1
}

func main() {
	fmt.Println(GetMinCoinCount(11, []int{5, 3}, 2))
}
