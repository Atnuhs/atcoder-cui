package acl

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestSplayset_BasicOperations(t *testing.T) {
	s := NewSplayset[int]()

	if !s.IsEmpty() {
		t.Errorf("new set should be empty")
	}

	s.Add(3)
	s.Add(1)
	s.Add(4)
	s.Add(1) // 重複: 無視されるべき
	s.Add(5)

	if s.Size() != 4 {
		t.Errorf("expected size 4, got %d", s.Size())
	}

	if !s.Contains(3) {
		t.Errorf("expected to contain 3")
	}
	if s.Contains(2) {
		t.Errorf("expected not to contain 2")
	}

	if !s.Delete(3) {
		t.Errorf("expected to delete 3")
	}
	if s.Delete(3) {
		t.Errorf("expected delete(3) to fail second time")
	}
	if s.Contains(3) {
		t.Errorf("expected 3 to be deleted")
	}
	if s.Size() != 3 {
		t.Errorf("expected size 3 after delete, got %d", s.Size())
	}
}

func TestSplayset_InOrder(t *testing.T) {
	s := NewSplayset[int]()
	for _, k := range []int{5, 2, 8, 1, 7} {
		s.Add(k)
	}
	got := s.InOrder()
	want := []int{1, 2, 5, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InOrder: want %v got %v", want, got)
	}
}

func TestSplayset_At(t *testing.T) {
	s := NewSplayset[int]()
	for _, k := range []int{30, 10, 20, 40} {
		s.Add(k)
	}
	cases := []struct {
		k      int
		key    int
		ok     bool
	}{
		{-1, 0, false},
		{0, 10, true},
		{1, 20, true},
		{2, 30, true},
		{3, 40, true},
		{4, 0, false},
	}
	for _, c := range cases {
		key, ok := s.At(c.k)
		if ok != c.ok || (ok && key != c.key) {
			t.Errorf("At(%d): want (%d,%v) got (%d,%v)", c.k, c.key, c.ok, key, ok)
		}
	}
}

func TestSplayset_LowerUpperBound(t *testing.T) {
	rng := rand.New(rand.NewSource(7))
	keys := rng.Perm(500)
	s := NewSplayset[int]()
	for _, k := range keys {
		s.Add(k)
	}
	sorted := make([]int, len(keys))
	copy(sorted, keys)
	sort.Ints(sorted)

	for q := -3; q <= 502; q++ {
		// LowerBound = sort.SearchInts (smallest i where a[i] >= q)
		wantLB := sort.SearchInts(sorted, q)
		if got := s.LowerBound(q); got != wantLB {
			t.Errorf("LowerBound(%d): want %d got %d", q, wantLB, got)
		}
		// UpperBound = smallest i where a[i] > q
		wantUB := sort.SearchInts(sorted, q+1)
		if got := s.UpperBound(q); got != wantUB {
			t.Errorf("UpperBound(%d): want %d got %d", q, wantUB, got)
		}
	}
}

func TestSplayset_EmptyBounds(t *testing.T) {
	s := NewSplayset[int]()
	if s.LowerBound(0) != 0 {
		t.Errorf("LowerBound on empty set should be 0")
	}
	if s.UpperBound(0) != 0 {
		t.Errorf("UpperBound on empty set should be 0")
	}
	if got := s.InOrder(); got != nil {
		t.Errorf("InOrder on empty set should be nil, got %v", got)
	}
}
