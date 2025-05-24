package main

import "fmt"

// Pair は2つの値を持つ構造体
type Pair[T any] struct {
	u, v T
}

// NewPair Pairを生成する
func NewPair[T any](u, v T) *Pair[T] {
	return &Pair[T]{u, v}
}

// String Pairの文字列を、空白区切りで返す
func (p *Pair[T]) String() string {
	return fmt.Sprintf("%v %v", p.u, p.v)
}
