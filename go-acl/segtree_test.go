package main

import (
	"testing"

	"github.com/Atnuhs/atcoder-cui/go-acl/testlib"
)

func TestMoMax(t *testing.T) {
	mo := MoMax()
	
	testlib.AclAssert(t, -INF, mo.E)
	testlib.AclAssert(t, 5, mo.Op(3, 5))
	testlib.AclAssert(t, 5, mo.Op(5, 3))
	testlib.AclAssert(t, 0, mo.Op(-1, 0))
}

func TestMoMin(t *testing.T) {
	mo := MoMin()
	
	testlib.AclAssert(t, INF, mo.E)
	testlib.AclAssert(t, 3, mo.Op(3, 5))
	testlib.AclAssert(t, 3, mo.Op(5, 3))
	testlib.AclAssert(t, -1, mo.Op(-1, 0))
}

func TestMoSum(t *testing.T) {
	mo := MoSum[int]()
	
	testlib.AclAssert(t, 0, mo.E)
	testlib.AclAssert(t, 8, mo.Op(3, 5))
	testlib.AclAssert(t, -1, mo.Op(-5, 4))
}

func TestMoXOR(t *testing.T) {
	mo := MoXOR()
	
	testlib.AclAssert(t, 0, mo.E)
	testlib.AclAssert(t, 6, mo.Op(3, 5)) // 3 ^ 5 = 6
	testlib.AclAssert(t, 0, mo.Op(5, 5)) // 5 ^ 5 = 0
}

func TestNewSegmentTree(t *testing.T) {
	tests := map[string]struct {
		arr      []int
		monoid   *Monoid[int]
		testFunc func(*testing.T, *SegmentTree[int])
	}{
		"sum tree": {
			arr:    []int{1, 2, 3, 4, 5},
			monoid: MoSum[int](),
			testFunc: func(t *testing.T, st *SegmentTree[int]) {
				testlib.AclAssert(t, 15, st.Query(0, 5)) // Sum of all
				testlib.AclAssert(t, 6, st.Query(1, 3))  // Sum of [2, 3]
				testlib.AclAssert(t, 1, st.Query(0, 1))  // Sum of [1]
			},
		},
		"max tree": {
			arr:    []int{3, 1, 4, 1, 5},
			monoid: MoMax(),
			testFunc: func(t *testing.T, st *SegmentTree[int]) {
				testlib.AclAssert(t, 5, st.Query(0, 5)) // Max of all
				testlib.AclAssert(t, 4, st.Query(1, 3)) // Max of [1, 4]
				testlib.AclAssert(t, 3, st.Query(0, 1)) // Max of [3]
			},
		},
		"min tree": {
			arr:    []int{3, 1, 4, 1, 5},
			monoid: MoMin(),
			testFunc: func(t *testing.T, st *SegmentTree[int]) {
				testlib.AclAssert(t, 1, st.Query(0, 5)) // Min of all
				testlib.AclAssert(t, 1, st.Query(1, 3)) // Min of [1, 4]
				testlib.AclAssert(t, 3, st.Query(0, 1)) // Min of [3]
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			st := NewSegmentTree(tc.arr, tc.monoid)
			tc.testFunc(t, st)
		})
	}
}

func TestSegmentTree_Update(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	st := NewSegmentTree(arr, MoSum[int]())
	
	// Initial sum
	testlib.AclAssert(t, 15, st.Query(0, 5))
	
	// Update index 2 from 3 to 10
	st.Update(2, 10)
	testlib.AclAssert(t, 22, st.Query(0, 5)) // 1+2+10+4+5 = 22
	testlib.AclAssert(t, 12, st.Query(1, 3)) // 2+10 = 12
	
	// Update index 0 from 1 to 0
	st.Update(0, 0)
	testlib.AclAssert(t, 21, st.Query(0, 5)) // 0+2+10+4+5 = 21
}

func TestSegmentTree_QueryEdgeCases(t *testing.T) {
	arr := []int{1, 2, 3}
	st := NewSegmentTree(arr, MoSum[int]())
	
	// Empty range should return identity element
	testlib.AclAssert(t, 0, st.Query(0, 0))
	testlib.AclAssert(t, 0, st.Query(1, 1))
	testlib.AclAssert(t, 0, st.Query(3, 3))
	
	// Single element ranges
	testlib.AclAssert(t, 1, st.Query(0, 1))
	testlib.AclAssert(t, 2, st.Query(1, 2))
	testlib.AclAssert(t, 3, st.Query(2, 3))
}