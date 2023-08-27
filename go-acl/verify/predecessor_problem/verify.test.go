// verification-helper: PROBLEM https://judge.yosupo.jp/problem/predecessor_problem
package main

import (
	"go-acl/splay"
	. "go-acl/util"
)

func main() {
	defer Out.Flush()

	_, q := Readi(), Readi()
	t := Reads()

	a := splay.NewSplaySet()
	for i, v := range t {
		if v == '1' {
			a.Push(i)
		}
	}

	for i := 0; i < q; i++ {
		c, k := Readi(), Readi()
		switch c {
		case 0:
			if !a.Has(k) {
				a.Push(k)
			}
		case 1:
			if a.Has(k) {
				a.Remove(k)
			}
		case 2:
			if a.Has(k) {
				Ans(1)
			} else {
				Ans(0)
			}
		case 3:
			v := a.Ge(k)
			if 0 <= v && v < a.Size() {
				Ans(a.At(v))
			} else {
				Ans(-1)
			}
		case 4:
			v := a.Le(k)
			if 0 <= v && v < a.Size() {
				Ans(a.At(v))
			} else {
				Ans(-1)
			}
		}
	}
}
