package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/Atnuhs/atcoder-cui/go-acl/testlib"
)

func Test_lIdx(t *testing.T) {
	tests := map[string]struct {
		idx  int
		want int
	}{
		"odd":  {idx: 5, want: 4},
		"even": {idx: 6, want: 6},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := lIdx(tc.idx)
			if tc.want != got {
				t.Errorf("want %d but got %d", tc.want, got)
			}
		})
	}
}

func Fuzz_lIdx(f *testing.F) {
	f.Add(0)
	f.Add(INF)

	f.Fuzz(func(t *testing.T, idx int) {
		if idx < 0 {
			return
		}

		want := (idx / 2) * 2
		got := lIdx(idx)
		if want != got {
			t.Errorf("idx: %d, want: %d, got: %d", idx, want, got)
		}
	})
}

func Test_rIdx(t *testing.T) {
	tests := map[string]struct {
		idx  int
		want int
	}{
		"odd":  {idx: 5, want: 5},
		"even": {idx: 6, want: 7},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := rIdx(tc.idx)
			if tc.want != got {
				t.Errorf("want %d but got %d", tc.want, got)
			}
		})
	}
}

func Fuzz_rIdx(f *testing.F) {
	f.Add(0)
	f.Add(INF)

	f.Fuzz(func(t *testing.T, idx int) {
		if idx < 0 {
			return
		}

		want := (idx/2)*2 + 1
		got := rIdx(idx)
		if want != got {
			t.Errorf("idx: %d, want: %d, got: %d", idx, want, got)
		}
	})
}

func Test_pIdx(t *testing.T) {
	tests := map[string]struct {
		idx  int
		want int
	}{
		"odd":  {idx: 5, want: 0},
		"even": {idx: 6, want: 2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := pIdx(tc.idx)
			if tc.want != got {
				t.Errorf("want %d but got %d", tc.want, got)
			}
		})
	}
}

func Fuzz_pIdx(f *testing.F) {
	f.Add(2)
	f.Add(3)
	f.Add(INF)

	f.Fuzz(func(t *testing.T, idx int) {
		if idx < 2 {
			return
		}

		want := ((idx - 2) / 4) * 2
		got := pIdx(idx)
		if want != got {
			t.Errorf("idx: %d, want: %d, got: %d", idx, want, got)
		}
	})
}

func Test_cIdx(t *testing.T) {
	tests := map[string]struct {
		idx  int
		want int
	}{
		"root0": {idx: 0, want: 2},
		"root1": {idx: 1, want: 3},
		"c1":    {idx: 2, want: 6},
		"c2":    {idx: 3, want: 7},
		"c3":    {idx: 4, want: 10},
		"c4":    {idx: 5, want: 11},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cIdx(tc.idx)
			if tc.want != got {
				t.Errorf("want %d but got %d", tc.want, got)
			}
		})
	}
}

func TestDEPQ_up(t *testing.T) {
	tests := map[string]struct {
		values []int
		idx    int
		want   []int
	}{
		"target r-r, less than pr": {values: []int{10, 5, 7, 6, 6, 4}, idx: 5, want: []int{10, 4, 7, 6, 6, 5}},
		"target r-r, equal to pr":  {values: []int{10, 5, 7, 6, 6, 5}, idx: 5, want: []int{10, 5, 7, 6, 6, 5}},
		"target r-r, more than pr": {values: []int{10, 5, 7, 6, 6, 6}, idx: 5, want: []int{10, 5, 7, 6, 6, 6}},

		"target r-r, less than pl": {values: []int{10, 5, 7, 6, 6, 9}, idx: 5, want: []int{10, 5, 7, 6, 9, 6}},
		"target r-r, equal to pl":  {values: []int{10, 5, 7, 6, 6, 10}, idx: 5, want: []int{10, 5, 7, 6, 10, 6}},
		"target r-r, more than pl": {values: []int{10, 5, 7, 6, 6, 11}, idx: 5, want: []int{11, 5, 7, 6, 10, 6}},

		"target r-l, less than pr": {values: []int{10, 5, 7, 6, 4, 6}, idx: 4, want: []int{10, 4, 7, 6, 6, 5}},
		"target r-l, equal to pr":  {values: []int{10, 5, 7, 6, 5, 6}, idx: 4, want: []int{10, 5, 7, 6, 6, 5}},
		"target r-l, more than pr": {values: []int{10, 5, 7, 6, 6, 6}, idx: 4, want: []int{10, 5, 7, 6, 6, 6}},

		"target r-l, less than pl": {values: []int{10, 5, 7, 6, 9, 6}, idx: 4, want: []int{10, 5, 7, 6, 9, 6}},
		"target r-l, equal to pl":  {values: []int{10, 5, 7, 6, 10, 6}, idx: 4, want: []int{10, 5, 7, 6, 10, 6}},
		"target r-l, more than pl": {values: []int{10, 5, 7, 6, 11, 6}, idx: 4, want: []int{11, 5, 7, 6, 10, 6}},

		"target l-r, less than pr": {values: []int{10, 5, 7, 4, 6, 5}, idx: 3, want: []int{10, 4, 7, 5, 6, 5}},
		"target l-r, equal to pr":  {values: []int{10, 5, 7, 5, 6, 5}, idx: 3, want: []int{10, 5, 7, 5, 6, 5}},
		"target l-r, more than pr": {values: []int{10, 5, 7, 6, 6, 5}, idx: 3, want: []int{10, 5, 7, 6, 6, 5}},

		"target l-r, less than pl": {values: []int{10, 5, 7, 9, 6, 5}, idx: 3, want: []int{10, 5, 9, 7, 6, 5}},
		"target l-r, equal to pl":  {values: []int{10, 5, 7, 10, 6, 5}, idx: 3, want: []int{10, 5, 10, 7, 6, 5}},
		"target l-r, more than pl": {values: []int{10, 5, 7, 11, 6, 5}, idx: 3, want: []int{11, 5, 10, 7, 6, 5}},

		"target l-l, less than pr": {values: []int{10, 5, 4, 6, 6, 5}, idx: 2, want: []int{10, 4, 6, 5, 6, 5}},
		"target l-l, equal to pr":  {values: []int{10, 5, 5, 6, 6, 5}, idx: 2, want: []int{10, 5, 6, 5, 6, 5}},
		"target l-l, more than pr": {values: []int{10, 5, 6, 6, 6, 5}, idx: 2, want: []int{10, 5, 6, 6, 6, 5}},

		"target l-l, less than pl": {values: []int{10, 5, 9, 6, 6, 5}, idx: 2, want: []int{10, 5, 9, 6, 6, 5}},
		"target l-l, equal to pl":  {values: []int{10, 5, 10, 6, 6, 5}, idx: 2, want: []int{10, 5, 10, 6, 6, 5}},
		"target l-l, more than pl": {values: []int{10, 5, 11, 6, 6, 5}, idx: 2, want: []int{11, 5, 10, 6, 6, 5}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			pq := &DEPQ[int]{values: tc.values}
			pq.up(tc.idx)
			got := pq.values
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("want(-), got(+)\n%s", diff)
			}
		})
	}
}

