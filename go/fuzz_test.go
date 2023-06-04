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
                return  CountDivisors(Factorize(x))
            },
            // get len from Factorize with sieve
            sv.CountDivisors,
        }

        for i, method1 := range methods {
            got1 := method1(a)
            for j := i+1; j < len(methods); j++ {
                method2 := methods[j]
                got2 := method2(a)
                t.Logf("%d num divisors, method1: %v:%v, but method2 %v:%v", a, method1, got1, method2, got2)
		        if !reflect.DeepEqual(got1, got2) {
                    t.Errorf("%d num divisors, method1: %v:%v, but method2 %v:%v", a, method1, got1, method2, got2)
                }
            } 
	    }
    })
}
