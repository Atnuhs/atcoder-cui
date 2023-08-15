package main

import (
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const N = 10

func generate(n int) ([]*SplayNode, []int) {
	actual := make([]*SplayNode, n)
	expected := make([]int, 0, n)

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
	return actual, expected
}

func assertValues(t *testing.T, root *SplayNode, expected []int) {
	t.Helper()
	actual := root.values()
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Log(expected)
		t.Log(actual)
		t.Errorf("%s", diff)
	}
}

func TestSplayNode_splay(t *testing.T) {
	n := N
	actual, expected := generate(n)
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
	actual, expected := generate(n)
	root := actual[n-1]
	assertValues(t, root, expected)
	if root.size != n {
		t.Errorf("expected node size: %d but got: %d", n, actual[0].size)
	}
}

func TestSplayNode_Get(t *testing.T) {
	n := N
	actual, expected := generate(n)

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

func TestSplay_Insert_Delete(t *testing.T) {
	n := N
	var root *SplayNode = nil

	for i := 0; i < n; i++ {
		root = MergeSN(root, NewSplayNode(-1, i))
		t.Logf("\nInsert: %d/%d\n%s", i, n, root.describe(0))
	}

	t.Log(root.describe(0))

	for i := 0; i < n; i++ {
		root = root.Get(i)
		t.Logf("\nGet: %d/%d\n%s", i, n, root.describe(0))
	}

	for i := 0; i < n; i++ {
		root = InsertSN(i*2, root, NewSplayNode(1, -i))
		t.Logf("\nAdd oddIndex:%d/%d\n%s", i, n, root.describe(0))
	}

	for {
		t.Logf("##POP %d", root.size/2)
		left, del := DeleteSN(root.size/2, root)
		root = left
		if root == nil {
			break
		}
		t.Logf("\nDelete: %d\n%s", root.size, root.describe(0))
		t.Logf("\nPoped: %d\n%s", del.size, del.describe(0))
	}
}
