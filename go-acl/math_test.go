package main

import "testing"

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
			invX := Inv(tc.x, tc.p)

			// x * invX = should be 1
			got := (tc.x * invX) % tc.p
			if got != 1 {
				t.Errorf("actual should be 1 but got %d, invX: %d", got, invX)
			}
		})
	}
}

func FuzzInv(f *testing.F) {
	f.Add(4, 7)
	f.Add(1000000007, 97)
	f.Add(97, 1000000007)
	f.Add(4, 1)
	f.Add(4, -11)
	f.Fuzz(func(f *testing.T, x, mod int) {
		if !IsPrime(mod) || mod <= 1 {
			return
		}

		if Gcd(x, mod) != 1 {
			return
		}

		invX := Inv(x, mod)
		got := (invX * x) % mod

		if got != 1 {
			f.Errorf("expected 1, but got %d, x: %d, mod: %d, invX: %d", got, x, mod, invX)
		}
	})
}

func TestGcd(t *testing.T) {
	testCases := []struct {
		desc       string
		x, y, want int
	}{
		{desc: "gcd(2, 2) => 2", x: 2, y: 2, want: 2},
		{desc: "gcd(4, 2) => 2", x: 4, y: 2, want: 2},
		{desc: "gcd(4, 6) => 2", x: 4, y: 6, want: 2},
		{desc: "gcd(11, 13) => 1", x: 11, y: 13, want: 1},
		{desc: "gcd(11, 13) => 1", x: 11, y: 13, want: 1},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := Gcd(tc.x, tc.y)
			if tc.want != got {
				t.Errorf("expected %d but got %d", tc.want, got)
			}
		})
	}
}

func FuzzGcd(f *testing.F) {
	f.Add(0, 100000)
	f.Add(100000, 0)
	f.Add(0, 0)
	f.Add(12345678, 1000000007)
	f.Add(-12345678, 1000000007)
	f.Add(-12345678, 0)
	f.Add(0, -12345678)
	f.Fuzz(func(f *testing.T, x, y int) {
		_ = Gcd(x, y)
	})
}
