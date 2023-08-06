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
