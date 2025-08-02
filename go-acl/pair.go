package main

import "fmt"

// Pair は2つの値を持つ構造体
type Pair[A any, B any] struct {
	U A
	V B
}

// NewPair Pairを生成する
func NewPair[A any, B any](u A, v B) *Pair[A, B] {
	return &Pair[A, B]{u, v}
}

// String Pairの文字列を、空白区切りで返す
func (p *Pair[A, B]) String() string {
	return fmt.Sprintf("%v %v", p.U, p.V)
}

type Pos struct {
	H, W int
}

func NewPos(h, w int) *Pos {
	return &Pos{h, w}
}

func (p *Pos) Neighbors4() []*Pos {
	return []*Pos{
		NewPos(p.H-1, p.W),
		NewPos(p.H, p.W-1),
		NewPos(p.H+1, p.W),
		NewPos(p.H, p.W+1),
	}
}

func (p *Pos) Neighbors8() []*Pos {
	return []*Pos{
		NewPos(p.H-1, p.W),
		NewPos(p.H, p.W-1),
		NewPos(p.H+1, p.W),
		NewPos(p.H, p.W+1),
		NewPos(p.H-1, p.W-1),
		NewPos(p.H-1, p.W+1),
		NewPos(p.H+1, p.W-1),
		NewPos(p.H+1, p.W+1),
	}
}

func (p *Pos) Neighbors4In(w, h int) []*Pos {
	return Filter(p.Neighbors4(), func(p *Pos) bool { return InGrid(p, h, w) })
}

func (p *Pos) Neighbors8In(w, h int) []*Pos {
	return Filter(p.Neighbors8(), func(p *Pos) bool { return InGrid(p, h, w) })
}

func InGrid(p *Pos, h, w int) bool {
	return InArea(p, 0, h, 0, w)
}

func InArea(p *Pos, hl, hr, wl, wr int) bool {
	return InRange(p.H, hl, hr) && InRange(p.W, wl, wr)
}
