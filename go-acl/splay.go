package main

import (
	"fmt"
	"strings"
)

const (
	SplaySetDefaultValue = -1 // SplaySetで使用するデフォルト値
)


// SplayNode はスプレー木のノードを表す構造体
type SplayNode struct {
	l, r, p         *SplayNode
	size            int
	key             int
	value, min, max int
}

// NewSplayNode は新しいスプレー木のノードを生成する
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
	if sn == nil {
		return ""
	}
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

func (sn *SplayNode) findAtSub(idx int) (found *SplayNode) {
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
			now.splay()
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

func (sn *SplayNode) FindAt(idx int) *SplayNode {
	node := sn.findAtSub(idx)
	if node != nil {
		node.splay()
	}
	return node
}

func (sn *SplayNode) findSub(key int) (found *SplayNode) {
	now := sn
	for now != nil {
		if now.key == key {
			now.splay()
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

func (sn *SplayNode) Find(key int) (found *SplayNode) {
	found = sn.findSub(key)
	if found != nil {
		found.splay()
	}
	return found
}

func (sn *SplayNode) Ge(key int) (idx int) {
	if sn == nil {
		return -1
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
	sn = sn.FindAt(sn.size - 1) // always found
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

func (sn *SplayNode) Insert(node *SplayNode) *SplayNode {
	if sn == nil {
		return node
	}

	idx := sn.Ge(node.key) // idx shoud be 0 <= idx <= sn.size
	if idx == sn.size {
		return sn.MergeR(node)
	}

	found := sn.FindAt(idx)
	if found.key != node.key {
		return found.InsertAt(idx, node)
	}

	return found
}

func (sn *SplayNode) Delete(node *SplayNode) (root *SplayNode, removed *SplayNode) {
	root = sn.Find(node.key)
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

// SplaySet はスプレー木のセットを表す構造体
type SplaySet struct {
	root *SplayNode
}

// NewSplaySet は新しいスプレー木のセットを生成する
func NewSplaySet(values ...int) *SplaySet {
	s := &SplaySet{
		root: nil,
	}
	for _, v := range values {
		s.Push(v)
	}
	return s
}

func NewSplaySetNode(value int) *SplayNode {
	return NewSplayNode(value, SplaySetDefaultValue)
}

func (ss *SplaySet) String() string {
	return ss.root.String()
}

func (ss *SplaySet) Push(value int) {
	ss.root = ss.root.Insert(NewSplaySetNode(value))
}

func (ss *SplaySet) Remove(value int) {
	ss.root, _ = ss.root.Delete(NewSplaySetNode(value))
}

func (ss *SplaySet) Has(value int) bool {
	found := ss.root.Find(value)
	if found == nil {
		return false
	}
	ss.root = found
	return true
}

func (ss *SplaySet) Values() (arr []int) {
	return ss.root.values()
}

func (ss *SplaySet) Size() int {
	if ss.root == nil {
		return 0
	}
	return ss.root.size
}

func (ss *SplaySet) IsEmpty() bool {
	return ss.root == nil
}

func (ss *SplaySet) At(idx int) int {
	found := ss.root.FindAt(idx)
	if found == nil {
		panic("out of index")
	}
	ss.root = found
	return ss.root.key
}

func (ss *SplaySet) Ge(value int) int {
	idx := ss.root.Ge(value)
	if 0 <= idx && idx < ss.root.size {
		ss.root = ss.root.FindAt(idx)
	}
	return idx
}

func (ss *SplaySet) Gt(value int) int {
	if ss.root == nil {
		return -1
	}
	return ss.Ge(value + 1)
}

func (ss *SplaySet) Le(value int) int {
	if ss.root == nil {
		return -1
	}
	return ss.Ge(value+1) - 1
}

func (ss *SplaySet) Lt(value int) int {
	if ss.root == nil {
		return -1
	}
	return ss.Ge(value) - 1
}

// SplayMap はスプレー木のマップを表す構造体
type SplayMap struct {
	root *SplayNode
}

// NewSplayMap は新しいスプレー木のマップを生成する
func NewSplayMap() *SplayMap {
	return &SplayMap{
		root: nil,
	}
}

func NewSplayMapNode(key int, value int) *SplayNode {
	return NewSplayNode(key, value)
}

func (ss *SplayMap) Push(key int, value int) {
	node := NewSplayMapNode(key, value)
	if ss.root == nil {
		ss.root = node
		return
	}
	ss.root = ss.root.Insert(node)
}

func (ss *SplayMap) Remove(key int) int {
	var removed *SplayNode
	ss.root, removed = ss.root.Delete(NewSplayMapNode(key, SplaySetDefaultValue))
	if removed != nil {
		return removed.value
	}
	return 0
}

func (ss *SplayMap) Has(key int) bool {
	found := ss.root.Find(key)
	if found == nil {
		return false
	}
	ss.root = found
	return true
}

func (ss *SplayMap) Values() (arr []int) {
	return ss.root.values()
}

func (ss *SplayMap) Size() int {
	if ss.root == nil {
		return 0
	}
	return ss.root.size
}

func (ss *SplayMap) IsEmpty() bool {
	return ss.root == nil
}

func (ss *SplayMap) At(key int) int {
	found := ss.root.Find(key)
	if found == nil {
		return 0
	}
	ss.root = found
	return ss.root.value
}

func (ss *SplayMap) String() string {
	if ss.root == nil {
		return ""
	}
	return ss.root.String()
}

func (ss *SplayMap) Ge(value int) int {
	if ss.root == nil {
		return 0
	}
	idx := ss.root.Ge(value)
	found := ss.root.findAtSub(idx)
	if found != nil {
		ss.root = found
		return ss.root.key
	}
	return 0
}

func (ss *SplayMap) Gt(value int) int {
	return ss.Ge(value + 1)
}

func (ss *SplayMap) Le(value int) int {
	return ss.Ge(value+1) - 1
}

func (ss *SplayMap) Lt(value int) int {
	return ss.Ge(value) - 1
}
