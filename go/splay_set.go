package main

type SplaySet struct {
	root *SplayNode
}

func NewSplaySet() *SplaySet {
	return &SplaySet{
		root: nil,
	}
}

func NewSplaySetNode(value int) *SplayNode {
	return NewSplayNode(value, -1)
}

func (ss *SplaySet) Push(value int) {
	ss.root = ss.root.Insert(NewSplaySetNode(value))
}

func (ss *SplaySet) Remove(value int) {
	ss.root, _ = ss.root.Delete(NewSplaySetNode(value))
}

func (ss *SplaySet) Has(value int) bool {
	return ss.root.Has(value)
}

func (ss *SplaySet) Values() (arr []int) {
	return ss.root.values()
}

func (ss *SplaySet) Size() int {
	return ss.root.size
}

func (ss *SplaySet) IsEmpty() bool {
	return ss.root == nil
}

func (ss *SplaySet) Ge(value int) int {
	idx := ss.root.Ge(value)
	ss.root = ss.root.FindAt(idx)
	return ss.root.key
}

func (ss *SplaySet) Gt(value int) int {
	return ss.Ge(value + 1)
}

func (ss *SplaySet) Le(value int) int {
	return ss.Ge(value+1) - 1
}

func (ss *SplaySet) Lt(value int) int {
	return ss.Ge(value) - 1
}
