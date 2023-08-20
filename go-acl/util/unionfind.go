package util

type UnionFind struct {
	data []int
}

func NewUnionFind(n int) *UnionFind {
	data := make([]int, n)
	for i := range data {
		data[i] = -1
	}
	return &UnionFind{
		data: data,
	}
}

func (uf *UnionFind) Root(x int) int {
	if uf.data[x] < 0 {
		return x
	} else {
		uf.data[x] = uf.Root(uf.data[x])
		return uf.data[x]
	}
}

func (uf *UnionFind) Family(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

func (uf *UnionFind) Size(x int) int {
	return -uf.data[uf.Root(x)]
}

func (uf *UnionFind) Union(x, y int) {
	rx := uf.Root(x)
	ry := uf.Root(y)

	if rx == ry {
		return
	}

	if uf.Size(rx) < uf.Size(ry) {
		rx = rx ^ ry
		ry = rx ^ ry
		rx = rx ^ ry
	}

	uf.data[rx] += uf.data[ry]
	uf.data[ry] = rx
}
