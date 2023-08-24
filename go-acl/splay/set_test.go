package splay

import (
	"testing"

	"go-acl/testlib"
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
			s := NewSplayNodeFromSlice(tc.pushValues)

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
			s := NewSplayNodeFromSlice(tc.pushValues)

			for _, v := range tc.deleteValues {
				s.Remove(v)
			}

			got := s.Values()
			testlib.AclAssert(t, tc.want, got)
		})
	}
}
