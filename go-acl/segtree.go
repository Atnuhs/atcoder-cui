package main

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
	n := NextPow2(len(arr))

	data := L1[T](2*n - 1)
	for i := range data {
		data[i] = mo.E
	}

	for i := range arr {
		j := i + n - 1
		data[j] = arr[i]
	}

	for i := n - 2; i >= 0; i-- {
		c1, c2 := (i<<1)+1, (i<<1)+2
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
		i = (i - 1) >> 1
		c1, c2 := (i<<1)+1, (i<<1)+2
		st.data[i] = st.mo.Op(st.data[c1], st.data[c2])
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

	c1, c2 := (n<<1)+1, (n<<1)+2
	mid := (l + r) >> 1
	vl := st.querySub(a, b, c1, l, mid)
	vr := st.querySub(a, b, c2, mid, r)
	return st.mo.Op(vl, vr)
}
