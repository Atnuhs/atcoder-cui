package main

import (
	"math"
)

// ModPow return x^e % mod
func ModPow(x, e, mod int) int {
	i := 1
	ret := 1

	for i <= e {
		if i&e > 0 {
			ret = (ret * x) % mod
		}
		i <<= 1
		x = (x * x) % mod
	}
	return ret
}

func Inv(x, mod int) int {
	return ModPow(x, mod-2, mod)
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func Lcm(a, b int) int {
	return a / Gcd(a, b) * b
}

func Sqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

// NextPerm returns [1,2,3,4] => [1,2,4,3] ... [4,3,2,1]
func NextPerm(a []int) bool {
	// search i
	i := len(a) - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := len(a) - 1
	for j >= 0 && a[j] <= a[i] {
		j--
	}

	a[i], a[j] = a[j], a[i]

	l := i + 1
	r := len(a) - 1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	return true
}

// Extrema returns min, max
func Extrema[T Ordered](vals ...T) (T, T) {
	mi, ma := vals[0], vals[0]
	for _, v := range vals {
		if v < mi {
			mi = v
		}
		if v > ma {
			ma = v
		}
	}
	return mi, ma
}

func Max[T Ordered](vals ...T) T {
	_, ma := Extrema(vals...)
	return ma
}

func Min[T Ordered](vals ...T) T {
	mi, _ := Extrema(vals...)
	return mi
}

func Sum[T Ordered](vals ...T) T {
	var sum T
	for _, v := range vals {
		sum += v
	}
	return sum
}

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

// IsPrime is O(Sqrt(N))
func IsPrime(x int) bool {
	if x == 1 {
		return false
	}

	rx := Sqrt(x)
	for i := 2; i <= rx; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// Factorize is O(Sqrt(N))
// got, ret
// 6, []Pair{{2,1}, {3.1}}
func Factorize(x int) []*Pair[int] {
	if x == 1 {
		return []*Pair[int]{}
	}

	rx := Sqrt(x)
	n := x
	ret := make([]*Pair[int], 0)
	for i := 2; i <= rx; i++ {
		if n%i != 0 {
			continue
		}
		exp := 0
		for n%i == 0 {
			n /= i
			exp++
		}
		ret = append(ret, NewPair(i, exp))
	}
	if n != 1 {
		ret = append(ret, NewPair(n, 1))
	}
	return ret
}

// Mobius is O(sqrt(n)) returns
// 0 <= 4, 12, 18, 50
// 1 <= 1, 6, 210
// -1 <= 2, 30, 140729
func Mobius(x int) int {
	ret := 1

	rx := Sqrt(x)
	n := x
	for i := 2; i <= rx; i++ {
		if n%i != 0 {
			continue
		}

		if (n/i)%i == 0 {
			return 0
		}
		n /= i
		ret = -ret
	}

	if n != 1 {
		ret = -ret
	}
	return ret
}

// Divisors is O(sqrt(n)) returns
// 2 => 1, 2
// 10 => 1, 2, 5, 10
func Divisors(x int) []int {
	ret := make([]int, 0)

	rx := Sqrt(x)
	for i := 1; i <= rx; i++ {
		if x%i != 0 {
			continue
		}
		ret = append(ret, i)
		if i != x/i {
			ret = append(ret, x/i)
		}
	}
	return ret
}

// CountDivisors is O(sqrt(n)) returns
// 1 => 1
// 2 => 2
// 10 => 4
func CountDivisors(pairs []*Pair[int]) int {
	ans := 1
	for _, pe := range pairs {
		ans *= (pe.v + 1)
	}
	return ans
}
