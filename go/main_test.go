package main

import (
	"testing"
)

func TestInv(t *testing.T) {
	testCases := []struct {
		desc string
		x    int
		p    int
	}{
		{desc: "x:20, p:7", x: 20, p: 7},
		{desc: "x:1234567, p:10^9+7", x: 1234567, p: 1000000007},
		{desc: "x:1234567, p:998244353", x: 1234567, p: 998244353},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			invX := inv(tc.x, tc.p)
			actual := (tc.x * invX) % tc.p
			if actual != 1 {
				t.Errorf("actual should be 1 but got %d, invX: %d", actual, invX)
			}

		})
	}
}

func benchmarkPushBack(n int, b *testing.B) {
	arr := NewVector(0)
	b.ResetTimer()
	for i := 0; i < n; i++ {
		arr.PushBack(i)
	}
}
func BenchmarkPushBack1000(b *testing.B) {
	benchmarkPushBack(1000, b)
}
func BenchmarkPushBack10000(b *testing.B) {
	benchmarkPushBack(10000, b)
}
func BenchmarkPushBack100000(b *testing.B) {
	benchmarkPushBack(100000, b)
}
func BenchmarkPushBack1000000(b *testing.B) {
	benchmarkPushBack(1000000, b)
}

func benchmarkPushFront(n int, b *testing.B) {
	arr := NewVector(0)
	b.ResetTimer()
	for i := 0; i < n; i++ {
		arr.PushFront(i)
	}
}
func BenchmarkPushFront1000(b *testing.B) {
	benchmarkPushFront(1000, b)
}
func BenchmarkPushFront10000(b *testing.B) {
	benchmarkPushFront(10000, b)
}
func BenchmarkPushFront100000(b *testing.B) {
	benchmarkPushFront(100000, b)
}

func benchmarkPopBack(n int, b *testing.B) {
	arr := NewVector(n)
	for i := 0; i < n; i++ {
		arr.PushBack(i)
	}

	b.ResetTimer()
	for arr.Len() > 0 {
		arr.PopBack()
	}
}

func BenchmarkPopBack1000(b *testing.B) {
	benchmarkPopBack(1000, b)
}
func BenchmarkPopBack10000(b *testing.B) {
	benchmarkPopBack(10000, b)
}
func BenchmarkPopBack100000(b *testing.B) {
	benchmarkPopBack(100000, b)
}
func BenchmarkPopBack1000000(b *testing.B) {
	benchmarkPopBack(1000000, b)
}

func benchmarkPopFront(n int, b *testing.B) {
	arr := NewVector(n)
	for i := 0; i < n; i++ {
		arr.PushBack(i)
	}

	b.ResetTimer()
	for arr.Len() > 0 {
		arr.PopFront()
	}
}

func BenchmarkPopFront1000(b *testing.B) {
	benchmarkPopFront(1000, b)
}
func BenchmarkPopFront10000(b *testing.B) {
	benchmarkPopFront(10000, b)
}
func BenchmarkPopFront100000(b *testing.B) {
	benchmarkPopFront(100000, b)
}
func BenchmarkPopFront1000000(b *testing.B) {
	benchmarkPopFront(1000000, b)
}
