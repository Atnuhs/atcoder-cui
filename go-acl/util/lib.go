package util

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

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

type Ordered interface {
	constraints.Ordered
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

func Readss(n int) []string {
	return NewArr(n, func(i int) string { return Reads() })
}

func Readrs(n int) [][]rune {
	return NewArr(n, func(i int) []rune { return Readr() })
}

func Readis(n int) []int {
	return NewArr(n, func(i int) int { return Readi() })
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
	ret := true
	for i, v := range vals {
		ret = ret && f(i, v)
	}
	return ret
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
