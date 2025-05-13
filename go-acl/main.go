package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
	"strings"
)

// lib
var (
	In  = bufio.NewReaderSize(os.Stdin, 1<<20)
	Out = bufio.NewWriterSize(os.Stdout, 1<<20)
	Dbg = bufio.NewWriterSize(os.Stderr, 1<<20)
)

const (
	MOD1 = 1000000007
	MOD2 = 998244353
	// INF is 10^18
	INF = 1000000000000000000
)

func init() {
}

func main() {
	defer Out.Flush()
}

// Ordered はconstraints.OrderedがAtCoderで使えないので、代わりに使う
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Pair は2つの値を持つ構造体
type Pair[T any] struct {
	u, v T
}

// NewPair Pairを生成する
func NewPair[T any](u, v T) *Pair[T] {
	return &Pair[T]{u, v}
}

// String Pairの文字列を、空白区切りで返す
func (p *Pair[T]) String() string {
	return fmt.Sprintf("%v %v", p.u, p.v)
}

// ILF は長さnの配列を、関数fで初期化する
func ILF[T any](n int, f func(i int) T) []T {
	ret := make([]T, n)
	for i := range ret {
		ret[i] = f(i)
	}
	return ret
}

// IGF はh行w列の配列を、関数fで初期化する
func IGF[T any](h, w int, f func(ih, iw int) T) [][]T {
	return ILF(h, func(ih int) []T {
		return ILF(w, func(iw int) T {
			return f(ih, iw)
		})
	})
}

// ILF2 は長さnの配列を、関数fで初期化する
func ILF2[T any](n int, f func() T) []T { return ILF(n, func(_ int) T { return f() }) }

// IGF2 はh行w列の配列を、関数fで初期化する
func IGF2[T any](h, w int, f func() T) [][]T {
	return IGF(h, w, func(_ int, _ int) T { return f() })
}

// IL は長さnの配列を、値vで初期化する
// vに配列を指定すると、すべてポインタが同じになる
func IL[T any](n int, v T) []T {
	return ILF(n, func(_ int) T { return v })
}

// IG はh行w列の配列を、値vで初期化する
// vに配列を指定すると、すべてポインタが同じになる
func IG[T any](h, w int, v T) [][]T {
	return IGF(h, w, func(_ int, _ int) T { return v })
}

// IJF はn行のグラフをjagged配列で初期化する
func IJ[T any](n int) [][]T {
	return ILF(n, func(_ int) []T { return make([]T, 0) })
}

// Prepend は配列の先頭に値を追加する
func Prepend[T any](arr []T, vars ...T) []T {
	return append(vars, arr...)
}

// S は文字列を読み込む
func S() string {
	var ret string
	fmt.Fscan(In, &ret)
	return ret
}

// R は文字列を[]runeとして読み込む
func R() []rune {
	return []rune(S())
}

// I は整数を読み込む
func I() int {
	var ret int
	fmt.Fscan(In, &ret)
	return ret
}

// I2 は整数を2つ読み込む
func II() (int, int) {
	var ret1, ret2 int
	fmt.Fscan(In, &ret1, &ret2)
	return ret1, ret2
}

// I3 は整数を3つ読み込む
func III() (int, int, int) {
	var ret1, ret2, ret3 int
	fmt.Fscan(In, &ret1, &ret2, &ret3)
	return ret1, ret2, ret3
}

// Is は整数をn個読み込む
func Is(n int) []int { return ILF2(n, I) }

// Iss は整数をh行w列の配列として読み込む
func Iss(h, w int) [][]int { return IGF2(h, w, I) }

// Sss は文字列をn個読み込む
func Ss(n int) []string { return ILF2(n, S) }

// Rss は文字列をn個読み込む
func Rs(n int) [][]rune { return ILF2(n, R) }

// All は配列のすべての要素が条件を満たすかどうかを判定する
func All[T any](vals []T, f func(i int, v T) bool) bool {
	for i, v := range vals {
		if !f(i, v) {
			return false
		}
	}
	return true
}

// Any は配列のいずれかの要素が条件を満たすかどうかを判定する
func Any[T any](vals []T, f func(i int, v T) bool) bool {
	for i, v := range vals {
		if f(i, v) {
			return true
		}
	}
	return false
}

