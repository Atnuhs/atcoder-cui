package main

import (
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

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

func TestSplayNode_splay(t *testing.T) {
	n := N
	expected, actual := generate(n)
	for i := range actual {
		node := actual[i]
		node.splay()
		if node.p != nil {
			t.Errorf("splayed node should be root but not root")
		}
		if node.size != n {
			t.Errorf("expected size: %d but got %d", n, node.size)
		}
		assertValues(t, node, expected)

	}
}

func TestSplayNode_values(t *testing.T) {
	n := N
	expected, actual := generate(n)
	root := actual[n-1]
	assertValues(t, root, expected)
	if root.size != n {
		t.Errorf("expected node size: %d but got: %d", n, actual[0].size)
	}
}

func TestSplayNode_Get(t *testing.T) {
	n := N
	expected, actual := generate(n)

	root := actual[n-1]
	for i := range actual {
		root = root.Get(i)

		if root == nil {
			t.Fatalf("root should not be nil in safe index access %d: %d", i, len(actual))
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

func generate2(t *testing.T, n int) (expected []int, actual *SplayNode) {
	actual = nil
	expected = make([]int, 0, n)

	for i := 0; i < n; i++ {
		nn := NewSplayNode(i, -1)
		if actual == nil {
			actual = nn
			expected = append(expected, i)
		} else {
			j := rand.Intn(actual.size + 1)
			actual = actual.Insert(j, nn)
			expected = insert(j, i, expected)
		}
		assertValues(t, actual, expected)
	}
	t.Logf("\n%s", actual.describe(0))
	return expected, actual
}

func TestSplayNode_Insert(t *testing.T) {
	n := N
	expected, root := generate2(t, n)

	for i := 0; i < n; i++ {
		root = root.Get(i)
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

func TestSplayNode_Delete(t *testing.T) {
	n := N
	expected, root := generate2(t, n)

	for i := 0; i < n; i++ {
		j := rand.Intn(root.size)
		root, _ = root.Delete(j)
		expected = delete(j, expected)
		assertValues(t, root, expected)
	}
}
