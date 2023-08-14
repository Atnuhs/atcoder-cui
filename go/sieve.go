package main

// Eratosthenes sieve
type EratosthenesSieve struct {
	isPrime   []bool
	minFactor []int
	mobius    []int
}

// NewSieve is O(N loglog N)
func NewSieve(n int) *EratosthenesSieve {
	isPrime := make([]bool, n+1)
	minFactor := make([]int, n+1)
	mobius := make([]int, n+1)

	for i := range isPrime {
		isPrime[i] = true
		minFactor[i] = -1
		mobius[i] = 1
	}

	isPrime[0] = false
	isPrime[1] = false
	minFactor[1] = 1

	// sieve
	for i := range isPrime {
		if !isPrime[i] {
			continue
		}

		minFactor[i] = i
		mobius[i] = -1

		for j := i * 2; j <= n; j += i {
			isPrime[j] = false

			if minFactor[j] == -1 {
				minFactor[j] = i
			}

			if (j/i)%i == 0 {
				mobius[j] = 0
			} else {
				mobius[j] = -mobius[j]
			}
		}
	}
	return &EratosthenesSieve{
		isPrime:   isPrime,
		minFactor: minFactor,
		mobius:    mobius,
	}
}

// IsPrime is O(1)
func (sv *EratosthenesSieve) IsPrime(x int) bool {
	return sv.isPrime[x]
}

// Factorize is O(Sqrt(1))
// got, ret
// 6, []Pair{{2,1}, {3.1}}
func (sv *EratosthenesSieve) Factorize(x int) []*Pair[int] {
	ret := make([]*Pair[int], 0)
	n := x
	for n > 1 {
		p := sv.minFactor[n]
		exp := 0

		for sv.minFactor[n] == p {
			n /= p
			exp++
		}
		ret = append(ret, NewPair(p, exp))
	}
	return ret
}

// Mobius is O(1) return
// 0 <= 4, 12, 18, 50
// 1 <= 1, 6, 210
// -1 <= 2, 30, 140729
func (sv *EratosthenesSieve) Mobius(x int) int {
	return sv.mobius[x]
}

// Divisors is O(sqrt(n)) returns
// 2 => 1, 2
// 10 => 1, 2, 5, 10
func (sv *EratosthenesSieve) Divisors(x int) []int {
	ret := []int{1}

	f := sv.Factorize(x)
	for _, pe := range f {
		n := len(ret)
		for i := 0; i < n; i++ {
			v := 1
			for j := 0; j < pe.v; j++ {
				v *= pe.u
				ret = append(ret, ret[i]*v)
			}
		}
	}
	return ret
}

// CountDivisors is O(1) returns len(sv.Divisors(x))
// 1 => 1
// 2 => 2
// 10 => 4
func (sv *EratosthenesSieve) CountDivisors(x int) int {
	return CountDivisors(sv.Factorize(x))
}