// Ans は出力を行う
func Ans(args ...interface{}) {
	for i, arg := range args {
		if reflect.TypeOf(arg).Kind() == reflect.Slice {
			fmt.Fprint(Out, strings.Trim(fmt.Sprint(arg), "[]"))
			// float64の場合、小数点以下14桁まで出力
		} else if reflect.TypeOf(arg).Kind() == reflect.Float64 {
			fmt.Fprintf(Out, "%.14f", arg)
		} else {
			fmt.Fprint(Out, arg)
		}
		if i < len(args)-1 {
			fmt.Fprint(Out, " ")
		}
	}
	fmt.Fprintln(Out)
}

// Yes は"Yes"を出力する
func Yes() {
	Ans("Yes")
}

// No は"No"を出力する
func No() {
	Ans("No")
}

// YesNo は条件に応じてYesまたはNoを出力する
func YesNo(f func() bool) {
	if f() {
		Yes()
	} else {
		No()
	}
}

// YesNo2 は条件に応じてYesまたはNoを出力する
func YesNo2(b bool) {
	if b {
		Yes()
	} else {
		No()
	}
}

// PopBack はO(1)で配列の末尾を削除して返す
func PopBack[T any](a *[]T) T {
	ret := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return ret
}

// PopFront はO(1)で配列の先頭を削除して返す
func PopFront[T any](a *[]T) T {
	ret := (*a)[0]
	*a = (*a)[1:]
	return ret
}

// UnionFind はUnion-Find木の実装
type UnionFind struct {
	data []int
}

// NewUnionFind は新しいUnion-Find木を生成する
func NewUnionFind(n int) *UnionFind {
	data := make([]int, n)
	for i := range data {
		data[i] = -1
	}
	return &UnionFind{
		data: data,
	}
}

func (uf *UnionFind) Root(x int) int {
	if uf.data[x] < 0 {
		return x
	} else {
		uf.data[x] = uf.Root(uf.data[x])
		return uf.data[x]
	}
}

