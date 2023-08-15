package main

import (
	"reflect"
	"sort"
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
			invX := Inv(tc.x, tc.p)

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
				got := PopBack(&a)
				want := tc.data[i]
				if want != got {
					t.Errorf("index: %d, expected %d, but got %d", i, want, got)
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
				got := PopFront(&a)
				want := tc.data[i]
				if want != got {
					t.Errorf("index: %d, expected %d, but got %d", i, want, got)
				}
			}
		})
	}
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

func TestEratosthenesSieve_IsPrime(t *testing.T) {
	maxX := 3 * 100000
	testCases := []struct {
		desc string
		x    int
		want bool
	}{
		{desc: "one", x: 1, want: false},
		{desc: "small prime", x: 2, want: true},
		{desc: "small not prime", x: 4, want: false},
		{desc: "prime?", x: 57, want: false},
		{desc: "large prime", x: 104729, want: true},
		{desc: "large not prime", x: 111111, want: false},
	}
	sv := NewSieve(maxX)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := sv.IsPrime(tc.x)
			if got != tc.want {
				t.Errorf("%d is Prime?, expected %t, but got %t", tc.x, tc.want, got)
			}
		})
	}
}

func TestEratosthenesSieve_Factorize(t *testing.T) {
	maxX := 3 * 100000
	testCases := []struct {
		desc string
		x    int
		want []*Pair[int]
	}{
		{desc: "one", x: 1, want: []*Pair[int]{}},
		{desc: "simple prime number", x: 2, want: []*Pair[int]{NewPair(2, 1)}},
		{desc: "simple composite number", x: 12, want: []*Pair[int]{NewPair(2, 2), NewPair(3, 1)}},
		{desc: "large prime number", x: 104729, want: []*Pair[int]{NewPair(104729, 1)}},
		{
			desc: "large composite number",
			x:    1260,
			want: []*Pair[int]{NewPair(2, 2), NewPair(3, 2), NewPair(5, 1), NewPair(7, 1)},
		},
	}

	sv := NewSieve(maxX)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := sv.Factorize(tc.x)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("factorize %d result, expected %v, but got %v", tc.x, tc.want, got)
			}
		})
	}
}

func TestEratosthenesSieve_Divisors(t *testing.T) {
	maxX := 3 * 100000
	testCases := []struct {
		desc string
		x    int
		want []int
	}{
		{desc: "one", x: 1, want: []int{1}},
		{desc: "simple prime number", x: 2, want: []int{1, 2}},
		{desc: "simple composite number", x: 12, want: []int{1, 2, 3, 4, 6, 12}},
		{desc: "large prime number", x: 104729, want: []int{1, 104729}},
	}

	sv := NewSieve(maxX)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := sv.Divisors(tc.x)
			sort.Ints(got)
			sort.Ints(tc.want)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("factorize %d result, expected %v, but got %v", tc.x, tc.want, got)
			}
		})
	}
}

func TestEratosthenesSieve_Mobius(t *testing.T) {
	maxX := 3 * 100000
	testCases := []struct {
		desc string
		x    int
		want int
	}{
		{desc: "one", x: 1, want: 1},
		{desc: "simple prime number", x: 2, want: -1},
		{desc: "simple composite number", x: 6, want: 1},
		{desc: "simple composite number", x: 30, want: -1},
		{desc: "simple composite number", x: 4, want: 0},
		{desc: "simple composite number", x: 12, want: 0},
		{desc: "large prime number", x: 104729, want: -1},
	}

	sv := NewSieve(maxX)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := sv.Mobius(tc.x)

			if tc.want != got {
				t.Errorf("%d Mebius function, expected %d, but got %d", tc.x, tc.want, got)
			}
		})
	}
}

func TestCountDivisors(t *testing.T) {
	maxX := 3 * 100000
	testCases := []struct {
		desc string
		x    int
		want int
	}{
		{desc: "one", x: 1, want: 1},
		{desc: "simple prime number", x: 2, want: 2},
		{desc: "simple composite number", x: 12, want: 6},
		{desc: "large prime number", x: 104729, want: 2},
	}

	sv := NewSieve(maxX)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			f := sv.Factorize(tc.x)
			got := CountDivisors(f)

			if tc.want != got {
				t.Errorf("%d divisors num, expected %d, but got %d", tc.x, tc.want, got)
			}
		})
	}
}
