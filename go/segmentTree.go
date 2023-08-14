package main


type SegmentTree[T any] struct {
    data []T
    n int
    mo *Monoid[T]
}

func NewSegmentTree[T any](arr []T, mo *Monoid[T]) *SegmentTree[T] {
    n := 1
    for n < len(arr) {
        n *= 2
    }

    data := NewArr(2*n-1, func(i int) T {return mo.E})
    for i := range arr {
        j := i+n-1
        data[j] = arr[i]
    }

    for i := n-2; i>=0; i-- {
        c1 := 2*i+1
        c2 := 2*i+2
        data[i] = mo.Op(data[c1], data[c2])
    }

    return &SegmentTree[T]{
        data: data,
        n: n,
        mo: mo,
    }
}

func (st *SegmentTree[T]) Update(i int, x T) {
    i += st.n -1
    st.data[i] = x
    for i > 0 {
        i = (i-1)/2  
        st.data[i] = st.mo.Op(st.data[2*i+1], st.data[2*i+2])
    }
}

func (st *SegmentTree[T]) At(i int) T {
    return st.Query(i, i+1)
}

func (st *SegmentTree[T]) Query(a, b int) T {
    return st.querySub(a, b, 0, 0, st.n) 
}

func (st *SegmentTree[T]) querySub(a,b,n,l,r int) T {
    if r <= a || b <= l {
        return st.mo.E
    }

    if a <= l && r <= b {
        return st.data[n]
    }

    vl := st.querySub(a,b,2*n+1, l, (l+r)/2)
    vr := st.querySub(a, b, 2*n+2, (l+r)/2, r)
    return st.mo.Op(vl, vr)
}
