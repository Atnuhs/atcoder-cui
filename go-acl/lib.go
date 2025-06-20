package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	MOD1 = 1000000007
	MOD2 = 998244353
	// INF is 10^18
	INF = 1000000000000000000

	// Buffer size constants
	BufferSize = 1 << 20

	// Error messages for data structures
	ErrEmptyContainer = "operation on empty container"
	ErrOutOfIndex     = "index out of range"
)

var (
	In  = bufio.NewReaderSize(os.Stdin, BufferSize)
	Out = bufio.NewWriterSize(os.Stdout, BufferSize)
	Dbg = bufio.NewWriterSize(os.Stderr, BufferSize)
)

// Ordered はconstraints.OrderedがAtCoderで使えないので、代わりに使う
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Rep(n int) []struct{} {
	return make([]struct{}, n)
}

// MakeSlice は長さnの配列を、関数fで初期化する
func MakeSlice[T any](n int, f func(i int) T) []T {
	ret := make([]T, n)
	for i := range ret {
		ret[i] = f(i)
	}
	return ret
}

// MakeGrid はh行w列の配列を、関数fで初期化する
func MakeGrid[T any](h, w int, f func(ih, iw int) T) [][]T {
	return MakeSlice(h, func(ih int) []T {
		return MakeSlice(w, func(iw int) T {
			return f(ih, iw)
		})
	})
}

// MakeSliceWith は長さnの配列を、関数fで初期化する
func MakeSliceWith[T any](n int, f func() T) []T { return MakeSlice(n, func(_ int) T { return f() }) }

// MakeGridWith はh行w列の配列を、関数fで初期化する
func MakeGridWith[T any](h, w int, f func() T) [][]T {
	return MakeGrid(h, w, func(_ int, _ int) T { return f() })
}

// MakeSliceOf は長さnの配列を、値vで初期化する
// vに配列を指定すると、すべてポインタが同じになる
func MakeSliceOf[T any](n int, v T) []T {
	return MakeSlice(n, func(_ int) T { return v })
}

// MakeGridOf はh行w列の配列を、値vで初期化する
// vに配列を指定すると、すべてポインタが同じになる
func MakeGridOf[T any](h, w int, v T) [][]T {
	return MakeGrid(h, w, func(_ int, _ int) T { return v })
}

