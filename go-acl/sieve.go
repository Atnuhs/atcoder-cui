package main

// EratosthenesSieve はエラトステネスの篩の実装
type EratosthenesSieve struct {
	isPrime   []bool
	minFactor []int
	mobius    []int
}

// NewSieve はO(N loglog N)でエラトステネスの篩を初期化する
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

// IsPrime はO(1)で素数かどうかを判定する
func (sv *EratosthenesSieve) IsPrime(x int) bool {
	return sv.isPrime[x]
}

// Factorize は O(Sqrt(N))で素因数分解を行う
// 返り値は素因数とその指数のPairのスライス
// 例）got, ret
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

// Mobius はO(sqrt(n))でメビウス関数を計算する
// メビウス関数は、整数nに対して以下のように定義される
// 0 <= n: nが平方数で割り切れる場合
// 1 or -1 <= (-1)^k: nがk個の異なる素因数を持つ場合
// 具体的には以下のような値となる
// 0 <= 4, 12, 18, 50: 平方数で割り切れる
// 1 <= 1, 6, 210: 偶数個の素因数を持つ
// -1 <= 2, 30, 140729 : 奇数個の素因数を持つ
// 約数系包除原理で使う
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
