package splay

import (
	"fmt"
	"strings"

	. "go-acl/util"
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
	if sn == nil {
		return -1
	}
	if sn.l != nil {
		return sn.l.size
	}
	return 0
}

func (sn *SplayNode) update() {
	if sn == nil {
		return
	}
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
	if sn == nil {
		return
	}
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
	if sn == nil {
		return
	}
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
	if sn == nil {
		return ret
	}
	if sn.l != nil {
		ret = append(ret, sn.l.values()...)
	}
	ret = append(ret, sn.key)
	if sn.r != nil {
		ret = append(ret, sn.r.values()...)
	}
	return ret
}

func (sn *SplayNode) String() string {
	ret := strings.Builder{}
	ret.WriteString("(")
	if sn.l != nil {
		ret.WriteString(fmt.Sprintf("%s ", sn.l.String()))
	}
	ret.WriteString(fmt.Sprint(sn.key))
	if sn.r != nil {
		ret.WriteString(fmt.Sprintf(" %s", sn.r.String()))
	}
	ret.WriteString(")")
	return ret.String()
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

func (sn *SplayNode) FindAt(idx int) (found *SplayNode) {
	if sn == nil {
		return nil
	}
	if idx < 0 || sn.size <= idx {
		return nil
	}
	// n include [0, n)
	now := sn
	for now != nil {
		switch {
		case idx == now.index():
			return now
		case idx < now.index():
			now = now.l
		case idx > now.index():
			idx -= now.index() + 1
			now = now.r
		}
	}
	panic("must not reach this code")
}

func (sn *SplayNode) FindAtAndSplay(idx int) *SplayNode {
	node := sn.FindAt(idx)
	node.splay()
	return node
}

func (sn *SplayNode) Find(key int) (found *SplayNode) {
	now := sn
	for now != nil {
		if now.key == key {
			return now
		}

		if now.key > key {
			now = now.l
		} else {
			now = now.r
		}
	}
	return nil
}

func (sn *SplayNode) FindAndSplay(key int) (found *SplayNode) {
	found = sn.Find(key)
	found.splay()
	return found
}

func (sn *SplayNode) Has(key int) bool {
	found := sn.Find(key)
	if found == nil {
		return false
	}
	return found.key == key
}

func (sn *SplayNode) Ge(key int) (idx int) {
	if sn == nil {
		return 0
	}
	now := sn
	idx = sn.size
	i := 0
	for now != nil {
		if now.key >= key {
			idx = Min(idx, i+now.index())
			now = now.l
		} else {
			i += now.index() + 1
			now = now.r
		}
	}
	return idx
}

func (sn *SplayNode) MergeR(rroot *SplayNode) *SplayNode {
	if rroot == nil {
		return sn
	}
	if sn == nil {
		return rroot
	}
	sn = sn.FindAtAndSplay(sn.size - 1) // always found
	sn.r = rroot
	rroot.p = sn
	sn.update()
	return sn
}

func (sn *SplayNode) Split(idx int) (*SplayNode, *SplayNode) {
	if sn == nil {
		return nil, nil
	}
	if idx == sn.size {
		return sn, nil
	}

	rroot := sn.FindAtAndSplay(idx)
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

func (sn *SplayNode) Insert(node *SplayNode) *SplayNode {
	idx := sn.Ge(node.key)
	if found := sn.FindAt(idx); found != nil {
		if found.key == node.key {
			return sn
		}
	}
	return sn.InsertAt(idx, node)
}

func (sn *SplayNode) Delete(node *SplayNode) (root *SplayNode, removed *SplayNode) {
	root = sn.FindAndSplay(node.key)
	if root == nil {
		// target not found
		return sn, nil
	}
	if root.key == node.key {
		// target found
		root, removed = root.DeleteAt(root.index())
	}
	// target not found
	return root, removed
}
