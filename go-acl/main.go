package main

func main() {
	defer Out.Flush()

	h, w := II()
	g := Rs(h)
	passed := MakeGridOf(h, w, -1)
	a, b, c, d := IIII()
	a--
	b--
	c--
	d--

	todo := NewPQ(func(i, j *Pair[*Pair[int, int], int]) bool { return i.v < j.v })
	todo.Push(NewPair(NewPair(a, b), 0))
	dxy := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for !todo.IsEmpty() {
		now := todo.Pop()
		p := now.u
		cnt := now.v

		if p.u < 0 || h <= p.u || p.v < 0 || w <= p.v {
			continue
		}
		if passed[p.u][p.v] >= 0 {
			continue
		}
		passed[p.u][p.v] = cnt
		Ans("#")
		for _, v := range passed {
			Ans(Map(v, func(v int) any {
				if v == -1 {
					return "#"
				}
				return v
			}))
		}
		if p.u == c && p.v == d {
			Ans(cnt)
			return
		}

		if g[p.u][p.v] == '#' {
			for _, d := range dxy {
				u, v := p.u+d[0], p.v+d[1]
				if 0 <= u || u < h || 0 <= v || v < w || g[u][v] == '.' {
					todo.Push(NewPair(NewPair(u, v), cnt))
				}
			}
			continue
		}

		for _, d := range dxy {
			u, v := p.u+d[0], p.v+d[1]
			if u < 0 || h <= u || v < 0 || w <= v {
				continue
			}

			if g[u][v] == '.' {
				todo.Push(NewPair(NewPair(u, v), cnt))
			} else {
				todo.Push(NewPair(NewPair(u, v), cnt+1))
				u2, v2 := u+d[0], v+d[1]
				if 0 <= u2 || u2 < h || 0 <= v2 || v2 < w || g[u2][v2] == '#' {
					todo.Push(NewPair(NewPair(u2, v2), cnt+1))
				}
			}
		}
	}
}
