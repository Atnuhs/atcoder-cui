package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/exp/slices"
)

func TestSplayNode_splay(t *testing.T) {
	values := struct {
		pp, p, t, ppc, pc, tcl, tcr int
	}{
		pp:  1,
		p:   2,
		t:   3,
		ppc: 4,
		pc:  5,
		tcl: 6,
		tcr: 7,
	}
	type treeStruct struct {
		ppLeft bool
		pLeft  bool
	}

	genArr := func(ts treeStruct) []interface{} {
		t := []interface{}{values.tcl, values.t, values.tcr}
		var p []interface{}
		if ts.pLeft {
			p = append(t, values.p, values.pc)
		} else {
			p = append([]interface{}{values.pc, values.p}, t...)
		}

		var pp []interface{}
		if ts.ppLeft {
			pp = append(p, values.pp, values.ppc)
		} else {
			pp = append([]interface{}{values.ppc, values.pp}, p...)
		}
		return pp
	}

	expected := func(format string, ts treeStruct) string {
		return fmt.Sprintf(format, genArr(ts)...)
	}

	tests := map[string]struct {
		treeStrcut  treeStruct
		expectedFmt string
	}{
		"p is left child of pp and t is left child of p": {
			treeStrcut:  treeStruct{ppLeft: true, pLeft: true},
			expectedFmt: "((%d) %d ((%d) %d ((%d) %d (%d))))",
		},
		"p is left child of pp and t is right child of p": {
			treeStrcut:  treeStruct{ppLeft: true, pLeft: false},
			expectedFmt: "(((%d) %d (%d)) %d ((%d) %d (%d)))",
		},
		"p is right child of pp and t is left child of p": {
			treeStrcut:  treeStruct{ppLeft: false, pLeft: true},
			expectedFmt: "(((%d) %d (%d)) %d ((%d) %d (%d)))",
		},
		"p is right child of pp and t is right child of p": {
			treeStrcut:  treeStruct{ppLeft: false, pLeft: false},
			expectedFmt: "((((%d) %d (%d)) %d (%d)) %d (%d))",
		},
	}

	for caseName, test := range tests {
		t.Run(caseName, func(t *testing.T) {
			target := NewSplayNode(values.t, -1)
			target.l = NewSplayNode(values.tcl, -1)
			target.r = NewSplayNode(values.tcr, -1)
			target.l.p = target
			target.r.p = target
			p := NewSplayNode(values.p, -1)
			pc := NewSplayNode(values.pc, -1)
			target.p = p
			pc.p = p
			if test.treeStrcut.pLeft {
				p.l = target
				p.r = pc
			} else {
				p.l = pc
				p.r = target
			}

			pp := NewSplayNode(values.pp, -1)
			ppc := NewSplayNode(values.ppc, -1)
			p.p = pp
			ppc.p = pp
			if test.treeStrcut.ppLeft {
				pp.l = p
				pp.r = ppc
			} else {
				pp.l = ppc
				pp.r = p
			}

			t.Log(pp.String())
			target.splay()
			expected := expected(test.expectedFmt, test.treeStrcut)
			if diff := cmp.Diff(expected, target.String()); diff != "" {
				t.Errorf("\nexpected:\n%s\nactual:\n%s", expected, target.String())
				t.Errorf("\ndiff: %s", diff)
			}
		})
	}
}

const N = 10

func generate(n int) (expected []int, actual []*SplayNode) {
	actual = make([]*SplayNode, n)
	expected = make([]int, 0, n)

	for i := range actual {
		actual[i] = &SplayNode{
			key: i,
		}
		actual[i].update()

		if i == 0 {
			expected = append(expected, 0)
			continue
		}

		switch rand.Intn(2) {
		case 0:
			actual[i].r = actual[i-1]
			expected = append([]int{i}, expected...)
		case 1:
			actual[i].l = actual[i-1]
			expected = append(expected, i)

		}
		actual[i-1].p = actual[i]
		actual[i].update()
	}
	return expected, actual
}

func generateRandomTestCase(t *testing.T, n int) (expected []int, actual *SplayNode) {
	actual = nil
	expected = make([]int, 0, n)

	for i := 0; i < n; i++ {
		nn := NewSplayNode(i, -1)
		if actual == nil {
			actual = nn
			expected = append(expected, i)
		} else {
			j := rand.Intn(actual.size + 1)
			actual = actual.InsertAt(j, nn)
			expected = insert(j, i, expected)
		}
		assertValues(t, actual, expected)
	}
	t.Logf("\n%s", actual.describe(0))
	return expected, actual
}

func assertValues(t *testing.T, root *SplayNode, expected []int) {
	t.Helper()
	var actual []int
	if root != nil {
		actual = root.values()
	}
	opts := []cmp.Option{
		cmpopts.EquateEmpty(),
	}
	if diff := cmp.Diff(expected, actual, opts...); diff != "" {
		t.Logf("Expected: %v", expected)
		t.Logf("Actual  : %v", actual)
		t.Errorf("%s", diff)
	}
}

func TestSplayNode_values(t *testing.T) {
	n := N
	expected, root := generateRandomTestCase(t, n)
	assertValues(t, root, expected)
	if root.size != n {
		t.Errorf("expected node size: %d but got: %d", n, root.size)
	}
}

