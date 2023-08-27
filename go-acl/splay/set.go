package splay

type SplaySet struct {
	root *SplayNode
}

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
	return NewSplayNode(value, -1)
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
