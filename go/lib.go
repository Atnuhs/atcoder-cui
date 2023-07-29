package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
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

func Newss(n int, f func(i int) string) []string {
    ret := make([]string, n)
    for i := range ret {
        ret[i] = f(i)
    }
    return ret
}

func Newrs(n int, f func(i int) []rune) [][]rune {
    ret := make([][]rune, n)
    for i := range ret {
        ret[i] = f(i)
    }
    return ret
}

func Newis(n int, f func(i int) int) []int {
    ret := make([]int, n)
    for i := range ret {
        ret[i] = f(i)
    }
    return ret
}

func Reads() string {
	in.Scan()
	return in.Text()
}

func Readr() []rune {
    return []rune(Reads())
}

func Readi() int {
	in.Scan()
	ret, _ := strconv.Atoi(in.Text())
	return ret
}

func Readss(n int) []string {
    return Newss(n, func(i int) string {return Reads()})
}


func Readrs(n int) [][]rune {
    return Newrs(n, func(i int) []rune {return Readr()})
}


func Readis(n int) []int {
    return Newis(n, func(i int) int {return Readi()})
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

func All(vals []int, f func(i, v int) bool) bool {
    ret := true 
    for i, v := range vals {
        ret = ret && f(i, v)
    }
    return ret
}

func Any(vals []int, f func(i, v int) bool) bool {
    for i, v := range vals {
        if f(i, v) {
            return true
        }
    }
    return false
}

func Ans(a ...interface{}) {
    fmt.Fprintln(out, a...)
}

func Yes() {
    Ans("Yes")
}

func No() {
    Ans("No")
}

func YesNo(f func()bool) {
    if f() {
        Yes()
    } else {
        No()
    }
}
