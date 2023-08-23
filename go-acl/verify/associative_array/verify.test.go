// verification-helper: PROBLEM https://judge.yosupo.jp/problem/associative_array
package main

import (
	. "go-acl/util"
)

func main() {
	defer Out.Flush()

	q := Readi()
	mp := make(map[int]int)

	for i := 0; i < q; i++ {
		t := Readi()
		switch t {
		case 0:
			k, v := Readi(), Readi()
			mp[k] = v
		case 1:
			k := Readi()
			if val, ok := mp[k]; ok {
				Ans(val)
			} else {
				Ans(0)
			}
		}
	}
}
