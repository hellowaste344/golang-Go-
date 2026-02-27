package main

import (
	"fmt"
	"time"
)

func heavyComputation() int {
	sum := 0
	for i := 0; i <= 1e9; i++ {
		sum += i
	}
	return sum
}

func main() {
	start := time.Now()

	total := heavyComputation()
	fmt.Printf("total: %d\n", total)

	elapsed := time.Since(start)
	fmt.Printf("Elapsed %d milliseconds\n", elapsed.Milliseconds())
}
