package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type Pair struct {
	u, v int
}

func NewPair(u, v int) *Pair {
	return &Pair{u, v}
}
func (p *Pair) String() string {
	return fmt.Sprintf("%d %d", p.u, p.v)
}

func init() {
	in.Split(bufio.ScanWords)
	in.Buffer([]byte{}, math.MaxInt64)
}

func reads() string {
	in.Scan()
	return in.Text()
}

func readss(n int) []string {
	ret := make([]string, n)
	for i := range ret {
		ret[i] = reads()
	}
	return ret
}

func readrs(n int) [][]rune {
	ret := make([][]rune, n)
	for i := range ret {
		ret[i] = []rune(reads())
	}
	return ret
}

func readi() int {
	in.Scan()
	ret, _ := strconv.Atoi(in.Text())
	return ret
}

func readis(n int) []int {
	ret := make([]int, n)
	for i := range ret {
		ret[i] = readi()
	}
	return ret
}

func IsPrime(x int) bool {
	if x == 1 {
		return false
	}

	rx := sqrt(x)
	for i := 2; i <= rx; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func Factorize(x int) []*Pair {
	if x == 1 {
		return []*Pair{}
	}

	rx := sqrt(x)
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

func Divisors(x int) []int {
    ret := make([]int, 0)

    rx := sqrt(x)
    for i := 1; i<= rx; i++ {
        if x % i != 0 {
            continue
        }
        ret = append(ret, i)
        if i != x/i {
            ret = append(ret, x/i)
        }
    }
    return ret
}

// Eratosthenes sieve
type EratosthenesSieve struct {
	isPrime   []bool
	minFactor []int
}

func NewSieve(n int) *EratosthenesSieve {
	isPrime := make([]bool, n+1)
	minFactor := make([]int, n+1)

	for i := range isPrime {
		isPrime[i] = true
		minFactor[i] = -1
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

		for j := i * 2; j <= n; j += i {
			isPrime[j] = false

			if minFactor[j] == -1 {
				minFactor[j] = i
			}
		}
	}
	return &EratosthenesSieve{
		isPrime:   isPrime,
		minFactor: minFactor,
	}
}

func (sv *EratosthenesSieve) IsPrime(x int) bool {
	return sv.isPrime[x]
}

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

func modPow(x, e, mod int) int {
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

func inv(x, mod int) int {
	return modPow(x, mod-2, mod)
}

func popBack(a *[]int) int {
	ret := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return ret
}

func popFront(a *[]int) int {
	ret := (*a)[0]
	*a = (*a)[1:]
	return ret
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func sqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

func nextPerm(a []int) bool {
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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func NewUf(n int) []int {
	uf := make([]int, n)
	for i := range uf {
		uf[i] = -1
	}
	return uf
}

func root(uf []int, x int) int {
	if uf[x] < 0 {
		return x
	} else {
		uf[x] = root(uf, uf[x])
		return uf[x]
	}
}

func family(uf []int, x, y int) bool {
	return root(uf, x) == root(uf, y)
}

func size(uf []int, x int) int {
	return -uf[root(uf, x)]
}

func unite(uf []int, x, y int) {
	rx := root(uf, x)
	ry := root(uf, y)
	if rx == ry {
		return
	}
	if size(uf, rx) < size(uf, ry) {
		rx = rx ^ ry
		ry = rx ^ ry
		rx = rx ^ ry
	}
	uf[rx] += uf[ry]
	uf[ry] = rx
}

type splay_node struct {
	l, r, p *splay_node
	size    int
	value   int
}

func (n *splay_node) rotate() {
	p := n.p

	if pp := p.p; pp != nil {
		if pp.l == p {
			pp.l = n
		} else {
			pp.r = n
		}
		n.p = pp
	}

	var c *splay_node
	if p.l == n {
		c = n.r
		n.r = p
		p.l = c
	} else {
		c = n.l
		n.l = p
		p.r = c
	}

	if c != nil {
		c.p = p
	}
	p.p = n
	p.update()
	n.update()
}

func (n *splay_node) state() int {
	if n.p == nil {
		return 0
	}
	if n.p.l == n {
		return 1
	}
	if n.p.r == n {
		return -1
	}
	return 0
}

func (n *splay_node) splay() {
	for n.p != nil {
		// n has parent

		if n.p.state() == 0 {
			// n.p doesn't have p
			n.rotate()
		}

		if n.state() == n.p.state() {
			n.p.rotate()
			n.rotate()
		} else {
			n.rotate()
			n.rotate()
		}
	}
}

func (n *splay_node) update() {
	n.size = 1
	if n.l != nil {
		n.size += n.l.size
	}
	if n.r != nil {
		n.size += n.r.size
	}
}

func get(ind int, root *splay_node) *splay_node {
	now := root
	for {
		lsize := 0
		if now.l != nil {
			lsize = now.l.size
		}
		if ind < lsize {
			now = now.l
		}
		if ind == lsize {
			now.splay()
			return now
		}
		if ind > lsize {
			now = now.r
			ind -= lsize + 1
		}
	}
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