func TestDEPQ_down(t *testing.T) {
	tests := map[string]struct {
		values []int
		idx    int
		want   []int
	}{
		"min heap, p > cl, p < cr, cl < cr": {values: []int{10, 5, 7, 1, 6, 5}, idx: 1, want: []int{10, 1, 7, 5, 6, 5}},
		"min heap, p < cl, p > cr, cl > cr": {values: []int{10, 5, 7, 6, 6, 1}, idx: 1, want: []int{10, 1, 7, 6, 6, 5}},
		"min heap, p < cl, p < cr, cl < cr": {values: []int{10, 5, 7, 6, 6, 7}, idx: 1, want: []int{10, 5, 7, 6, 6, 7}},
		"min heap, p < cl, p < cr, cl > cr": {values: []int{10, 5, 7, 7, 6, 6}, idx: 1, want: []int{10, 5, 7, 7, 6, 6}},
		"min heap, p > cl, p > cr, cl < cr": {values: []int{10, 5, 7, 1, 6, 2}, idx: 1, want: []int{10, 1, 7, 5, 6, 2}},
		"min heap, p > cl, p > cr, cl > cr": {values: []int{10, 5, 7, 2, 6, 1}, idx: 1, want: []int{10, 1, 7, 2, 6, 5}},

		"max heap, p > cl, p < cr, cl < cr": {values: []int{10, 5, 7, 6, 20, 5}, idx: 0, want: []int{20, 5, 7, 6, 10, 5}},
		"max heap, p < cl, p > cr, cl > cr": {values: []int{10, 5, 20, 6, 6, 5}, idx: 0, want: []int{20, 5, 10, 6, 6, 5}},
		"max heap, p < cl, p < cr, cl < cr": {values: []int{10, 5, 20, 6, 30, 5}, idx: 0, want: []int{30, 5, 20, 6, 10, 5}},
		"max heap, p < cl, p < cr, cl > cr": {values: []int{10, 5, 30, 6, 20, 5}, idx: 0, want: []int{30, 5, 10, 6, 20, 5}},
		"max heap, p > cl, p > cr, cl < cr": {values: []int{10, 5, 7, 6, 9, 5}, idx: 0, want: []int{10, 5, 7, 6, 9, 5}},
		"max heap, p > cl, p > cr, cl > cr": {values: []int{10, 5, 7, 6, 6, 5}, idx: 0, want: []int{10, 5, 7, 6, 6, 5}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			pq := &DEPQ[int]{values: tc.values}
			pq.down(tc.idx)
			got := pq.values
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("want(-), got(+)\n%s", diff)
			}
		})
	}
}

