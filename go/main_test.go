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
