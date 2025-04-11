package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// lib
var (
	In  = bufio.NewScanner(os.Stdin)
	Out = bufio.NewWriter(os.Stdout)
	Dbg = bufio.NewWriter(os.Stderr)
)

const (
	MOD1 = 1000000007
	MOD2 = 998244353
	// INF is 10^18
	INF = 1000000000000000000
)

func init() {
	In.Split(bufio.ScanWords)
	In.Buffer([]byte{}, math.MaxInt64)
}

func main() {
	defer Out.Flush()
}

type Ordered interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type Pair[T any] struct {
	u, v T
}

func NewPair[T any](u, v T) *Pair[T] {
	return &Pair[T]{u, v}
}

// String return 'p.u p.v'
func (p *Pair[T]) String() string {
	return fmt.Sprintf("%v %v", p.u, p.v)
}

func NewArr[T any](n int, f func(i int) T) []T {
	ret := make([]T, n)
	for i := range ret {
		ret[i] = f(i)
	}
	return ret
}

func Reads() string {
	In.Scan()
	return In.Text()
}

func Readr() []rune {
	return []rune(Reads())
}

func Readi() int {
	In.Scan()
	ret, _ := strconv.Atoi(In.Text())
	return ret
}

func ReadLine[T any](n int, f func() T) []T {
	return NewArr(n, func(_ int) T { return f() })
}

func ReadGrid[T any](h, w int, f func() T) [][]T {
	return NewArr[[]T](h, func(_ int) []T {
		return ReadLine[T](w, f)
	})
}

// PopBack is O(1)
func PopBack[T any](a *[]T) T {
	ret := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return ret
}

// PopFront is O(1)
func PopFront[T any](a *[]T) T {
	ret := (*a)[0]
	*a = (*a)[1:]
	return ret
}

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

type PriorityQueue[T Ordered] struct {
	value heapImpl[T]
}

func NewPriorityQueue[T Ordered]() *PriorityQueue[T] {
	value := &heapImpl[T]{}
	heap.Init(value)
	return &PriorityQueue[T]{}
}

func (pq *PriorityQueue[T]) Push(x T) {
	heap.Push(&pq.value, x)
}

func (pq *PriorityQueue[T]) Pop() T {
	x := heap.Pop(&pq.value)
	return x.(T)
}

func NewGraph(n int) [][]int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, 0)
	}
	return g
}

func All[T any](vals []T, f func(i int, v T) bool) bool {
	for i, v := range vals {
		if f(i, v) {
			return false
		}
	}
	return true
}

func Any[T any](vals []T, f func(i int, v T) bool) bool {
	for i, v := range vals {
		if f(i, v) {
			return true
		}
	}
	return false
}

func Ans(args ...interface{}) {
	for i, arg := range args {
		if reflect.TypeOf(arg).Kind() == reflect.Slice {
			fmt.Fprint(Out, strings.Trim(fmt.Sprint(arg), "[]"))
		} else {
			fmt.Fprint(Out, arg)
		}
		if i < len(args)-1 {
			fmt.Fprint(Out, " ")
		}
	}
	fmt.Fprintln(Out)
}

func Yes() {
	Ans("Yes")
}

func No() {
	Ans("No")
}

func YesNo(f func() bool) {
	if f() {
		Yes()
	} else {
		No()
	}
}

func lIdx(idx int) int {
	return idx & ^1
}

func rIdx(idx int) int {
	return idx | 1
}

func pIdx(idx int) int {
	return ((idx >> 1) - 1) & ^1
}

func cIdx(idx int) int {
	return (idx & ^1)<<1 | 2 | idx&1
}

const depqCap = 1000

type DEPQ[T Ordered] struct {
	values []T
}

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

// Monoid
type (
	Operator[T any] func(x1, x2 T) T
	Monoid[T any]   struct {
		Op Operator[T]
		E  T
	}
)

func MoMax() *Monoid[int] {
	return &Monoid[int]{
		Op: func(x1, x2 int) int {
			return Max(x1, x2)
		},
		E: -INF,
	}
}

func MoMin() *Monoid[int] {
	return &Monoid[int]{
		Op: func(x1, x2 int) int {
			return Min(x1, x2)
		},
		E: INF,
	}
}

func MoSum[T int | float64]() *Monoid[T] {
	return &Monoid[T]{
		Op: func(x1, x2 T) T {
			return x1 + x2
		},
		E: 0,
	}
}

func MoXOR() *Monoid[int] {
	return &Monoid[int]{
		Op: func(x1, x2 int) int {
			return x1 ^ x2
		},
		E: 0,
	}
}

func MoMODMul(mod int) *Monoid[int] {
	return &Monoid[int]{
		Op: func(x1, x2 int) int {
			return (x1 * x2) % mod
		},
		E: 1,
	}
}


type SplayNode struct {
	l, r, p         *SplayNode
	size            int
	key             int
	value, min, max int
}

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


type SplaySet struct {
	root *SplayNode
}

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

// NewUnionFind
type UnionFind struct {
	data []int
}

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
		rx = rx ^ ry
		ry = rx ^ ry
		rx = rx ^ ry
	}

	uf.data[rx] += uf.data[ry]
	uf.data[ry] = rx
}

type SplayMap struct {
	root *SplayNode
}

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

type SegmentTree[T any] struct {
	data []T
	n    int
	mo   *Monoid[T]
}

func NewSegmentTree[T any](arr []T, mo *Monoid[T]) *SegmentTree[T] {
	n := 1
	for n < len(arr) {
		n *= 2
	}

	data := NewArr(2*n-1, func(i int) T { return mo.E })
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

type WEdge struct {
	from, to, weight int
}

func Kruskal(n int, edges []WEdge) (int, []WEdge) {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	uf := NewUnionFind(n)
	ret := make([]WEdge, 0)
	sum := 0
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

// Manacher algorithm
func Manacher(s string) []int {
	m := len(s)
	rad := make([]int, m)

	fmt.Println(s)
	i,j := 0, 0
	for i < m {
		fmt.Println(i,j,rad)
		for i-j >= 0 && i+j < m && s[i-j] == s[i+j] {
			j++
		}
		fmt.Println(j)
		rad[i] = j
		k := 1
		for i-k >= 0 && k + rad[i-k] < j {
			rad[i+k] = rad[i-k]
			k++	
		}
		i += k
		j -= k
	}
	return rad
}