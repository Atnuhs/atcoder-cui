package main

// SplayTree は汎用的なスプレー木の実装
type SplayTree[K Ordered, V any] struct {
	root *splayNode[K, V]
	size int
}

// splayNode はスプレー木のノード
type splayNode[K Ordered, V any] struct {
	l, r, p *splayNode[K, V]
	key     K
	value   V
	size    int
}

// NewSplayTree は新しいスプレー木を作成
func NewSplayTree[K Ordered, V any]() *SplayTree[K, V] {
	return &SplayTree[K, V]{}
}

// Insert はキーと値のペアを挿入
func (st *SplayTree[K, V]) Insert(key K, value V) {
	if st.root == nil {
		st.root = &splayNode[K, V]{
			key:   key,
			value: value,
			size:  1,
		}
		st.size = 1
		return
	}
	
	oldSize := st.root.size
	st.root = st.root.insert(key, value)
	// 新しいノードが追加された場合のみサイズを増加
	if st.root.size > oldSize {
		st.size++
	}
}

// Find はキーで検索し、値と存在フラグを返す
func (st *SplayTree[K, V]) Find(key K) (V, bool) {
	if st.root == nil {
		var zero V
		return zero, false
	}
	
	node := st.root.find(key)
	if node != nil {
		st.root = node
		return node.value, true
	}
	
	var zero V
	return zero, false
}

// Delete はキーで要素を削除
func (st *SplayTree[K, V]) Delete(key K) bool {
	if st.root == nil {
		return false
	}
	
	newRoot, deleted := st.root.delete(key)
	st.root = newRoot
	if deleted {
		st.size--
		return true
	}
	return false
}

// Size は要素数を返す
func (st *SplayTree[K, V]) Size() int {
	return st.size
}

// IsEmpty は空かどうかを返す
func (st *SplayTree[K, V]) IsEmpty() bool {
	return st.size == 0
}

// InOrder は中順巡回でキーと値のペアを返す
func (st *SplayTree[K, V]) InOrder() []struct{ Key K; Value V } {
	if st.root == nil {
		return nil
	}
	return st.root.inOrder()
}

// ---- ノードのメソッド ----

func (n *splayNode[K, V]) update() {
	if n == nil {
		return
	}
	n.size = 1
	if n.l != nil {
		n.size += n.l.size
	}
	if n.r != nil {
		n.size += n.r.size
	}
}

func (n *splayNode[K, V]) splay() {
	// スプレー操作の実装（既存のSplayNodeと同様）
	for n.p != nil {
		if n.p.p == nil {
			// Zig step
			n.rotate()
		} else if (n.p.l == n) == (n.p.p.l == n.p) {
			// Zig-zig step
			n.p.rotate()
			n.rotate()
		} else {
			// Zig-zag step
			n.rotate()
			n.rotate()
		}
	}
}

func (n *splayNode[K, V]) rotate() {
	p := n.p
	if p == nil {
		return
	}
	
	pp := p.p
	if pp != nil {
		if pp.l == p {
			pp.l = n
		} else {
			pp.r = n
		}
	}
	n.p = pp
	
	if p.l == n {
		p.l = n.r
		if n.r != nil {
			n.r.p = p
		}
		n.r = p
	} else {
		p.r = n.l
		if n.l != nil {
			n.l.p = p
		}
		n.l = p
	}
	p.p = n
	
	p.update()
	n.update()
}

func (n *splayNode[K, V]) find(key K) *splayNode[K, V] {
	current := n
	for current != nil {
		if key == current.key {
			current.splay()
			return current
		} else if key < current.key {
			current = current.l
		} else {
			current = current.r
		}
	}
	return nil
}

func (n *splayNode[K, V]) insert(key K, value V) *splayNode[K, V] {
	current := n
	for {
		if key == current.key {
			// キーが既に存在する場合は値を更新
			current.value = value
			current.splay()
			return current
		} else if key < current.key {
			if current.l == nil {
				newNode := &splayNode[K, V]{
					key:   key,
					value: value,
					size:  1,
					p:     current,
				}
				current.l = newNode
				current.update()
				newNode.splay()
				return newNode
			}
			current = current.l
		} else {
			if current.r == nil {
				newNode := &splayNode[K, V]{
					key:   key,
					value: value,
					size:  1,
					p:     current,
				}
				current.r = newNode
				current.update()
				newNode.splay()
				return newNode
			}
			current = current.r
		}
	}
}

func (n *splayNode[K, V]) delete(key K) (*splayNode[K, V], bool) {
	target := n.find(key)
	if target == nil {
		return n, false
	}
	
	// targetをルートにスプレー
	root := target
	
	var newRoot *splayNode[K, V]
	if root.l == nil {
		newRoot = root.r
	} else if root.r == nil {
		newRoot = root.l
	} else {
		// 左の最大要素を見つけてルートにする
		maxLeft := root.l
		for maxLeft.r != nil {
			maxLeft = maxLeft.r
		}
		maxLeft.splay()
		
		// maxLeftの右に元の右部分木を接続
		maxLeft.r = root.r
		if root.r != nil {
			root.r.p = maxLeft
		}
		newRoot = maxLeft
	}
	
	if newRoot != nil {
		newRoot.p = nil
		newRoot.update()
	}
	
	return newRoot, true
}

func (n *splayNode[K, V]) inOrder() []struct{ Key K; Value V } {
	if n == nil {
		return nil
	}
	
	var result []struct{ Key K; Value V }
	
	// 左の子を巡回
	if n.l != nil {
		result = append(result, n.l.inOrder()...)
	}
	
	// 現在のノード
	result = append(result, struct{ Key K; Value V }{Key: n.key, Value: n.value})
	
	// 右の子を巡回
	if n.r != nil {
		result = append(result, n.r.inOrder()...)
	}
	
	return result
}