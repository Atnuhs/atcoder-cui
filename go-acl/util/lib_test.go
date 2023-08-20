package util

import "testing"

func TestPopBack(t *testing.T) {
	testCases := []struct {
		desc string
		data []int
	}{
		{
			desc: "[1,2,3,4,5]",
			data: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			n := len(tc.data)
			a := make([]int, n)
			copy(a, tc.data)

			for i := n - 1; i >= 0; i-- {
				got := PopBack(&a)
				want := tc.data[i]
				if want != got {
					t.Errorf("index: %d, expected %d, but got %d", i, want, got)
				}
			}
		})
	}
}

func TestPopFront(t *testing.T) {
	testCases := []struct {
		desc string
		data []int
	}{
		{
			desc: "[1,2,3,4,5]",
			data: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			n := len(tc.data)
			a := make([]int, n)
			copy(a, tc.data)

			for i := 0; i < n; i++ {
				got := PopFront(&a)
				want := tc.data[i]
				if want != got {
					t.Errorf("index: %d, expected %d, but got %d", i, want, got)
				}
			}
		})
	}
}