// MakeJaggedSlice はn行のグラフをjagged配列で初期化する
func MakeJaggedSlice[T any](n int) [][]T {
	return MakeSlice(n, func(_ int) []T { return make([]T, 0) })
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

// II は整数を2つ読み込む
func II() (int, int) {
	var ret1, ret2 int
	fmt.Fscan(In, &ret1, &ret2)
	return ret1, ret2
}

// III は整数を3つ読み込む
func III() (int, int, int) {
	var ret1, ret2, ret3 int
	fmt.Fscan(In, &ret1, &ret2, &ret3)
	return ret1, ret2, ret3
}

// IIII は整数を4つ読み込む
func IIII() (int, int, int, int) {
	var ret1, ret2, ret3, ret4 int
	fmt.Fscan(In, &ret1, &ret2, &ret3, &ret4)
	return ret1, ret2, ret3, ret4
}

// Is は整数をn個読み込む
func Is(n int) []int { return MakeSliceWith(n, I) }

// Iss は整数をh行w列の配列として読み込む
func Iss(h, w int) [][]int { return MakeGridWith(h, w, I) }

// Sss は文字列をn個読み込む
func Ss(n int) []string { return MakeSliceWith(n, S) }

// Rss は文字列をn個読み込む
func Rs(n int) [][]rune { return MakeSliceWith(n, R) }

// formatSlice はスライスを文字列に変換する
func formatSlice[T any](slice []T, formatter func(T) string) {
	for i, x := range slice {
		if i > 0 {
			fmt.Fprint(Out, " ")
		}
		fmt.Fprint(Out, formatter(x))
	}
}

// Ans は出力を行う
func Ans(args ...any) {
	for i, arg := range args {
		switch v := arg.(type) {
		case float64:
			fmt.Fprintf(Out, "%.14f", v)
		case []int:
			formatSlice(v, func(x int) string { return fmt.Sprintf("%d", x) })
		case []string:
			formatSlice(v, func(x string) string { return x })
		case []float64:
			formatSlice(v, func(x float64) string { return fmt.Sprintf("%.14f", x) })
		default:
			fmt.Fprint(Out, v)
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
func YesNo(b bool) {
	if b {
		Yes()
	} else {
		No()
	}
}

// YesNoFunc は関数の結果に応じてYesまたはNoを出力する
func YesNoFunc(f func() bool) {
	YesNo(f())
}

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

func UnReduce[A, B any](a A, f func(A) (A, *B)) []B {
	nextA, elem := f(a)
	if elem == nil {
		return nil
	}
	return append([]B{*elem}, UnReduce(nextA, f)...)
}

func Reduce[A any, B any](src []A, f func(val A, acc B) B, acc B) B {
	ret := acc
	for _, v := range src {
		ret = f(v, ret)
	}
	return ret
}

func Scan[A, B any](vals []A, init B, f func(B, A) B) []B {
	acc := make([]B, 0, len(vals)+1)
	acc = append(acc, init)
	return Reduce(vals, func(val A, acc []B) []B {
		next := f(acc[len(acc)-1], val)
		return append(acc, next)
	}, acc)
}

func Map[A any, B any](vals []A, f func(val A) B) []B {
	return Reduce(vals, func(val A, acc []B) []B {
		return append(acc, f(val))
	}, make([]B, 0, len(vals)))
}

// MapAccumLは状態Sを持った状態で、[]A -> []Bの写像作成し状態Sと[]Bを返す
func MapAccumL[S, A, B any](vals []A, acc S, f func(S, A) (S, B)) (S, []B) {
	type pair Pair[S, []B]
	ret := Reduce(vals, func(val A, acc pair) pair {
		newS, y := f(acc.U, val)
		return pair{newS, append(acc.V, y)}
	}, pair{acc, nil})
	return ret.U, ret.V
}

// IterateNはx := f(x)を0回~N回繰り返した結果を配列で返す関数
func IterateN[T any](x0 T, n int, f func(T) T) []T {
	_, ys := MapAccumL(Rep(n), x0, func(state T, _ struct{}) (T, T) {
		cur := state
		state = f(state)
		return state, cur
	})
	return ys
}

// Rangeはnを引数で、[0, 1, 2, ..., n-1]の配列を返す
func Range(n int) []int {
	return IterateN(0, n, func(x int) int { return x + 1 })
}

// ZipWithは[]Aと[]Bから、(A, B) => Cの関数で[]Cを作る
func ZipWith[A, B, C any](as []A, bs []B, f func(A, B) C) []C {
	n := Min(len(as), len(bs))
	_, ys := MapAccumL(Rep(n), 0, func(i int, _ struct{}) (int, C) {
		return i + 1, f(as[i], bs[i])
	})
	return ys
}

// Pairwiseはvals []Aの隣り合った要素(v_i, v_i+1)からRを生成する関数で、[]Bを返す
func Pairwise[A, B any](vals []A, f func(A, A) B) []B {
	if len(vals) < 2 {
		return nil
	}
	return ZipWith(vals[:len(vals)-1], vals[1:], f)
}

// Windowは[]Tの各要素で、幅kのウィンドウを生成する。
// vals = [1,2,3,4,5], k=3 => [1,2,3], [2,3,4], [3,4,5]
func Window[T any](vals []T, k int) [][]T {
	if k <= 0 || len(vals) < k {
		return nil
	}
	m := len(vals) - k + 1
	init := make([][]T, 0, m)
	return Reduce(Range(m), func(i int, acc [][]T) [][]T {
		return append(acc, vals[i:i+k])
	}, init)
}

func Uniq[T Ordered](vals []T) []T {
	return Reduce(vals, func(val T, acc []T) []T {
		if len(acc) == 0 || acc[len(acc)-1] != val {
			acc = append(acc, val)
		}
		return acc
	}, make([]T, 0, len(vals)))
}

func Filter[T any](vals []T, f func(val T) bool) []T {
	return Reduce(vals, func(val T, acc []T) []T {
		if f(val) {
			return append(acc, val)
		}
		return acc
	}, make([]T, 0, len(vals)))
}

func Count[T any](vals []T, f func(val T) bool) int {
	return Reduce(vals, func(val T, acc int) int {
		if f(val) {
			return acc + 1
		}
		return acc
	}, 0)
}

func RotateCW90[T any](src [][]T) [][]T {
	if len(src) == 0 || len(src[0]) == 0 {
		return nil
	}
	h, w := len(src), len(src[0])

	ret := make([][]T, w)
	for i := range ret {
		ret[i] = make([]T, h)
	}
	for ih := range src {
		for iw := range src[ih] {
			jh, jw := iw, h-1-ih
			ret[jh][jw] = src[ih][iw]
		}
	}
	return ret
}

func Bisect(ok, ng int, pred func(int) bool) int {
	for Abs(ng-ok) > 1 {
		mid := (ok + ng) >> 1
		if pred(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

type Compress[T Ordered] struct {
	toOrig []T       // idx -> value
	toIdx  map[T]int // value -> idx
}

func NewCompress[T Ordered](vals []T) *Compress[T] {
	// Copy and sort the values
	v := make([]T, len(vals))
	copy(v, vals)
	sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
	v = Uniq(v)

	m := make(map[T]int, len(v))
	for i, x := range v {
		m[x] = i
	}
	return &Compress[T]{
		toOrig: v,
		toIdx:  m,
	}
}

func (c *Compress[T]) Idx(x T) int {
	if i, ok := c.toIdx[x]; ok {
		return i
	}
	return -1
}

func (c *Compress[T]) Val(i int) T {
	return c.toOrig[i]
}

func (c *Compress[T]) Size() int {
	return len(c.toOrig)
}
