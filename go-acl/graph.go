package main

import "sort"

// WEdge は重み付き辺を表す構造体
type WEdge struct {
	from, to, weight int
}

// NewWEdge は新しい重み付き辺を生成する
func NewWEdge(from, to, weight int) *WEdge {
	return &WEdge{
		from:   from,
		to:     to,
		weight: weight,
	}
}

// Kruskal はクラスカル法を用いて最小全域木を求める
func Kruskal(n int, edges []*WEdge) (int, []*WEdge) {
	// はじめに辺を重みでソートする
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	// その後、Union-Findを用いて最小全域木を求める
	uf := NewUnionFind(n)
	ret := make([]*WEdge, 0)
	sum := 0

	// すべての辺を調べる
	for _, e := range edges {
		if uf.Family(e.from, e.to) {
			continue
		}
		ret = append(ret, e)
		sum += e.weight
		uf.Union(e.from, e.to)
	}
	if uf.Size(0) != n {
		return -1, nil
	}
	return sum, ret
}
