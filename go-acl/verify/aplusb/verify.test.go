// verification-helper: PROBLEM https://judge.yosupo.jp/problem/aplusb
package main

import (
	"go-acl/util"
)

func main() {
	defer util.Out.Flush()
	a, b := util.Readi(), util.Readi()

	util.Ans(a + b)
}
