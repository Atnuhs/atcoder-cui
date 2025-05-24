package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

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

// Ordered はconstraints.OrderedがAtCoderで使えないので、代わりに使う
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
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

// Is は整数をn個読み込む
func Is(n int) []int { return ILF2(n, I) }

// Iss は整数をh行w列の配列として読み込む
func Iss(h, w int) [][]int { return IGF2(h, w, I) }

// Sss は文字列をn個読み込む
func Ss(n int) []string { return ILF2(n, S) }

// Rss は文字列をn個読み込む
func Rs(n int) [][]rune { return ILF2(n, R) }

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

func Reduce[S any, T any](src []S, f func(T, S) T, acc T) T {
	ret := acc
	for _, v := range src {
		ret = f(ret, v)
	}
	return ret
}

func Uniq[T Ordered](vals []T) []T {
	return Reduce(vals, func(acc []T, val T) []T {
		if len(acc) == 0 || acc[len(acc)-1] != val {
			acc = append(acc, val)
		}
		return acc
	}, make([]T, 0, len(vals)))
}

type Compress[T Ordered] struct {
	toOrig []T       // idx -> value
	toIdx  map[T]int // value -> idx
}

func NewCompress[T Ordered](vals []T) *Compress[T] {
	v := ILF(len(vals), func(i int) T { return vals[i] })
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
