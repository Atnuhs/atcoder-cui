package main

import (
	"sort"
	"testing"
)

func TestPQ(t *testing.T) {
	p := NewPQOrdered[int]()

	values := []int{3, 4, 1, 2, 5, 8, 7, 6, -121}
	for _, v := range values {
		p.Push(v)
	}
	got := make([]int, p.Len())
	for i := range got {
		got[i] = p.Pop()
	}
	sort.Ints(values)
	for i := range got {
		if got[i] != values[i] {
			t.Errorf("invalid at index %d want %d, got %d", i, values[i], got[i])
		}
	}
}