func check_heap[T Ordered](t *testing.T, pq *DEPQ[T]) {
	t.Helper()
	for i := range pq.values {
		cl := cIdx(i)
		cr := cl + 2

		if cl >= pq.Size() {
			continue
		}

		if i&1 == 1 {
			// min heap
			if pq.values[cl] < pq.values[i] {
				t.Errorf("parent(ID,value)=(%d,%v), child(ID,value)=(%d, %v) at min heap", i, pq.values[i], cl, pq.values[cl])
			}
		} else {
			// max heap
			if pq.values[cl] > pq.values[i] {
				t.Errorf("parent(ID,value)=(%d,%v), child(ID,value)=(%d, %v) at max heap", i, pq.values[i], cl, pq.values[cl])
			}
		}

		if cr >= pq.Size() {
			continue
		}

		if i&1 == 1 {
			// min heap
			if pq.values[cr] < pq.values[i] {
				t.Errorf("parent(ID,value)=(%d,%v), child(ID,value)=(%d, %v) at min heap", i, pq.values[i], cr, pq.values[cr])
			}
		} else {
			// max heap
			if pq.values[cr] > pq.values[i] {
				t.Errorf("parent(ID,value)=(%d,%v), child(ID,value)=(%d, %v) at max heap", i, pq.values[i], cr, pq.values[cr])
			}
		}
	}
}

func TestDEPQ_Push(t *testing.T) {
	tests := map[string]struct {
		values []int
	}{
		"0~9": {values: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			pq := NewDEPQ[int]()
			for _, v := range tc.values {
				pq.Push(v)
			}
			t.Log(pq)
			check_heap(t, pq)
		})
	}
}

func TestDEPQ_Constructor(t *testing.T) {
	tests := map[string]struct {
		values []int
	}{
		"0~9": {values: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			pq := NewDEPQ[int](tc.values...)
			t.Log(pq)
			check_heap(t, pq)
		})
	}
}

func TestDEPQ_PopMax(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	want := make([]int, len(values))
	copy(want, values)
	for i, j := 0, len(want)-1; i < j; i, j = i+1, j-1 {
		want[i], want[j] = want[j], want[i]
	}
	pq := NewDEPQ(values...)
	got := make([]int, 0, len(want))
	for !pq.Empty() {
		got = append(got, pq.PopMax())
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("want(-), got(+)\n%s", diff)
	}
}

func TestDEPQ_PopMin(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	want := make([]int, len(values))
	copy(want, values)
	pq := NewDEPQ(values...)
	got := make([]int, 0, len(want))
	for !pq.Empty() {
		got = append(got, pq.PopMin())
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("want(-), got(+)\n%s", diff)
	}
}

func TestDEPQ_EmptyOperations(t *testing.T) {
	tests := map[string]struct {
		setup func() *DEPQ[int]
		want  bool
	}{
		"empty queue": {
			setup: func() *DEPQ[int] { return NewDEPQ[int]() },
			want:  true,
		},
		"non-empty queue": {
			setup: func() *DEPQ[int] { return NewDEPQ(1, 2, 3) },
			want:  false,
		},
		"queue after popping all": {
			setup: func() *DEPQ[int] {
				pq := NewDEPQ(1)
				pq.PopMin()
				return pq
			},
			want: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			pq := tc.setup()
			got := pq.Empty()
			testlib.AclAssert(t, tc.want, got)
		})
	}
}

func TestDEPQ_SizeOperations(t *testing.T) {
	tests := map[string]struct {
		setup func() *DEPQ[int]
		want  int
	}{
		"empty": {
			setup: func() *DEPQ[int] { return NewDEPQ[int]() },
			want:  0,
		},
		"single element": {
			setup: func() *DEPQ[int] { return NewDEPQ(42) },
			want:  1,
		},
		"multiple elements": {
			setup: func() *DEPQ[int] { return NewDEPQ(1, 2, 3, 4, 5) },
			want:  5,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			pq := tc.setup()
			got := pq.Size()
			testlib.AclAssert(t, tc.want, got)
		})
	}
}

func TestDEPQ_GetOperations(t *testing.T) {
	pq := NewDEPQ(5, 1, 9, 3, 7)
	
	testlib.AclAssert(t, 1, pq.GetMin())
	testlib.AclAssert(t, 9, pq.GetMax())
	
	// Verify getting doesn't modify the queue
	testlib.AclAssert(t, 5, pq.Size())
	testlib.AclAssert(t, 1, pq.GetMin())
	testlib.AclAssert(t, 9, pq.GetMax())
}

func TestDEPQ_EdgeCasesOperations(t *testing.T) {
	t.Run("single element operations", func(t *testing.T) {
		pq := NewDEPQ(42)
		testlib.AclAssert(t, 42, pq.GetMin())
		testlib.AclAssert(t, 42, pq.GetMax())
		
		min := pq.PopMin()
		testlib.AclAssert(t, 42, min)
		testlib.AclAssert(t, true, pq.Empty())
	})
	
	t.Run("two element operations", func(t *testing.T) {
		pq := NewDEPQ(3, 1)
		testlib.AclAssert(t, 1, pq.GetMin())
		testlib.AclAssert(t, 3, pq.GetMax())
		
		max := pq.PopMax()
		testlib.AclAssert(t, 3, max)
		testlib.AclAssert(t, 1, pq.Size())
		testlib.AclAssert(t, 1, pq.GetMin())
		testlib.AclAssert(t, 1, pq.GetMax())
	})
}
