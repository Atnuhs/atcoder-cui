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

			// x * invX = should be 1
			got := (tc.x * invX) % tc.p
			if got != 1 {
				t.Errorf("actual should be 1 but got %d, invX: %d", got, invX)
			}

		})
	}
}

func TestPopBack(t *testing.T) {
	testCases := []struct {
		desc string
		data []int
	}{
		{
			desc: "[1,2,3,4,5]",
			data: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			n := len(tc.data)
			a := make([]int, n)
			copy(a, tc.data)

			for i := n - 1; i >= 0; i-- {
				got := popBack(&a)
				want := tc.data[i]
				if want != got {
                    t.Errorf("index: %d, expected %d, but got %d",i, want, got)
				}
			}
		})
	}
}

func TestPopFront(t *testing.T) {
	testCases := []struct {
		desc string
		data []int
	}{
		{
			desc: "[1,2,3,4,5]",
			data: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			n := len(tc.data)
			a := make([]int, n)
			copy(a, tc.data)

			for i := 0; i < n; i++ {
				got := popFront(&a)
				want := tc.data[i]
				if want != got {
                    t.Errorf("index: %d, expected %d, but got %d",i, want, got)
				}
			}

		})
	}
}
