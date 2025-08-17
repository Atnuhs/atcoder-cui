package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

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

func TestYesNo(t *testing.T) {
	testOut := new(bytes.Buffer)
	Out = bufio.NewWriter(testOut)

	tests := map[string]struct {
		input bool
		want  string
	}{
		"true":  {input: true, want: "Yes\n"},
		"false": {input: false, want: "No\n"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			testOut.Reset()
			YesNo(tc.input)
			Out.Flush()
			actual := testOut.String()
			if tc.want != actual {
				t.Errorf("expected: %q, but got: %q", tc.want, actual)
			}
		})
	}
}

func TestYesNoFunc(t *testing.T) {
	testOut := new(bytes.Buffer)
	Out = bufio.NewWriter(testOut)

	tests := map[string]struct {
		inputFunc func() bool
		want      string
	}{
		"function returns true": {
			inputFunc: func() bool { return true },
			want:      "Yes\n",
		},
		"function returns false": {
			inputFunc: func() bool { return false },
			want:      "No\n",
		},
		"complex function true": {
			inputFunc: func() bool { return 5 > 3 },
			want:      "Yes\n",
		},
		"complex function false": {
			inputFunc: func() bool { return 2 > 5 },
			want:      "No\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			testOut.Reset()
			YesNoFunc(tc.inputFunc)
			Out.Flush()
			actual := testOut.String()
			if tc.want != actual {
				t.Errorf("expected: %q, but got: %q", tc.want, actual)
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
