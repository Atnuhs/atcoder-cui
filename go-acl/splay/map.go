package splay

type SplayMap struct {
	root *SplayNode
}

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
	}
	ss.root = ss.root.Insert(node)
}

func (ss *SplayMap) Remove(key int) int {
	var removed *SplayNode
	ss.root, removed = ss.root.Delete(NewSplayMapNode(key, -1))
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
	idx := ss.root.Ge(value)
	ss.root = ss.root.FindAtSub(idx)
	return ss.root.key
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
