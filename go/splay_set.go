package main

type SplaySet struct {
	root *SplayNode
}

func NewSplaySet() *SplaySet {
	return &SplaySet{
		root: nil,
	}
}

func (ss *SplaySet) Le(value int) (int, bool) {
	if ss.root == nil {
		return -1, false
	}

	ok, ng := -1, ss.root.size
	now := ss.root
	for Abs(ok-ng) > 1 {
		m := (ok + ng) / 2
		now = now.Get(m)
		if now.key <= value {
			ok = m
		} else {
			ng = m
		}
	}
	ss.root = now
	return ok, true
}

func (ss *SplaySet) Has(value int) bool {
	_, ok := ss.Le(value)
	if !ok {
		return false
	}
	return ss.root.key == value
}

func (ss *SplaySet) Add(value int) {
	if ss.Has(value) {
		return
	}

	nn := NewSplayNode(value, -1)

	if i, ok := ss.Le(value); ok {
		l, r := SplitSN(i, ss.root)
		ss.root = MergeSN(MergeSN(l, nn), r)
	} else {
		ss.root = nn
	}
}
