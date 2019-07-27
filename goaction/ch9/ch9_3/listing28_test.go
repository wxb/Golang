package ch9_3_test

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {

	number := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
	b.StopTimer()
}

func BenchmarkItoa(b *testing.B) {
	number := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
	b.StopTimer()
}

func BenchmarkFormat(b *testing.B) {
	number := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
	b.StopTimer()
}
