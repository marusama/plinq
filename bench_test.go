package main

import (
	"testing"

	"github.com/ahmetb/go-linq"
)

func BenchmarkFor(b *testing.B) {

	m1 := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		m1 = append(m1, i+1)
	}

	b.ResetTimer()
	result1 := make([]int, 100)
	for n := 0; n < b.N; n++ {
		result1 = result1[0:0]
		for _, x := range m1 {
			if x%2 == 0 {
				y := x * x
				result1 = append(result1, y)
				if len(result1) == 10 {
					break
				}
			}
		}
	}
}

//func BenchmarkMy(b *testing.B) {
//
//	m1 := make([]int, 0, 100)
//	for i := 0; i < 100; i++ {
//		m1 = append(m1, i+1)
//	}
//
//	b.ResetTimer()
//	result1 := make([]int, 100)
//	for n := 0; n < b.N; n++ {
//		result1 = result1[0:0]
//		FromSlice(m1).
//			Where(func(x int) bool { return x%2 == 0 }).
//			Select(func(x int) int { return x * x }).
//			Take(10).ToSlice(&result1)
//	}
//}

func BenchmarkMyX(b *testing.B) {

	m1 := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		m1 = append(m1, i+1)
	}

	b.ResetTimer()
	result1 := make([]int, 100)
	for n := 0; n < b.N; n++ {
		result1 = result1[0:0]
		FromIntSlice(m1).
			WhereX(func(x interface{}) bool { return x.(int)%2 == 0 }).
			SelectX(func(x interface{}) interface{} { y := x.(int); return y * y }).
			Take(10).ToSlice(&result1)
	}
}

//func BenchmarkGoLinqT(b *testing.B) {
//
//	m1 := make([]int, 0, 100)
//	for i := 0; i < 100; i++ {
//		m1 = append(m1, i+1)
//	}
//
//	b.ResetTimer()
//	result1 := make([]string, 100)
//	for n := 0; n < b.N; n++ {
//		result1 = result1[0:0]
//		linq.From(m1).
//			WhereT(func(x int) bool { return x%2 == 0 }).
//			SelectT(func(x int) int { return x * x }).
//			SelectT(func(x int) string { return strconv.Itoa(x) }).
//			Take(10).ToSlice(&result1)
//	}
//}

func BenchmarkGoLinq(b *testing.B) {

	m1 := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		m1 = append(m1, i+1)
	}

	b.ResetTimer()
	result1 := make([]int, 100)
	for n := 0; n < b.N; n++ {
		result1 = result1[0:0]
		linq.From(m1).
			Where(func(x interface{}) bool { return x.(int)%2 == 0 }).
			Select(func(x interface{}) interface{} { y := x.(int); return y * y }).
			Take(10).ToSlice(&result1)
	}
}
