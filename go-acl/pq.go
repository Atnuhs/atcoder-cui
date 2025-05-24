package main

import "container/heap"

// heaapImpl はヒープの実装
type heapImpl[T Ordered] []T

func (h heapImpl[T]) Len() int { return len(h) }
func (h heapImpl[T]) Less(i, j int) bool {
	hi, hj := h[i], h[j]
	return hi < hj
}
func (h heapImpl[T]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *heapImpl[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *heapImpl[T]) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// PQ は優先度付きキューの実装
type PQ[T Ordered] struct {
	value heapImpl[T]
}

func NewPQ[T Ordered]() *PQ[T] {
	pq := &PQ[T]{}
	heap.Init(&pq.value)
	return pq
}

func (pq *PQ[T]) Push(x T) {
	heap.Push(&pq.value, x)
}

func (pq *PQ[T]) Pop() T {
	x := heap.Pop(&pq.value)
	return x.(T)
}