func TestSplayNode_FindAt(t *testing.T) {
	n := N
	expected, root := generateRandomTestCase(t, n)

	for i := 0; i < n; i++ {
		root = root.FindAt(i)

		if root == nil {
			t.Fatalf("root should not be nil in safe index access %d: %d", i, n)
		}

		if root.p != nil {
			t.Fatal("return value of Get should be root but has parent")
		}
		assertValues(t, root, expected)
	}
}

func insert(i int, v int, arr []int) []int {
	l := arr[:i]
	var r []int
	if i < len(arr) {
		r = arr[i:]
	}
	return append(l, append([]int{v}, r...)...)
}

func TestSplayNode_InsertAt(t *testing.T) {
	n := N
	expected, root := generateRandomTestCase(t, n)

	for i := 0; i < n; i++ {
		root = root.FindAt(i)
		assertValues(t, root, expected)
	}
}

func delete(i int, arr []int) []int {
	l := arr[:i]
	var r []int
	if i+1 < len(arr) {
		r = arr[i+1:]
	}
	return append(l, r...)
}

func TestSplayNode_DeleteAt(t *testing.T) {
	n := N
	expected, root := generateRandomTestCase(t, n)

	for i := 0; i < n; i++ {
		j := rand.Intn(root.size)
		root, _ = root.DeleteAt(j)
		expected = delete(j, expected)
		assertValues(t, root, expected)
	}
}

func TestSplayNode_maxRank(t *testing.T) {
	n := 10
	m := 10
	result := make([][]int, n+1)
	for in := 0; in <= n; in++ {
		nn := 1 << in
		result[in] = make([]int, m+1)
		for im := 0; im <= m; im++ {
			t.Log(in, im)
			mm := 1 << im
			root := NewSplayNode(0, -1)
			for i := 1; i < nn; i++ {
				root = root.InsertAt(root.size, NewSplayNode(i, -1))
			}

			for i := 0; i < mm; i++ {
				root = root.FindAt(rand.Intn(root.size))
			}

			result[in][im] = root.maxRank(0)
		}
	}

	for i := range result {
		t.Log(result[i])
	}
}

func TestSplayNode_Ge(t *testing.T) {
	n := N
	m := N * 100
	expected := make([]int, n)
	for i := range expected {
		expected[i] = rand.Intn(n)
	}
	sort.Ints(expected)

	root := NewSplayNode(expected[0], -1)
	for i := 1; i < n; i++ {
		root = root.InsertAt(root.size, NewSplayNode(expected[i], -1))
	}

	for i := 0; i < m; i++ {
		j := rand.Intn(root.size)
		root = root.FindAt(j)
		if i&(i-1) == 0 {
			t.Logf("i: %d, rank: %d", i, root.maxRank(0))
		}

	}
	// t.Logf("\n%s", root.describe(0))

	assertValues(t, root, expected)
	for i := 0; i < m; i++ {
		x := rand.Intn(n)
		expectedI := sort.Search(n, func(i int) bool { return expected[i] >= x })
		actualI := root.Ge(x)
		if expectedI != actualI {
			t.Log(expected)
			t.Log(root.values())
			t.Errorf("expected: %d but got %d at %d", expectedI, actualI, x)
		}
	}
}

func TestSplayNode_Insert(t *testing.T) {
	n := N
	expected := make([]int, n)
	for i := range expected {
		expected[i] = rand.Intn(n)
	}
	root := NewSplayNode(expected[0], -1)
	for i := 1; i < n; i++ {
		root = root.Insert(NewSplayNode(expected[i], -1))
	}
	slices.Sort(expected)
	expected = slices.Compact(expected)
	assertValues(t, root, expected)
}

func TestSplayNode_Delete(t *testing.T) {
	type expected struct {
		hasDropped bool
		dropped    int
		rest       []int
	}

	tests := map[string]struct {
		data     []int
		delete   int
		expected expected
	}{
		"delete an existing value": {
			data:     []int{1, 2, 3, 4, 5},
			delete:   1,
			expected: expected{hasDropped: true, dropped: 1, rest: []int{2, 3, 4, 5}},
		},
		"delete a non-existing value": {
			data:     []int{1, 2, 3, 4, 5},
			delete:   6,
			expected: expected{hasDropped: false, dropped: 0, rest: []int{1, 2, 3, 4, 5}},
		},
		"delete an existing value and no rest": {
			data:     []int{1},
			delete:   1,
			expected: expected{hasDropped: true, dropped: 1, rest: nil},
		},
	}

	for caseName, test := range tests {
		t.Run(caseName, func(t *testing.T) {
			// construct splay tree
			root := NewSplayNode(test.data[0], -1)
			for i := 1; i < len(test.data); i++ {
				root = root.Insert(NewSplayNode(test.data[i], -1))
			}

			// do delete
			root, actualDropped := root.Delete(NewSplayNode(test.delete, -1))

			// check dropped
			if test.expected.hasDropped != (actualDropped != nil) {
				t.Errorf("expected has dropped?: %t, but actual has dropped?: %t", test.expected.hasDropped, actualDropped != nil)
			}

			if !test.expected.hasDropped {
				return
			}

			expectedDropped := NewSplayNode(test.expected.dropped, -1)
			if diff := cmp.Diff(expectedDropped.key, actualDropped.key); diff != "" {
				t.Errorf("expected dropped value is %v but actual is %v", test.expected.dropped, actualDropped)
				t.Error(diff)
			}

			// check rest
			assertValues(t, root, test.expected.rest)
		})
	}
}