func (uf *UnionFind) Family(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

func (uf *UnionFind) Size(x int) int {
	return -uf.data[uf.Root(x)]
}

func (uf *UnionFind) Union(x, y int) {
	rx := uf.Root(x)
	ry := uf.Root(y)

	if rx == ry {
		return
	}

	if uf.Size(rx) < uf.Size(ry) {
		rx, ry = ry, rx
	}

	uf.data[rx] += uf.data[ry]
	uf.data[ry] = rx
}

// heaapImpl はヒープの実装
type heapImpl[T Ordered] []T

func (h heapImpl[T]) Len() int { return len(h) }
func (h heapImpl[T]) Less(i, j int) bool {
	hi, hj := h[i], h[j]
	return hi < hj
}
func (h heapImpl[T]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *heapImpl[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *heapImpl[T]) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// PQ は優先度付きキューの実装
type PQ[T Ordered] struct {
	value heapImpl[T]
}

func NewPQ[T Ordered]() *PQ[T] {
	pq := &PQ[T]{}
	heap.Init(&pq.value)
	return pq
}

func (pq *PQ[T]) Push(x T) {
	heap.Push(&pq.value, x)
}

func (pq *PQ[T]) Pop() T {
	x := heap.Pop(&pq.value)
	return x.(T)
}

// lIdx　は左の子ノードのインデックスを返す
func lIdx(idx int) int {
	return idx & ^1
}

// rIdx　は右の子ノードのインデックスを返す
func rIdx(idx int) int {
	return idx | 1
}

// pIdx　は親ノードのインデックスを返す
func pIdx(idx int) int {
	return ((idx >> 1) - 1) & ^1
}

// cIdx　は子ノードのインデックスを返す
func cIdx(idx int) int {
	return (idx & ^1)<<1 | 2 | idx&1
}

// DEPQ は二重優先度キューの実装
// 二重優先度キューは、最大値と最小値をO(logN)で取得できるデータ構造
type DEPQ[T Ordered] struct {
	values []T
}

// NewDEPQ は二重優先度キューを初期化する
func NewDEPQ[T Ordered](values ...T) *DEPQ[T] {
	pq := &DEPQ[T]{
		values: values,
	}
	for i := pq.Size() - 1; i >= 0; i-- {
		if i&1 == 1 && pq.values[i-1] < pq.values[i] {
			pq.values[i-1], pq.values[i] = pq.values[i], pq.values[i-1]
		}
		idx := pq.down(i)
		pq.upAt(idx, i)
	}
	return pq
}

func (pq *DEPQ[T]) Size() int {
	return len(pq.values)
}

func (pq *DEPQ[T]) Empty() bool {
	return len(pq.values) == 0
}

func (pq *DEPQ[T]) Push(x T) {
	pq.values = append(pq.values, x)
	pq.up(pq.Size() - 1)
}

func (pq *DEPQ[T]) GetMax() T {
	return pq.values[0]
}

func (pq *DEPQ[T]) GetMin() T {
	if pq.Size() < 2 {
		return pq.values[0]
	}
	return pq.values[1]
}

func (pq *DEPQ[T]) PopMax() T {
	ret := pq.GetMax()
	pq.values[0] = pq.values[pq.Size()-1]
	pq.values = pq.values[:pq.Size()-1]
	idx := pq.down(0)
	pq.up(idx)
	return ret
}

func (pq *DEPQ[T]) PopMin() T {
	ret := pq.GetMin()
	if pq.Size() < 2 {
		pq.values = []T{}
		return ret
	}
	pq.values[1] = pq.values[pq.Size()-1]
	pq.values = pq.values[:pq.Size()-1]
	idx := pq.down(1)
	pq.up(idx)
	return ret
}

func (pq *DEPQ[T]) upAt(idx, root int) {
	l, r := lIdx(idx), rIdx(idx)

	// sould be value[l] >= value[r]
	if r < pq.Size() && pq.values[l] < pq.values[r] {
		pq.values[l], pq.values[r] = pq.values[r], pq.values[l]
		idx ^= 1
	}

	for p := pIdx(idx); idx > root && pq.values[p] < pq.values[idx]; idx, p = p, pIdx(p) {
		// max heap
		pq.values[idx], pq.values[p] = pq.values[p], pq.values[idx]
	}

	for p := pIdx(idx) | 1; idx > root && pq.values[p] > pq.values[idx]; idx, p = p, pIdx(p)|1 {
		// min heap
		pq.values[idx], pq.values[p] = pq.values[p], pq.values[idx]
	}
}

func (pq *DEPQ[T]) up(idx int) {
	pq.upAt(idx, 1)
}

func (pq *DEPQ[T]) down(idx int) int {
	if idx&1 == 1 {
		// min heap
		for c := cIdx(idx); c < pq.Size(); idx, c = c, cIdx(c) {
			if c+2 < pq.Size() && pq.values[c] > pq.values[c+2] {
				c += 2
			}
			if pq.values[c] < pq.values[idx] {
				pq.values[idx], pq.values[c] = pq.values[c], pq.values[idx]
			}
		}
	} else {
		// max heap
		for c := cIdx(idx); c < pq.Size(); idx, c = c, cIdx(c) {
			if c+2 < pq.Size() && pq.values[c] < pq.values[c+2] {
				c += 2
			}
			if pq.values[c] > pq.values[idx] {
				pq.values[idx], pq.values[c] = pq.values[c], pq.values[idx]
			}
		}
	}
	return idx
}

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

// ModPow return x^e % mod
func ModPow(x, e, mod int) int {
	ret := 1
	for e > 0 {
		if e&1 == 1 {
			ret = (ret * x) % mod
		}
		x = (x * x) % mod
		e >>= 1
	}
	return ret
}

// Inv return x^(-1) % mod
func Inv(x, mod int) int {
	return ModPow(x, mod-2, mod)
}

// Gcd return greatest common divisor on O(log N)
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// Lcm return least common multiple on O(log N)
func Lcm(a, b int) int {
	return a / Gcd(a, b) * b
}

// Sqrt return square root of x
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

// Sum returns sum of vals
func Sum[T Ordered](vals ...T) T {
	var sum T
	for _, v := range vals {
		sum += v
	}
	return sum
}

// Abs returns absolute value of x
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

// Monoid
type (
	Operator[T any] func(x1, x2 T) T
	Monoid[T any]   struct {
		Op Operator[T]
		E  T
	}
)

// MoMax は最大値を求めるモノイド
func MoMax() *Monoid[int] {
	return &Monoid[int]{
		Op: func(x1, x2 int) int {
			return Max(x1, x2)
		},
		E: -INF,
	}
}

// MoMin は最小値を求めるモノイド
func MoMin() *Monoid[int] {
	return &Monoid[int]{
		Op: func(x1, x2 int) int {
			return Min(x1, x2)
		},
		E: INF,
	}
}

// MoSum は和を求めるモノイド
func MoSum[T int | float64]() *Monoid[T] {
	return &Monoid[T]{
		Op: func(x1, x2 T) T {
			return x1 + x2
		},
		E: 0,
	}
}

// MoXOR はXORを求めるモノイド
func MoXOR() *Monoid[int] {
	return &Monoid[int]{
		Op: func(x1, x2 int) int {
			return x1 ^ x2
		},
		E: 0,
	}
}

// MoMODMul はmodを考慮した掛け算を求めるモノイド
func MoMODMul(mod int) *Monoid[int] {
	return &Monoid[int]{
		Op: func(x1, x2 int) int {
			return (x1 * x2) % mod
		},
		E: 1,
	}
}

// SegmentTree はセグメント木の実装
type SegmentTree[T any] struct {
	data []T
	n    int
	mo   *Monoid[T]
}

// NewSegmentTree はセグメント木を初期化する
func NewSegmentTree[T any](arr []T, mo *Monoid[T]) *SegmentTree[T] {
	n := 1
	for n < len(arr) {
		n *= 2
	}

	data := ILF(2*n-1, func(i int) T { return mo.E })
	for i := range arr {
		j := i + n - 1
		data[j] = arr[i]
	}

	for i := n - 2; i >= 0; i-- {
		c1 := 2*i + 1
		c2 := 2*i + 2
		data[i] = mo.Op(data[c1], data[c2])
	}

	return &SegmentTree[T]{
		data: data,
		n:    n,
		mo:   mo,
	}
}

func (st *SegmentTree[T]) Update(i int, x T) {
	i += st.n - 1
	st.data[i] = x
	for i > 0 {
		i = (i - 1) / 2
		st.data[i] = st.mo.Op(st.data[2*i+1], st.data[2*i+2])
	}
}

func (st *SegmentTree[T]) At(i int) T {
	return st.Query(i, i+1)
}

func (st *SegmentTree[T]) Query(a, b int) T {
	return st.querySub(a, b, 0, 0, st.n)
}

func (st *SegmentTree[T]) querySub(a, b, n, l, r int) T {
	if r <= a || b <= l {
		return st.mo.E
	}

	if a <= l && r <= b {
		return st.data[n]
	}

	vl := st.querySub(a, b, 2*n+1, l, (l+r)/2)
	vr := st.querySub(a, b, 2*n+2, (l+r)/2, r)
	return st.mo.Op(vl, vr)
}

// SplayNode はスプレー木のノードを表す構造体
type SplayNode struct {
	l, r, p         *SplayNode
	size            int
	key             int
	value, min, max int
}

// NewSplayNode は新しいスプレー木のノードを生成する
func NewSplayNode(key, value int) *SplayNode {
	ret := &SplayNode{
		l:     nil,
		r:     nil,
		p:     nil,
		key:   key,
		value: value,
	}
	ret.update()
	return ret
}

func (sn *SplayNode) index() int {
	if sn == nil {
		return -1
	}
	if sn.l != nil {
		return sn.l.size
	}
	return 0
}

func (sn *SplayNode) update() {
	if sn == nil {
		return
	}
	sn.size = 1
	sn.min = sn.value
	sn.max = sn.value

	if sn.l != nil {
		sn.size += sn.l.size
		sn.min = Min(sn.min, sn.l.min)
		sn.max = Max(sn.max, sn.l.max)
	}
	if sn.r != nil {
		sn.size += sn.r.size
		sn.min = Min(sn.min, sn.r.min)
		sn.max = Max(sn.max, sn.r.max)
	}
}

func (sn *SplayNode) state() int {
	if sn.p == nil {
		return 0
	}
	if sn.p.l == sn {
		return 1
	}
	if sn.p.r == sn {
		return -1
	}
	return INF
}

func (sn *SplayNode) rotate() {
	if sn == nil {
		return
	}
	ns := sn.state()
	if ns == 0 {
		return
	}

	p := sn.p
	ps := p.state()

	// edge 1
	pp := p.p
	switch ps {
	case 1:
		pp.l = sn
	case -1:
		pp.r = sn
	}
	sn.p = pp

	// edge 2, 3
	var c *SplayNode
	switch ns {
	case 1:
		c = sn.r
		sn.r = p
		p.l = c
	case -1:
		c = sn.l
		sn.l = p
		p.r = c
	}

	p.p = sn
	if c != nil {
		c.p = p
	}
	p.update()
	sn.update()
}

func (sn *SplayNode) splay() {
	if sn == nil {
		return
	}
	for sn.p != nil {
		// sn is not root

		if sn.p.state() == 0 {
			// sn.p is root
			sn.rotate()
			continue
		}

		if sn.state() == sn.p.state() {
			sn.p.rotate()
			sn.rotate()
		} else {
			sn.rotate()
			sn.rotate()
		}
	}
}

func (sn *SplayNode) values() []int {
	ret := make([]int, 0)
	if sn == nil {
		return ret
	}
	if sn.l != nil {
		ret = append(ret, sn.l.values()...)
	}
	ret = append(ret, sn.key)
	if sn.r != nil {
		ret = append(ret, sn.r.values()...)
	}
	return ret
}

func (sn *SplayNode) String() string {
	if sn == nil {
		return ""
	}
	ret := strings.Builder{}
	ret.WriteString("(")
	if sn.l != nil {
		ret.WriteString(fmt.Sprintf("%s ", sn.l.String()))
	}
	ret.WriteString(fmt.Sprint(sn.key))
	if sn.r != nil {
		ret.WriteString(fmt.Sprintf(" %s", sn.r.String()))
	}
	ret.WriteString(")")
	return ret.String()
}

func (sn *SplayNode) describe(rank int) string {
	ret := ""
	if sn.r != nil {
		ret += sn.r.describe(rank + 1)
	}
	ret += fmt.Sprintf(
		strings.Repeat("    ", rank)+"-[k:%d, v:%d, sz: %d, rank: %d]\n",
		sn.key,
		sn.value,
		sn.size,
		rank,
	)

	if sn.l != nil {
		ret += sn.l.describe(rank + 1)
	}
	return ret
}

func (sn *SplayNode) maxRank(rank int) int {
	ret := rank
	if sn.r != nil {
		ret = Max(ret, sn.r.maxRank(rank+1))
	}
	if sn.l != nil {
		ret = Max(ret, sn.l.maxRank(rank+1))
	}
	return ret
}

func (sn *SplayNode) FindAtSub(idx int) (found *SplayNode) {
	if sn == nil {
		return nil
	}
	if idx < 0 || sn.size <= idx {
		return nil
	}
	// n include [0, n)
	now := sn
	for now != nil {
		switch {
		case idx == now.index():
			now.splay()
			return now
		case idx < now.index():
			now = now.l
		case idx > now.index():
			idx -= now.index() + 1
			now = now.r
		}
	}
	panic("must not reach this code")
}

func (sn *SplayNode) FindAt(idx int) *SplayNode {
	node := sn.FindAtSub(idx)
	node.splay()
	return node
}

func (sn *SplayNode) FindSub(key int) (found *SplayNode) {
	now := sn
	for now != nil {
		if now.key == key {
			now.splay()
			return now
		}

		if now.key > key {
			now = now.l
		} else {
			now = now.r
		}
	}
	return nil
}

func (sn *SplayNode) Find(key int) (found *SplayNode) {
	found = sn.FindSub(key)
	found.splay()
	return found
}

func (sn *SplayNode) Ge(key int) (idx int) {
	if sn == nil {
		return -1
	}
	now := sn
	idx = sn.size
	i := 0
	for now != nil {
		if now.key >= key {
			idx = Min(idx, i+now.index())
			now = now.l
		} else {
			i += now.index() + 1
			now = now.r
		}
	}
	return idx
}

func (sn *SplayNode) MergeR(rroot *SplayNode) *SplayNode {
	if rroot == nil {
		return sn
	}
	if sn == nil {
		return rroot
	}
	sn = sn.FindAt(sn.size - 1) // always found
	sn.r = rroot
	rroot.p = sn
	sn.update()
	return sn
}

func (sn *SplayNode) Split(idx int) (*SplayNode, *SplayNode) {
	if sn == nil {
		return nil, nil
	}
	if idx == sn.size {
		return sn, nil
	}

	rroot := sn.FindAt(idx)
	if rroot == nil {
		// idx is out of index
		return nil, nil
	}

	lroot := rroot.l
	if lroot != nil {
		lroot.p = nil
	}
	rroot.l = nil

	rroot.update()
	// lroot not need to update()
	return lroot, rroot
}

func (sn *SplayNode) InsertAt(idx int, node *SplayNode) *SplayNode {
	lroot, rroot := sn.Split(idx)
	if lroot == nil {
		return node.MergeR(rroot)
	} else {
		return lroot.MergeR(node).MergeR(rroot)
	}
}

func (sn *SplayNode) DeleteAt(idx int) (root *SplayNode, dropped *SplayNode) {
	lroot, rroot := sn.Split(idx)
	if rroot == nil {
		return lroot, nil
	}
	del, rroot := rroot.Split(1)
	if lroot == nil {
		return rroot, del
	} else {
		root = lroot.MergeR(rroot)
		return root, del
	}
}

func (sn *SplayNode) Insert(node *SplayNode) *SplayNode {
	if sn == nil {
		return node
	}

	idx := sn.Ge(node.key) // idx shoud be 0 <= idx <= sn.size
	if idx == sn.size {
		return sn.MergeR(node)
	}

	found := sn.FindAt(idx)
	if found.key != node.key {
		return found.InsertAt(idx, node)
	}

	return found
}

func (sn *SplayNode) Delete(node *SplayNode) (root *SplayNode, removed *SplayNode) {
	root = sn.Find(node.key)
	if root == nil {
		// target not found
		return sn, nil
	}
	if root.key == node.key {
		// target found
		root, removed = root.DeleteAt(root.index())
	}
	// target not found
	return root, removed
}

// SplaySet はスプレー木のセットを表す構造体
type SplaySet struct {
	root *SplayNode
}

// NewSplaySet は新しいスプレー木のセットを生成する
func NewSplaySet(values ...int) *SplaySet {
	s := &SplaySet{
		root: nil,
	}
	for _, v := range values {
		s.Push(v)
	}
	return s
}

func NewSplaySetNode(value int) *SplayNode {
	return NewSplayNode(value, -1)
}

func (ss *SplaySet) String() string {
	return ss.root.String()
}

func (ss *SplaySet) Push(value int) {
	ss.root = ss.root.Insert(NewSplaySetNode(value))
}

func (ss *SplaySet) Remove(value int) {
	ss.root, _ = ss.root.Delete(NewSplaySetNode(value))
}

func (ss *SplaySet) Has(value int) bool {
	found := ss.root.Find(value)
	if found == nil {
		return false
	}
	ss.root = found
	return true
}

func (ss *SplaySet) Values() (arr []int) {
	return ss.root.values()
}

func (ss *SplaySet) Size() int {
	if ss.root == nil {
		return 0
	}
	return ss.root.size
}

func (ss *SplaySet) IsEmpty() bool {
	return ss.root == nil
}

func (ss *SplaySet) At(idx int) int {
	found := ss.root.FindAt(idx)
	if found == nil {
		panic("out of index")
	}
	ss.root = found
	return ss.root.key
}

func (ss *SplaySet) Ge(value int) int {
	idx := ss.root.Ge(value)
	if 0 <= idx && idx < ss.root.size {
		ss.root = ss.root.FindAt(idx)
	}
	return idx
}

func (ss *SplaySet) Gt(value int) int {
	if ss.root == nil {
		return -1
	}
	return ss.Ge(value + 1)
}

func (ss *SplaySet) Le(value int) int {
	if ss.root == nil {
		return -1
	}
	return ss.Ge(value+1) - 1
}

func (ss *SplaySet) Lt(value int) int {
	if ss.root == nil {
		return -1
	}
	return ss.Ge(value) - 1
}

// SplayMap はスプレー木のマップを表す構造体
type SplayMap struct {
	root *SplayNode
}

// NewSplayMap は新しいスプレー木のマップを生成する
func NewSplayMap() *SplayMap {
	return &SplayMap{
		root: nil,
	}
}

func NewSplayMapNode(key int, value int) *SplayNode {
	return NewSplayNode(key, value)
}

func (ss *SplayMap) Push(key int, value int) {
	node := NewSplayMapNode(key, value)
	if ss.root == nil {
		ss.root = node
	}
	ss.root = ss.root.Insert(node)
}

func (ss *SplayMap) Remove(key int) int {
	var removed *SplayNode
	ss.root, removed = ss.root.Delete(NewSplayMapNode(key, -1))
	if removed != nil {
		return removed.value
	}
	return 0
}

func (ss *SplayMap) Has(key int) bool {
	found := ss.root.Find(key)
	if found == nil {
		return false
	}
	ss.root = found
	return true
}

func (ss *SplayMap) Values() (arr []int) {
	return ss.root.values()
}

func (ss *SplayMap) Size() int {
	return ss.root.size
}

func (ss *SplayMap) IsEmpty() bool {
	return ss.root == nil
}

func (ss *SplayMap) At(key int) int {
	found := ss.root.Find(key)
	if found == nil {
		return 0
	}
	ss.root = found
	return ss.root.value
}

func (ss *SplayMap) String() string {
	if ss.root == nil {
		return ""
	}
	return ss.root.String()
}

func (ss *SplayMap) Ge(value int) int {
	idx := ss.root.Ge(value)
	ss.root = ss.root.FindAtSub(idx)
	return ss.root.key
}

func (ss *SplayMap) Gt(value int) int {
	return ss.Ge(value + 1)
}

func (ss *SplayMap) Le(value int) int {
	return ss.Ge(value+1) - 1
}

func (ss *SplayMap) Lt(value int) int {
	return ss.Ge(value) - 1
}

// WEdge は重み付き辺を表す構造体
type WEdge struct {
	from, to, weight int
}

// NewWEdge は新しい重み付き辺を生成する
func NewWEdge(from, to, weight int) *WEdge {
	return &WEdge{
		from:   from,
		to:     to,
		weight: weight,
	}
}

// Kruskal はクラスカル法を用いて最小全域木を求める
func Kruskal(n int, edges []*WEdge) (int, []*WEdge) {
	// はじめに辺を重みでソートする
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	// その後、Union-Findを用いて最小全域木を求める
	uf := NewUnionFind(n)
	ret := make([]*WEdge, 0)
	sum := 0

	// すべての辺を調べる
	for _, e := range edges {
		if uf.Family(e.from, e.to) {
			continue
		}
		ret = append(ret, e)
		sum += e.weight
		uf.Union(e.from, e.to)
	}
	if uf.Size(0) != n {
		return -1, nil
	}
	return sum, ret
}

// Manacher は文字が与えられたとき、各iについて、
// 文字iを中心とした回文の半径を記録した配列を返す
// 例）"ababa" => [0, 1, 2, 1, 0]
// O(|S|)
// 偶数調の回文を考慮する場合は、"a$b$a$b$a"のように$を挿入すると検出できるようになる
func Manacher(s string) []int {
	m := len(s)
	rad := make([]int, m)

	i, j := 0, 0
	for i < m {
		for i-j >= 0 && i+j < m && s[i-j] == s[i+j] {
			j++
		}
		rad[i] = j
		k := 1
		for i-k >= 0 && k+rad[i-k] < j {
			rad[i+k] = rad[i-k]
			k++
		}
		i += k
		j -= k
	}
	return rad
}
