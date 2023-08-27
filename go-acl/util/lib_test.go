package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

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

func TestAns(t *testing.T) {
	testOut := new(bytes.Buffer)
	Out = bufio.NewWriter(testOut)
	testCases := map[string]struct {
		data     []interface{}
		expected string
	}{
		"only int":    {data: []interface{}{1, 2, 3}, expected: "1 2 3\n"},
		"only string": {data: []interface{}{"a", "b", "c"}, expected: "a b c\n"},
		"only []int":  {data: []interface{}{[]int{1, 2, 3}}, expected: "1 2 3\n"},
		"combined":    {data: []interface{}{1, 2, 3, "4", "a", []int{5, 6, 7}}, expected: "1 2 3 4 a 5 6 7\n"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			testOut.Reset()
			Ans(tc.data...)
			Out.Flush()
			actual := testOut.String()
			if tc.expected != actual {
				t.Errorf("expected: %q, but got: %q", tc.expected, actual)
			}
		})
	}
}

func BenchmarkOutputToOut(b *testing.B) {
	text := strings.Repeat("a", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Fprintln(Out, text)
	}
}

func BenchmarkOutputToDiscard(b *testing.B) {
	text := strings.Repeat("a", 100)
	Discard := bufio.NewWriter(io.Discard)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Fprintln(Discard, text)
	}
}
