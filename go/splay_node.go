package main

import (
	"fmt"
	"strings"
)

type SplayNode struct {
	l, r, p         *SplayNode
	size            int
	key             int
	value, min, max int
}

func NewSplayNode(key, value int) *SplayNode {
	ret := &SplayNode{
		l:     nil,
		r:     nil,
		p:     nil,
		key:   key,
		value: value,
	}
	ret.update()
	return ret
}

func (sn *SplayNode) index() int {
	if sn.l != nil {
		return sn.l.size
	}
	return 0
}

func (sn *SplayNode) update() {
	sn.size = 1
	sn.min = sn.value
	sn.max = sn.value

	if sn.l != nil {
		sn.size += sn.l.size
		sn.min = Min(sn.min, sn.l.min)
		sn.max = Max(sn.max, sn.l.max)
	}
	if sn.r != nil {
		sn.size += sn.r.size
		sn.min = Min(sn.min, sn.r.min)
		sn.max = Max(sn.max, sn.r.max)
	}
}

func (sn *SplayNode) state() int {
	if sn.p == nil {
		return 0
	}
	if sn.p.l == sn {
		return 1
	}
	if sn.p.r == sn {
		return -1
	}
	return INF
}

func (sn *SplayNode) rotate() {
	ns := sn.state()
	if ns == 0 {
		return
	}

	p := sn.p
	ps := p.state()

	// edge 1
	pp := p.p
	switch ps {
	case 1:
		pp.l = sn
	case -1:
		pp.r = sn
	}
	sn.p = pp

	// edge 2, 3
	var c *SplayNode
	switch ns {
	case 1:
		c = sn.r
		sn.r = p
		p.l = c
	case -1:
		c = sn.l
		sn.l = p
		p.r = c
	}

	p.p = sn
	if c != nil {
		c.p = p
	}
	p.update()
	sn.update()
}

func (sn *SplayNode) splay() {
	for sn.p != nil {
		// sn is not root

		if sn.p.state() == 0 {
			// sn.p is root
			sn.rotate()
			continue
		}

		if sn.state() == sn.p.state() {
			sn.p.rotate()
			sn.rotate()
		} else {
			sn.rotate()
			sn.rotate()
		}
	}
}

func (sn *SplayNode) values() []int {
	ret := make([]int, 0)
	if sn.l != nil {
		ret = append(ret, sn.l.values()...)
	}
	ret = append(ret, sn.key)
	if sn.r != nil {
		ret = append(ret, sn.r.values()...)
	}
	return ret
}

func (sn *SplayNode) describe(rank int) string {
	ret := ""
	if sn.r != nil {
		ret += sn.r.describe(rank + 1)
	}
	ret += fmt.Sprintf(
		strings.Repeat("    ", rank)+"-[k:%d, v:%d, sz: %d, rank: %d]\n",
		sn.key,
		sn.value,
		sn.size,
		rank,
	)

	if sn.l != nil {
		ret += sn.l.describe(rank + 1)
	}
	return ret
}

func (sn *SplayNode) maxRank(rank int) int {
	ret := rank
	if sn.r != nil {
		ret = Max(ret, sn.r.maxRank(rank+1))
	}
	if sn.l != nil {
		ret = Max(ret, sn.l.maxRank(rank+1))
	}
	return ret
}

func (sn *SplayNode) FindAt(idx int) *SplayNode {
	if idx < 0 || sn.size <= idx {
		return nil
	}
	// n include [0, n)
	for sn.index() != idx {
		switch {
		case idx < sn.index():
			sn = sn.l
		case idx > sn.index():
			idx -= sn.index() + 1
			sn = sn.r
		}
	}
	sn.splay()
	return sn
}

func (sn *SplayNode) Ge(val int) int {
	now := sn
	ret := sn.size
	i := 0
	for now != nil {
		if now.key >= val {
			ret = Min(ret, i+now.index())
			now = now.l
		} else {
			i += now.index() + 1
			now = now.r
		}
	}
	return ret
}

func (sn *SplayNode) MergeR(rroot *SplayNode) *SplayNode {
	if rroot == nil {
		return sn
	}
	if sn == nil {
		panic("sn is nil")
	}
	sn = sn.FindAt(sn.size - 1) // always found
	sn.r = rroot
	rroot.p = sn
	sn.update()
	return sn
}

func (sn *SplayNode) Split(idx int) (*SplayNode, *SplayNode) {
	if idx == sn.size {
		return sn, nil
	}

	rroot := sn.FindAt(idx)
	if rroot == nil {
		// idx is out of index
		return nil, nil
	}

	lroot := rroot.l
	if lroot != nil {
		lroot.p = nil
	}
	rroot.l = nil

	rroot.update()
	// lroot not need to update()
	return lroot, rroot
}

func (sn *SplayNode) InsertAt(idx int, node *SplayNode) *SplayNode {
	lroot, rroot := sn.Split(idx)
	if lroot == nil {
		return node.MergeR(rroot)
	} else {
		return lroot.MergeR(node).MergeR(rroot)
	}
}

func (sn *SplayNode) DeleteAt(idx int) (root *SplayNode, dropped *SplayNode) {
	lroot, rroot := sn.Split(idx)
	if rroot == nil {
		return lroot, nil
	}
	del, rroot := rroot.Split(1)
	if lroot == nil {
		return rroot, del
	} else {
		root = lroot.MergeR(rroot)
		return root, del
	}
}
