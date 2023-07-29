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
	return 0
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
	for sn.state() == 0 {
		// sn is root
		return
	}

	if sn.p.state() == 0 {
		// sn.p is root
		sn.rotate()
		return
	}

	if sn.state() == sn.p.state() {
		sn.p.rotate()
		sn.rotate()
	} else {
		sn.rotate()
		sn.rotate()
	}
}

func (sn *SplayNode) describe(rank int) string {
	ret := ""
	if sn.r != nil {
		ret += sn.r.describe(rank + 1)
	}
	ret += fmt.Sprintf(
		strings.Repeat("    ", rank)+"-[k:%d, v:%d, sz: %d]\n",
		sn.key,
		sn.value,
		sn.size,
	)

	if sn.l != nil {
		ret += sn.l.describe(rank + 1)
	}
	return ret
}

func get_subSN(ind int, node *SplayNode) (int, *SplayNode) {
	if node == nil {
		return -1, nil
	}
	ls := 0
	if node.l != nil {
		ls = node.l.size
	}

	switch {
	case ind < ls:
		return ind, node.l
	case ind == ls:
		return -1, node
	case ind > ls:
		return ind - (ls + 1), node.r
	}
	return -1, nil
}

func GetSN(ind int, node *SplayNode) *SplayNode {
	for ind != -1 {
		ind, node = get_subSN(ind, node)
	}
	// node found
	if node != nil {
		node.splay()
	}
	return node
}

func MergeSN(lroot, rroot *SplayNode) *SplayNode {
	if lroot == nil {
		return rroot
	}
	if rroot == nil {
		return lroot
	}
	lroot = GetSN(lroot.size-1, lroot) // always found
	lroot.r = rroot
	rroot.p = lroot
	lroot.update()
	return lroot
}

func SplitSN(ind int, root *SplayNode) (*SplayNode, *SplayNode) {
	if root == nil {
		return nil, nil
	}
	if ind == root.size {
		return root, nil
	}

	rroot := GetSN(ind, root)
	if rroot == nil {
		// rroot not found
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

func InsertSN(ind int, root *SplayNode, node *SplayNode) *SplayNode {
	lroot, rroot := SplitSN(ind, root)
	return MergeSN(MergeSN(lroot, node), rroot)
}

func DeleteSN(ind int, root *SplayNode) (*SplayNode, *SplayNode) {
	lroot, rroot := SplitSN(ind, root)
	del, rroot := SplitSN(1, rroot)
	root = MergeSN(lroot, rroot)
	return root, del
}
