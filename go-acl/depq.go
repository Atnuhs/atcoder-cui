package main

// lIdx　は左の子ノードのインデックスを返す
func lIdx(idx int) int {
	return idx & ^1
}

// rIdx　は右の子ノードのインデックスを返す
func rIdx(idx int) int {
	return idx | 1
}

// pIdx　は親ノードのインデックスを返す
func pIdx(idx int) int {
	return ((idx >> 1) - 1) & ^1
}

// cIdx　は子ノードのインデックスを返す
func cIdx(idx int) int {
	return (idx & ^1)<<1 | 2 | idx&1
}

// DEPQ は二重優先度キューの実装
// 二重優先度キューは、最大値と最小値をO(logN)で取得できるデータ構造
type DEPQ[T Ordered] struct {
	values []T
}

// NewDEPQ は二重優先度キューを初期化する
func NewDEPQ[T Ordered](values ...T) *DEPQ[T] {
	pq := &DEPQ[T]{
		values: values,
	}
	for i := pq.Size() - 1; i >= 0; i-- {
		if i&1 == 1 && pq.values[i-1] < pq.values[i] {
			pq.values[i-1], pq.values[i] = pq.values[i], pq.values[i-1]
		}
		idx := pq.down(i)
		pq.upAt(idx, i)
	}
	return pq
}

func (pq *DEPQ[T]) Size() int {
	return len(pq.values)
}

func (pq *DEPQ[T]) Empty() bool {
	return len(pq.values) == 0
}

func (pq *DEPQ[T]) Push(x T) {
	pq.values = append(pq.values, x)
	pq.up(pq.Size() - 1)
}

func (pq *DEPQ[T]) GetMax() T {
	return pq.values[0]
}

func (pq *DEPQ[T]) GetMin() T {
	if pq.Size() < 2 {
		return pq.values[0]
	}
	return pq.values[1]
}

func (pq *DEPQ[T]) PopMax() T {
	ret := pq.GetMax()
	pq.values[0] = pq.values[pq.Size()-1]
	pq.values = pq.values[:pq.Size()-1]
	idx := pq.down(0)
	pq.up(idx)
	return ret
}

func (pq *DEPQ[T]) PopMin() T {
	ret := pq.GetMin()
	if pq.Size() < 2 {
		pq.values = []T{}
		return ret
	}
	pq.values[1] = pq.values[pq.Size()-1]
	pq.values = pq.values[:pq.Size()-1]
	idx := pq.down(1)
	pq.up(idx)
	return ret
}

func (pq *DEPQ[T]) upAt(idx, root int) {
	l, r := lIdx(idx), rIdx(idx)

	// sould be value[l] >= value[r]
	if r < pq.Size() && pq.values[l] < pq.values[r] {
		pq.values[l], pq.values[r] = pq.values[r], pq.values[l]
		idx ^= 1
	}

	for p := pIdx(idx); idx > root && pq.values[p] < pq.values[idx]; idx, p = p, pIdx(p) {
		// max heap
		pq.values[idx], pq.values[p] = pq.values[p], pq.values[idx]
	}

	for p := pIdx(idx) | 1; idx > root && pq.values[p] > pq.values[idx]; idx, p = p, pIdx(p)|1 {
		// min heap
		pq.values[idx], pq.values[p] = pq.values[p], pq.values[idx]
	}
}

func (pq *DEPQ[T]) up(idx int) {
	pq.upAt(idx, 1)
}

func (pq *DEPQ[T]) down(idx int) int {
	if idx&1 == 1 {
		// min heap
		for c := cIdx(idx); c < pq.Size(); idx, c = c, cIdx(c) {
			if c+2 < pq.Size() && pq.values[c] > pq.values[c+2] {
				c += 2
			}
			if pq.values[c] < pq.values[idx] {
				pq.values[idx], pq.values[c] = pq.values[c], pq.values[idx]
			}
		}
	} else {
		// max heap
		for c := cIdx(idx); c < pq.Size(); idx, c = c, cIdx(c) {
			if c+2 < pq.Size() && pq.values[c] < pq.values[c+2] {
				c += 2
			}
			if pq.values[c] > pq.values[idx] {
				pq.values[idx], pq.values[c] = pq.values[c], pq.values[idx]
			}
		}
	}
	return idx
}
