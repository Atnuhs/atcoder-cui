package main

import (
	"math/rand"
	"testing"

	"github.com/Atnuhs/atcoder-cui/go-acl/testlib"
)

func TestSplaySet_Push(t *testing.T) {
	tests := map[string]struct {
		pushValues []int
		want       []int
	}{
		"single value": {
			pushValues: []int{10},
			want:       []int{10},
		},

		"sorted values": {
			pushValues: []int{10, 20, 30},
			want:       []int{10, 20, 30},
		},
		"unsorted values": {
			pushValues: []int{30, 10, 20},
			want:       []int{10, 20, 30},
		},
		"unsorted minus values": {
			pushValues: []int{-30, -10, -20},
			want:       []int{-30, -20, -10},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewSplaySet(tc.pushValues...)

			got := s.Values()
			testlib.AclAssert(t, tc.want, got)
		})
	}
}

func TestSplaySet_Remove(t *testing.T) {
	tests := map[string]struct {
		pushValues   []int
		deleteValues []int
		want         []int
	}{
		"single push, single delete": {
			pushValues:   []int{10},
			deleteValues: []int{10},
			want:         []int{},
		},
		"multi push, single delete": {
			pushValues:   []int{20, 10, 30},
			deleteValues: []int{10},
			want:         []int{20, 30},
		},
		"multi push, multi delete": {
			pushValues:   []int{20, 10, 30},
			deleteValues: []int{30, 20},
			want:         []int{10},
		},
		"single push, multi delete, no value left": {
			pushValues:   []int{20},
			deleteValues: []int{20, 10},
			want:         []int{},
		},
		"single push, multi delete, value remains": {
			pushValues:   []int{20},
			deleteValues: []int{30, 10, 40, 50, 50},
			want:         []int{20},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewSplaySet(tc.pushValues...)

			for _, v := range tc.deleteValues {
				s.Remove(v)
			}

			got := s.Values()
			testlib.AclAssert(t, tc.want, got)
		})
	}
}

func TestSplaySet_Has(t *testing.T) {
	tests := map[string]struct {
		pushValues []int
		trialValue int
		want       bool
	}{
		"set has value": {
			pushValues: []int{10, 20, 30},
			trialValue: 20,
			want:       true,
		},
		"not set has value": {
			pushValues: []int{10, 20, 30},
			trialValue: 40,
			want:       false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewSplaySet(tc.pushValues...)
			got := s.Has(tc.trialValue)

			testlib.AclAssert(t, tc.want, got)
		})
	}
}

func TestSplaySet_Ge(t *testing.T) {
	pushValues := []int{10, 20, 30}
	tests := map[string]struct {
		pushValues []int
		trialValue int
		want       int
	}{
		"minimmum": {
			pushValues: pushValues,
			trialValue: -INF,
			want:       0,
		},
		"less than first value": {
			pushValues: pushValues,
			trialValue: 9,
			want:       0,
		},
		"equal to first value": {
			pushValues: pushValues,
			trialValue: 10,
			want:       0,
		},
		"more than first value": {
			pushValues: pushValues,
			trialValue: 11,
			want:       1,
		},
		"less than last value": {
			pushValues: pushValues,
			trialValue: 29,
			want:       2,
		},
		"equal to last value": {
			pushValues: pushValues,
			trialValue: 30,
			want:       2,
		},
		"more than last value": {
			pushValues: pushValues,
			trialValue: 31,
			want:       3,
		},
		"maximum": {
			pushValues: pushValues,
			trialValue: INF,
			want:       3,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewSplaySet(tc.pushValues...)
			got := s.Ge(tc.trialValue)
			t.Log(s)
			testlib.AclAssert(t, tc.want, got)
			if s.Size() != len(tc.pushValues) {
				t.Errorf("expected size: %d but got %d", len(tc.pushValues), s.Size())
			}
		})
	}
}

func TestSplaySet_Le(t *testing.T) {
	pushValues := []int{10, 20, 30}
	tests := map[string]struct {
		pushValues []int
		trialValue int
		want       int
	}{
		"minimmum": {
			pushValues: pushValues,
			trialValue: -INF,
			want:       -1,
		},
		"less than first value": {
			pushValues: pushValues,
			trialValue: 9,
			want:       -1,
		},
		"equal to first value": {
			pushValues: pushValues,
			trialValue: 10,
			want:       0,
		},
		"more than first value": {
			pushValues: pushValues,
			trialValue: 11,
			want:       0,
		},
		"less than last value": {
			pushValues: pushValues,
			trialValue: 29,
			want:       1,
		},
		"equal to last value": {
			pushValues: pushValues,
			trialValue: 30,
			want:       2,
		},
		"more than last value": {
			pushValues: pushValues,
			trialValue: 31,
			want:       2,
		},
		"maximum": {
			pushValues: pushValues,
			trialValue: INF,
			want:       2,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewSplaySet(tc.pushValues...)
			got := s.Le(tc.trialValue)
			t.Log(s)
			testlib.AclAssert(t, tc.want, got)
			if s.Size() != len(tc.pushValues) {
				t.Errorf("expected size: %d but got %d", len(tc.pushValues), s.Size())
			}
		})
	}
}

func BenchmarkSplaySet_Insert(b *testing.B) {
	s := NewSplaySet(ILF(100000, func(i int) int { return rand.Intn(100000) })...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vPush, vRemove := rand.Intn(b.N), rand.Intn(b.N)
		s.Push(vPush)
		s.Remove(vRemove)
	}
}

func BenchmarkSplaySet_Remove(b *testing.B) {
	s := NewSplaySet(ILF(100000, func(i int) int { return rand.Intn(100000) })...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vPush, vRemove := rand.Intn(b.N), rand.Intn(b.N)
		s.Push(vPush)
		s.Remove(vRemove)
	}
}
