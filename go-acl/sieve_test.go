package main

import (
	"reflect"
	"sort"
	"testing"
)

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

func FuzzEratosthenesSieve_IsPrime(f *testing.F) {
	maxX := 3 * 100000
	sv := NewSieve(maxX)
	f.Add(0)
	f.Add(1)
	f.Add(9)
	f.Add(123456)
	f.Add(1000000007)
	f.Fuzz(func(t *testing.T, a int) {
		if 1 > a || a > maxX {
			return
		}

		// Perform prime number determination in two different ways
		// method1 O(sqrt(N))
		ret1 := IsPrime(a)
		// method2 EratosthenesSieve
		ret2 := sv.IsPrime(a)
		if ret1 != ret2 {
			t.Errorf("%d is Prime?, method1: %t, but method2 %t", a, ret1, ret2)
		}
	})
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

func FuzzEratosthenesSieve_Factorize(f *testing.F) {
	maxX := 3 * 100000
	sv := NewSieve(maxX)
	f.Add(0)
	f.Add(1)
	f.Add(9)
	f.Add(123456)
	f.Add(1000000007)
	f.Fuzz(func(t *testing.T, a int) {
		if 1 > a || a > maxX {
			return
		}
		// Perform Factorization in two different ways
		// method1 O(sqrt(N))
		ret1 := Factorize(a)
		// method2 EratosthenesSieve
		ret2 := sv.Factorize(a)
		if !reflect.DeepEqual(ret1, ret2) {
			t.Errorf("factorize %d, method1: %v, but method2 %v", a, ret1, ret2)
		}
	})
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

func FuzzEratosthenesSieve_Divisors(f *testing.F) {
	maxX := 3 * 100000
	sv := NewSieve(maxX)
	f.Add(0)
	f.Add(1)
	f.Add(9)
	f.Add(123456)
	f.Add(1000000007)
	f.Fuzz(func(t *testing.T, a int) {
		if 1 > a || a > maxX {
			return
		}
		// Enumerate divisors in two different ways
		// method1 O(sqrt(N))
		ret1 := Divisors(a)
		// method2 EratosthenesSieve
		ret2 := sv.Divisors(a)
		sort.Ints(ret1)
		sort.Ints(ret2)
		if !reflect.DeepEqual(ret1, ret2) {
			t.Errorf("enumerate %d divisors, method1: %v, but method2 %v", a, ret1, ret2)
		}
	})
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

func FuzzEratosthenesSieve_Mobius(f *testing.F) {
	maxX := 3 * 100000
	sv := NewSieve(maxX)
	f.Add(0)
	f.Add(1)
	f.Add(9)
	f.Add(123456)
	f.Add(1000000007)
	f.Fuzz(func(t *testing.T, a int) {
		if 1 > a || a > maxX {
			return
		}
		// Mobius function in two different ways
		// method1 O(sqrt(N))
		ret1 := Mobius(a)
		// method2 EratosthenesSieve
		ret2 := sv.Mobius(a)
		if ret1 != ret2 {
			t.Errorf("%d mobius function, method1: %v, but method2 %v", a, ret1, ret2)
		}
	})
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

func FuzzCountDivisors(f *testing.F) {
	maxX := 3 * 100000
	sv := NewSieve(maxX)
	f.Add(0)
	f.Add(1)
	f.Add(9)
	f.Add(123456)
	f.Add(1000000007)
	f.Fuzz(func(t *testing.T, a int) {
		if 1 > a || a > maxX {
			return
		}
		// enumerate divisors counting methods
		type methodFunc func(x int) int
		methods := []methodFunc{
			// get len enumerate divisors
			func(x int) int {
				return len(Divisors(x))
			},
			// get len enumerate divisors with eratosthenes sieve
			func(x int) int {
				return len(sv.Divisors(x))
			},
			// get len from Factiroze
			func(x int) int {
				return CountDivisors(Factorize(x))
			},
			// get len from Factorize with sieve
			sv.CountDivisors,
		}

		for i, method1 := range methods {
			got1 := method1(a)
			for j := i + 1; j < len(methods); j++ {
				method2 := methods[j]
				got2 := method2(a)
				t.Logf(
					"%d num divisors, method1: %v:%v, but method2 %v:%v",
					a,
					method1,
					got1,
					method2,
					got2,
				)
				if !reflect.DeepEqual(got1, got2) {
					t.Errorf(
						"%d num divisors, method1: %v:%v, but method2 %v:%v",
						a,
						method1,
						got1,
						method2,
						got2,
					)
				}
			}
		}
	})
}
