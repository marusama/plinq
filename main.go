package main

import (
	"fmt"
	"strconv"
)

func main() {
	m1 := make([]int, 0)
	for i := 0; i < 100; i++ {
		m1 = append(m1, i+1)
	}

	result1 := make([]string, 0)
	FromSlice(m1).
		Where(func(x int) bool { return x%2 == 0 }).
		Select(func(x int) int { return x * x }).
		Select(func(x int) string { return strconv.Itoa(x) }).
		Take(5).ToSlice(&result1)
	fmt.Println(result1)

}
