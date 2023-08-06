package main


type SegmentTree struct {
    data []int
    n int
    mo *Monoid
}

func NewSegmentTree(arr []int, mo *Monoid) *SegmentTree {
    n := 1
    for n < len(arr) {
        n *= 2
    }

    data := Newis(2*n-1, func(i int) int {return mo.E})
    for i := range arr {
        j := i+n-1
        data[j] = arr[i]
    }

    for i := n-2; i>=0; i-- {
        c1 := 2*i+1
        c2 := 2*i+2
        data[i] = mo.Op(data[c1], data[c2])
    }

    return &SegmentTree{
        data: data,
        n: n,
        mo: mo,
    }
}

func (st *SegmentTree) Update(i, x int) {
    i += st.n -1
    st.data[i] = x
    for i > 0 {
        i = (i-1)/2  
        st.data[i] = st.mo.Op(st.data[2*i+1], st.data[2*i+2])
    }
}

func (st *SegmentTree) At(i int) int {
    return st.Query(i, i+1)
}

func (st *SegmentTree) Query(a, b int) int {
    return st.querySub(a, b, 0, 0, st.n) 
}

func (st *SegmentTree) querySub(a,b,n,l,r int) int {
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
