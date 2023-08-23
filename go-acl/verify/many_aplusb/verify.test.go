// verification-helper: PROBLEM https://judge.yosupo.jp/problem/many_aplusb
package main

import "go-acl/util"

func main() {
	defer util.Out.Flush()

	t := util.Readi()
	ans := make([]int, t)
	for i := 0; i < t; i++ {
		a, b := util.Readi(), util.Readi()
		ans[i] = a + b
	}

	for i := range ans {
		util.Ans(ans[i])
	}
}
