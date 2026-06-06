package acl

import "math"

// Sieve は0..n の各整数が素数かどうかを保持するテーブル
type Sieve []bool

// NewSieve は O(N loglog N) でエラトステネスの篩によりテーブルを生成する
func NewSieve(n int) Sieve {
	isPrime := L1[bool](n + 1)
	for i := range isPrime {
		isPrime[i] = true
	}

	isPrime[0] = false
	isPrime[1] = false
	for i := 2; i*i <= n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i * i; j <= n; j += i {
			isPrime[j] = false
		}
	}
	return Sieve(isPrime)
}

// IsPrime は x が素数かどうかを O(1) で返す
func (sv Sieve) IsPrime(x int) bool {
	return sv[x]
}

// Primes はテーブル中の素数を昇順で列挙する
// 容量は素数定理に基づく上界で確保している
func (sv Sieve) Primes() []int {
	n := len(sv) - 1
	m := func() int {
		if n < 17 {
			return [...]int{0, 0, 1, 2, 2, 3, 3, 4, 4, 4, 4, 5, 5, 6, 6, 6, 6}[n]
		}
		return int(1.3*float64(n)/math.Log(float64(n))) + 1
	}()
	ret := make([]int, 0, m)
	for i := 2; i <= n; i++ {
		if sv[i] {
			ret = append(ret, i)
		}
	}
	return ret
}

// MinFactorTable は各値 x に対しその最小素因数を保持するテーブル
type MinFactorTable []int

// NewMinFactor は O(N loglog N) でエラトステネスの篩によりテーブルを生成する
func NewMinFactor(n int) MinFactorTable {
	mf := make(MinFactorTable, n+1)
	mf[1] = 1
	for p := 2; p <= n; p++ {
		if mf[p] == 0 {
			mf[p] = p
			// p*p のオーバーフロー回避のため除算で比較する
			if p > n/p {
				continue
			}
			for j := p * p; j <= n; j += p {
				if mf[j] == 0 {
					mf[j] = p
				}
			}
		}
	}
	return mf

}

// IsPrime は x が素数かどうかを O(1) で返す
func (mf MinFactorTable) IsPrime(x int) bool {
	if x == 1 {
		return false
	}
	return mf[x] == x
}

// Factorize は O(log x) で素因数分解を行う
// 返り値は (素因数, 指数) のペアの昇順スライス
// 例) 6 => []Pair{{2,1}, {3,1}}
func (mf MinFactorTable) Factorize(x int) []Pair[int, int] {
	ret := make([]Pair[int, int], 0)
	n := x
	for n > 1 {
		p := mf[n]
		exp := 0

		for mf[n] == p {
			n /= p
			exp++
		}
		ret = append(ret, NewPair(p, exp))
	}
	return ret
}

// Divisors は x の約数を列挙する (O(log x + d(x)))
// 例) 2 => [1, 2]
//
//	10 => [1, 2, 5, 10]
//
// 返り値の順序は約数の昇順ではない点に注意
func (mf *MinFactorTable) Divisors(x int) []int {
	ret := []int{1}

	f := mf.Factorize(x)
	for _, pe := range f {
		n := len(ret)
		for i := 0; i < n; i++ {
			v := 1
			for j := 0; j < pe.V; j++ {
				v *= pe.U
				ret = append(ret, ret[i]*v)
			}
		}
	}
	return ret
}

// CountDivisors は x の約数の個数を O(log x) で返す
// 例) 1 => 1
//
//	2 => 2
//	10 => 4
func (mf *MinFactorTable) CountDivisors(x int) int {
	return CountDivisors(mf.Factorize(x))
}

// MobiusTable はメビウス関数 μ(x) の値を保持するテーブル
// μ(n) は次のように定義される
//
//	μ(n) =  0  : n が平方因子をもつ場合 (例: 4, 12, 18, 50)
//	μ(n) = +1  : n が偶数個の相異なる素因数の積の場合 (例: 1, 6, 210)
//	μ(n) = -1  : n が奇数個の相異なる素因数の積の場合 (例: 2, 30, 140729)
//
// 約数系の包除原理で用いる
type MobiusTable []int

// NewMobiusTable は MinFactorTable 経由で μ(0..n) を O(N loglog N) で構築する
func NewMobiusTable(n int) MobiusTable {
	mf := NewMinFactor(n)
	mu := L1[int](n + 1)
	for i := range mu {
		mu[i] = 1
	}

	mu[1] = 1
	for x := 2; x <= n; x++ {
		p := mf[x]
		y := x / p
		if y%p == 0 {
			mu[x] = 0
		} else {
			mu[x] = -mu[y]
		}
	}
	return MobiusTable(mu)
}

// Mobius は μ(x) を O(1) で返す
func (mu MobiusTable) Mobius(x int) int {
	return mu[x]
}

// SegmentedSieve は区間 [lo, hi] に含まれる整数が素数かどうかを保持する区間篩
type SegmentedSieve struct {
	start   int
	isPrime Sieve
}

// NewSegmentedSieve は区間 [lo, hi] の区間篩を構築する
// D = hi - lo として、構築は O(sqrt(hi) loglog hi + D loglog hi)
// hi が大きく D が比較的小さい場合に有効
func NewSegmentedSieve(lo, hi int) *SegmentedSieve {
	if lo >= hi {
		return nil
	}
	sqrtHi := Sqrt(hi) + 1
	sv1 := NewSieve(sqrtHi)

	m := hi - lo + 1
	isPrime := L1[bool](m)
	for i := range isPrime {
		isPrime[i] = true
	}

	for p := range sv1 {
		if !sv1.IsPrime(p) {
			continue
		}

		// l は lo 以上の最小のpの倍数
		l := ((lo + p - 1) / p) * p
		if l == p {
			l = p * p
		}
		for i := l; i <= hi; i += p {
			isPrime[i-lo] = false
		}
	}
	return &SegmentedSieve{
		start:   lo,
		isPrime: isPrime,
	}
}

// IsPrime は x が素数かどうかを O(1) で返す (x は構築時の区間 [lo, hi] に含まれること)
func (sv *SegmentedSieve) IsPrime(x int) bool {
	i := x - sv.start
	return sv.isPrime[i]
}
