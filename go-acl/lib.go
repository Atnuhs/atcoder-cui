package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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

func InRange(x, l, r int) bool {
	return l <= x && x < r
}

// L1 は長さnの配列を、関数fで初期化する
func L1[T any](n int) []T {
	return make([]T, n)
}

// L2 はh行w列の配列を、関数fで初期化する
func L2[T any](n1, n2 int) [][]T {
	ret := make([][]T, n1)
	for i := range ret {
		ret[i] = make([]T, n2)
	}
	return ret
}

func L3[T any](n1, n2, n3 int) [][][]T {
	ret := make([][][]T, n1)
	for i := range ret {
		ret[i] = make([][]T, n2)
		for j := range ret[i] {
			ret[i][j] = make([]T, n3)
		}
	}
	return ret
}

func F1[T any](vals []T, fill T) {
	for i := range vals {
		vals[i] = fill
	}
}

func F2[T any](vals [][]T, fill T) {
	for i := range vals {
		for j := range vals[i] {
			vals[i][j] = fill
		}
	}
}

func F3[T any](vals [][][]T, fill T) {
	for i := range vals {
		for j := range vals[i] {
			for k := range vals[i][j] {
				vals[i][j][k] = fill
			}
		}
	}
}

// Jag は長さNのJagged配列を生成する
func Jag[T any](n int) [][]T {
	return L2[T](n, 0)
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

func F() float64 {
	var ret float64
	fmt.Fscan(In, &ret)
	return ret
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
func Is(n int) []int {
	ret := make([]int, n)
	for i := range ret {
		fmt.Fscan(In, &ret[i])
	}
	return ret
}

// Iss は整数をh行w列の配列として読み込む
func Iss(h, w int) [][]int {
	ret := L2[int](h, w)
	for i := range ret {
		for j := range ret[i] {
			fmt.Fscan(In, &ret[i][j])
		}
	}
	return ret
}

// Sss は文字列をn個読み込む
func Ss(n int) []string {
	ret := L1[string](n)
	for i := range ret {
		fmt.Fscan(In, &ret[i])
	}
	return ret
}

// Rss は文字列をn個読み込む
func Rs(n int) [][]rune {
	ret := make([][]rune, n)
	for i := range ret {
		ret[i] = R()
	}
	return ret
}

func Fs(n int) []float64 {
	ret := make([]float64, n)
	for i := range ret {
		fmt.Fscan(In, &ret[i])
	}
	return ret
}

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

func Reverse[T any](vals []T) []T {
	ret := make([]T, len(vals))
	for i := len(vals) - 1; i >= 0; i-- {
		ret[i] = vals[len(vals)-1-i]
	}
	return ret
}

func ReverseS(s string) string {
	return string(Reverse([]byte(s)))
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

func Bisect(ok, ng int, pred func(mid int) bool) int {
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

// Geは昇順ソート済みの配列aに対して、x以上の要素の左端Indexを返す
// aのすべての要素がxより小さい場合、len(a)を返す
func Ge[T Ordered](a []T, x T) int {
	ok, ng := len(a), -1
	for ok-ng > 1 {
		m := (ok + ng) >> 1
		if x <= a[m] {
			ok = m
		} else {
			ng = m
		}
	}
	return ok
}

// Gtは昇順ソート済みの配列aに対して、xより大きい要素の左端Indexを返す
// aのすべての要素がx以下の場合、len(a)を返す
func Gt[T Ordered](a []T, x T) int {
	ok, ng := len(a), -1
	for ok-ng > 1 {
		m := (ok + ng) >> 1
		if x < a[m] {
			ok = m
		} else {
			ng = m
		}
	}
	return ok
}

// Leは昇順ソート済みの配列aに対し、x以下の要素の右端Indexを返す
// aのすべての要素がxより大きい場合、-1を返す
func Le[T Ordered](a []T, x T) int {
	ok, ng := -1, len(a)
	for ng-ok > 1 {
		m := (ok + ng) >> 1
		if x >= a[m] {
			ok = m
		} else {
			ng = m
		}
	}
	return ok
}

// Ltは昇順ソート済みの配列aに対して、xより小さい要素の右端を返す
// aのすべての要素がx以上の場合、-1を返す
func Lt[T Ordered](a []T, x T) int {
	ok, ng := -1, len(a)
	for ng-ok > 1 {
		m := (ok + ng) >> 1
		if x > a[m] {
			ok = m
		} else {
			ng = m
		}
	}
	return ok
}

func Uniq[T Ordered](vals []T) []T {
	ret := make([]T, 0, len(vals))
	if len(vals) == 0 {
		return ret
	}
	ret = append(ret, vals[0])
	for _, v := range vals[1:] {
		if ret[len(ret)-1] != v {
			ret = append(ret, v)
		}
	}
	return ret
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

type Key string

func KeyInts(a []int) Key {
	if len(a) == 0 {
		return ""
	}
	var b strings.Builder

	b.Grow(len(a) * 3)
	b.WriteString(strconv.Itoa(a[0]))
	for i := 1; i < len(a); i++ {
		b.WriteByte(' ')
		b.WriteString(strconv.Itoa(a[i]))
	}
	return Key(b.String())
}

func (k Key) ToInts() []int {
	toks := strings.Fields(string(k))
	ret := make([]int, 0, len(toks))
	for _, s := range toks {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Errorf("failed to parse int %s: element of %v %w", s, k, err))
		}
		ret = append(ret, x)
	}
	return ret
}

func Sort[T Ordered](arr []T) []T {
	ret := append([]T(nil), arr...)
	sort.Slice(ret, func(i int, j int) bool { return ret[i] < ret[j] })
	return ret
}

func SortIdx[T Ordered](arr []T) []int {
	ret := make([]int, len(arr))
	for i := range ret {
		ret[i] = i
	}
	sort.SliceStable(ret, func(i, j int) bool { return arr[ret[i]] < arr[ret[j]] })
	return ret
}

func SortF[T any](arr []T, less LessFunc[T]) []T {
	ret := append([]T(nil), arr...)
	sort.Slice(ret, func(i, j int) bool { return less(ret[i], ret[j]) })
	return ret
}

func SortIdxF[T any](arr []T, less LessFunc[T]) []int {
	ret := make([]int, len(arr))
	for i := range ret {
		ret[i] = i
	}
	sort.Slice(ret, func(i, j int) bool { return less(arr[ret[i]], arr[ret[j]]) })
	return ret
}

func Pow(x, e int) int {
	return int(math.Pow(float64(x), float64(e)))
}
