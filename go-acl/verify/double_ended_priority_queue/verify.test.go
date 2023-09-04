// verification-helper: PROBLEM https://judge.yosupo.jp/problem/double_ended_priority_queue
package main

import . "go-acl/util"

func main() {
	defer Out.Flush()

	n, q := Readi(), Readi()
	pq := NewDEPQ(Readis(n)...)

	for iq := 0; iq < q; iq++ {
		t := Readi()
		switch t {
		case 0:
			x := Readi()
			pq.Push(x)
		case 1:
			Ans(pq.PopMin())
		case 2:
			Ans(pq.PopMax())
		}
	}
}
