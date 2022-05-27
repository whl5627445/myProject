package main

import (
	"fmt"
)

func test(a []int) []int {
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	return a
}

func main() {
	// n :=0
	for i := 0; i < 10; i++ {
		i += 2
		fmt.Println(i)
	}
}
