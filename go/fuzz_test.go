package main

import (
	"reflect"
	"sort"
	"testing"
)

func FuzzGcd(f *testing.F) {
	f.Add(0, 100000)
	f.Add(100000, 0)
	f.Add(0, 0)
	f.Add(12345678, 1000000007)
	f.Add(-12345678, 1000000007)
	f.Add(-12345678, 0)
	f.Add(0, -12345678)
	f.Fuzz(func(f *testing.T, x, y int) {
		_ = gcd(x, y)
	})
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

		if gcd(x, mod) != 1 {
			return
		}

		invX := inv(x, mod)
		got := (invX * x) % mod

		if got != 1 {
			f.Errorf("expected 1, but got %d, x: %d, mod: %d, invX: %d", got, x, mod, invX)
		}
	})
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


