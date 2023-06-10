package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

const (
	MOD1 = 1000000007
	MOD2 = 998244353
	// INF is 10^18
	INF = 1000000000000000000
)

func init() {
	in.Split(bufio.ScanWords)
	in.Buffer([]byte{}, math.MaxInt64)
}

type Pair struct {
	u, v int
}

func NewPair(u, v int) *Pair {
	return &Pair{u, v}
}

// String return 'p.u p.v'
func (p *Pair) String() string {
	return fmt.Sprintf("%d %d", p.u, p.v)
}

func Reads() string {
	in.Scan()
	return in.Text()
}

func Readss(n int) []string {
	ret := make([]string, n)
	for i := range ret {
		ret[i] = Reads()
	}
	return ret
}

func Readrs(n int) [][]rune {
	ret := make([][]rune, n)
	for i := range ret {
		ret[i] = []rune(Reads())
	}
	return ret
}

func Readi() int {
	in.Scan()
	ret, _ := strconv.Atoi(in.Text())
	return ret
}

func Readis(n int) []int {
	ret := make([]int, n)
	for i := range ret {
		ret[i] = Readi()
	}
	return ret
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

// PopBack is O(1)
func PopBack(a *[]int) int {
	ret := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return ret
}

// PopFront is O(1)
func PopFront(a *[]int) int {
	ret := (*a)[0]
	*a = (*a)[1:]
	return ret
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
func Extrema(vals ...int) (int, int) {
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

func Max(vals ...int) int {
	_, ma := Extrema(vals...)
	return ma
}

func Min(vals ...int) int {
	mi, _ := Extrema(vals...)
	return mi
}

func Sum(vals ...int) int {
	sum := 0
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
func Factorize(x int) []*Pair {
	if x == 1 {
		return []*Pair{}
	}

	rx := Sqrt(x)
	n := x
	ret := make([]*Pair, 0)
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
func CountDivisors(pairs []*Pair) int {
	ans := 1
	for _, pe := range pairs {
		ans *= (pe.v + 1)
	}
	return ans
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
func (sv *EratosthenesSieve) Factorize(x int) []*Pair {
	ret := make([]*Pair, 0)
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

// Mobius is O(1) returns
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

func (sn *SplayNode) update() {
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
	return 0
}

func (sn *SplayNode) rotate() {
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
	for sn.state() == 0 {
		// sn is root
		return
	}

	if sn.p.state() == 0 {
		// sn.p is root
		sn.rotate()
		return
	}

	if sn.state() == sn.p.state() {
		sn.p.rotate()
		sn.rotate()
	} else {
		sn.rotate()
		sn.rotate()
	}
}

func (sn *SplayNode) describe(rank int) string {
	ret := ""
	if sn.r != nil {
		ret += sn.r.describe(rank + 1)
	}
	ret += fmt.Sprintf(
		strings.Repeat("    ", rank)+"-[k:%d, v:%d, sz: %d]\n",
		sn.key,
		sn.value,
		sn.size,
	)

	if sn.l != nil {
		ret += sn.l.describe(rank + 1)
	}
	return ret
}

func get_subSN(ind int, node *SplayNode) (int, *SplayNode) {
	if node == nil {
		return -1, nil
	}
	ls := 0
	if node.l != nil {
		ls = node.l.size
	}

	switch {
	case ind < ls:
		return ind, node.l
	case ind == ls:
		return -1, node
	case ind > ls:
		return ind - (ls + 1), node.r
	}
	return -1, nil
}

func GetSN(ind int, node *SplayNode) *SplayNode {
	for ind != -1 {
		ind, node = get_subSN(ind, node)
	}
	// node found
	if node != nil {
		node.splay()
	}
	return node
}

func MergeSN(lroot, rroot *SplayNode) *SplayNode {
	if lroot == nil {
		return rroot
	}
	if rroot == nil {
		return lroot
	}
	lroot = GetSN(lroot.size-1, lroot) // always found
	lroot.r = rroot
	rroot.p = lroot
	lroot.update()
	return lroot
}

func SplitSN(ind int, root *SplayNode) (*SplayNode, *SplayNode) {
	if root == nil {
		return nil, nil
	}
	if ind == root.size {
		return root, nil
	}

	rroot := GetSN(ind, root)
	if rroot == nil {
		// rroot not found
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

func InsertSN(ind int, root *SplayNode, node *SplayNode) *SplayNode {
	lroot, rroot := SplitSN(ind, root)
	return MergeSN(MergeSN(lroot, node), rroot)
}

func DeleteSN(ind int, root *SplayNode) (*SplayNode, *SplayNode) {
	lroot, rroot := SplitSN(ind, root)
	del, rroot := SplitSN(1, rroot)
	root = MergeSN(lroot, rroot)
	return root, del
}

type heapImpl []int

func (h heapImpl) Len() int           { return len(h) }
func (h heapImpl) Less(i, j int) bool { return h[i] < h[j] }
func (h heapImpl) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *heapImpl) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heapImpl) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type PriorityQueue struct {
	value heapImpl
}

func NewPriorityQueue() *PriorityQueue {
	value := &heapImpl{}
	heap.Init(value)
	return &PriorityQueue{}
}

func (pq *PriorityQueue) Push(x int) {
	heap.Push(&pq.value, x)
}

func (pq *PriorityQueue) Pop() int {
	x := heap.Pop(&pq.value)
	return x.(int)
}

func NewGraph(n int) [][]int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, 0)
	}
	return g
}

func main() {
	defer out.Flush()
}
