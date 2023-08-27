---
data:
  _extendedDependsOn:
  - icon: ':heavy_check_mark:'
    path: go-acl/main.go
    title: go-acl/main.go
  - icon: ':heavy_check_mark:'
    path: go-acl/splay/map.go
    title: go-acl/splay/map.go
  - icon: ':heavy_check_mark:'
    path: go-acl/splay/node.go
    title: go-acl/splay/node.go
  - icon: ':heavy_check_mark:'
    path: go-acl/splay/node_test.go
    title: go-acl/splay/node_test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/splay/set.go
    title: go-acl/splay/set.go
  - icon: ':heavy_check_mark:'
    path: go-acl/splay/set_test.go
    title: go-acl/splay/set_test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/testlib/assert.go
    title: go-acl/testlib/assert.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/lib.go
    title: go-acl/util/lib.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/lib_test.go
    title: go-acl/util/lib_test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/math_test.go
    title: go-acl/util/math_test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/monoid.go
    title: go-acl/util/monoid.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/segmentTree.go
    title: go-acl/util/segmentTree.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/sieve.go
    title: go-acl/util/sieve.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/sieve_test.go
    title: go-acl/util/sieve_test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/unionfind.go
    title: go-acl/util/unionfind.go
  - icon: ':heavy_check_mark:'
    path: go-acl/verify/aplusb/verify.test.go
    title: go-acl/verify/aplusb/verify.test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/verify/associative_array/verify.test.go
    title: go-acl/verify/associative_array/verify.test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/verify/many_aplusb/verify.test.go
    title: go-acl/verify/many_aplusb/verify.test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/verify/predecessor_problem/verify.test.go
    title: go-acl/verify/predecessor_problem/verify.test.go
  _extendedRequiredBy:
  - icon: ':heavy_check_mark:'
    path: go-acl/main.go
    title: go-acl/main.go
  - icon: ':heavy_check_mark:'
    path: go-acl/splay/map.go
    title: go-acl/splay/map.go
  - icon: ':heavy_check_mark:'
    path: go-acl/splay/node.go
    title: go-acl/splay/node.go
  - icon: ':heavy_check_mark:'
    path: go-acl/splay/node_test.go
    title: go-acl/splay/node_test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/splay/set.go
    title: go-acl/splay/set.go
  - icon: ':heavy_check_mark:'
    path: go-acl/splay/set_test.go
    title: go-acl/splay/set_test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/testlib/assert.go
    title: go-acl/testlib/assert.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/lib.go
    title: go-acl/util/lib.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/lib_test.go
    title: go-acl/util/lib_test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/math_test.go
    title: go-acl/util/math_test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/monoid.go
    title: go-acl/util/monoid.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/segmentTree.go
    title: go-acl/util/segmentTree.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/sieve.go
    title: go-acl/util/sieve.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/sieve_test.go
    title: go-acl/util/sieve_test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/unionfind.go
    title: go-acl/util/unionfind.go
  _extendedVerifiedWith:
  - icon: ':heavy_check_mark:'
    path: go-acl/verify/aplusb/verify.test.go
    title: go-acl/verify/aplusb/verify.test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/verify/associative_array/verify.test.go
    title: go-acl/verify/associative_array/verify.test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/verify/many_aplusb/verify.test.go
    title: go-acl/verify/many_aplusb/verify.test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/verify/predecessor_problem/verify.test.go
    title: go-acl/verify/predecessor_problem/verify.test.go
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':heavy_check_mark:'
  attributes: {}
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: go-acl/util/math.go\n"
  code: "package util\n\nimport (\n\t\"math\"\n)\n\n// ModPow return x^e % mod\nfunc\
    \ ModPow(x, e, mod int) int {\n\ti := 1\n\tret := 1\n\n\tfor i <= e {\n\t\tif\
    \ i&e > 0 {\n\t\t\tret = (ret * x) % mod\n\t\t}\n\t\ti <<= 1\n\t\tx = (x * x)\
    \ % mod\n\t}\n\treturn ret\n}\n\nfunc Inv(x, mod int) int {\n\treturn ModPow(x,\
    \ mod-2, mod)\n}\n\nfunc Gcd(a, b int) int {\n\tif b == 0 {\n\t\treturn a\n\t\
    }\n\treturn Gcd(b, a%b)\n}\n\nfunc Lcm(a, b int) int {\n\treturn a / Gcd(a, b)\
    \ * b\n}\n\nfunc Sqrt(x int) int {\n\treturn int(math.Sqrt(float64(x)))\n}\n\n\
    // NextPerm returns [1,2,3,4] => [1,2,4,3] ... [4,3,2,1]\nfunc NextPerm(a []int)\
    \ bool {\n\t// search i\n\ti := len(a) - 2\n\tfor i >= 0 && a[i] >= a[i+1] {\n\
    \t\ti--\n\t}\n\tif i < 0 {\n\t\treturn false\n\t}\n\tj := len(a) - 1\n\tfor j\
    \ >= 0 && a[j] <= a[i] {\n\t\tj--\n\t}\n\n\ta[i], a[j] = a[j], a[i]\n\n\tl :=\
    \ i + 1\n\tr := len(a) - 1\n\tfor l < r {\n\t\ta[l], a[r] = a[r], a[l]\n\t\tl++\n\
    \t\tr--\n\t}\n\treturn true\n}\n\n// Extrema returns min, max\nfunc Extrema[T\
    \ Ordered](vals ...T) (T, T) {\n\tmi, ma := vals[0], vals[0]\n\tfor _, v := range\
    \ vals {\n\t\tif v < mi {\n\t\t\tmi = v\n\t\t}\n\t\tif v > ma {\n\t\t\tma = v\n\
    \t\t}\n\t}\n\treturn mi, ma\n}\n\nfunc Max[T Ordered](vals ...T) T {\n\t_, ma\
    \ := Extrema(vals...)\n\treturn ma\n}\n\nfunc Min[T Ordered](vals ...T) T {\n\t\
    mi, _ := Extrema(vals...)\n\treturn mi\n}\n\nfunc Sum[T Ordered](vals ...T) T\
    \ {\n\tvar sum T\n\tfor _, v := range vals {\n\t\tsum += v\n\t}\n\treturn sum\n\
    }\n\nfunc Abs(x int) int {\n\tif x < 0 {\n\t\tx = -x\n\t}\n\treturn x\n}\n\n//\
    \ IsPrime is O(Sqrt(N))\nfunc IsPrime(x int) bool {\n\tif x == 1 {\n\t\treturn\
    \ false\n\t}\n\n\trx := Sqrt(x)\n\tfor i := 2; i <= rx; i++ {\n\t\tif x%i == 0\
    \ {\n\t\t\treturn false\n\t\t}\n\t}\n\treturn true\n}\n\n// Factorize is O(Sqrt(N))\n\
    // got, ret\n// 6, []Pair{{2,1}, {3.1}}\nfunc Factorize(x int) []*Pair[int] {\n\
    \tif x == 1 {\n\t\treturn []*Pair[int]{}\n\t}\n\n\trx := Sqrt(x)\n\tn := x\n\t\
    ret := make([]*Pair[int], 0)\n\tfor i := 2; i <= rx; i++ {\n\t\tif n%i != 0 {\n\
    \t\t\tcontinue\n\t\t}\n\t\texp := 0\n\t\tfor n%i == 0 {\n\t\t\tn /= i\n\t\t\t\
    exp++\n\t\t}\n\t\tret = append(ret, NewPair(i, exp))\n\t}\n\tif n != 1 {\n\t\t\
    ret = append(ret, NewPair(n, 1))\n\t}\n\treturn ret\n}\n\n// Mobius is O(sqrt(n))\
    \ returns\n// 0 <= 4, 12, 18, 50\n// 1 <= 1, 6, 210\n// -1 <= 2, 30, 140729\n\
    func Mobius(x int) int {\n\tret := 1\n\n\trx := Sqrt(x)\n\tn := x\n\tfor i :=\
    \ 2; i <= rx; i++ {\n\t\tif n%i != 0 {\n\t\t\tcontinue\n\t\t}\n\n\t\tif (n/i)%i\
    \ == 0 {\n\t\t\treturn 0\n\t\t}\n\t\tn /= i\n\t\tret = -ret\n\t}\n\n\tif n !=\
    \ 1 {\n\t\tret = -ret\n\t}\n\treturn ret\n}\n\n// Divisors is O(sqrt(n)) returns\n\
    // 2 => 1, 2\n// 10 => 1, 2, 5, 10\nfunc Divisors(x int) []int {\n\tret := make([]int,\
    \ 0)\n\n\trx := Sqrt(x)\n\tfor i := 1; i <= rx; i++ {\n\t\tif x%i != 0 {\n\t\t\
    \tcontinue\n\t\t}\n\t\tret = append(ret, i)\n\t\tif i != x/i {\n\t\t\tret = append(ret,\
    \ x/i)\n\t\t}\n\t}\n\treturn ret\n}\n\n// CountDivisors is O(sqrt(n)) returns\n\
    // 1 => 1\n// 2 => 2\n// 10 => 4\nfunc CountDivisors(pairs []*Pair[int]) int {\n\
    \tans := 1\n\tfor _, pe := range pairs {\n\t\tans *= (pe.v + 1)\n\t}\n\treturn\
    \ ans\n}\n"
  dependsOn:
  - go-acl/splay/node.go
  - go-acl/splay/set_test.go
  - go-acl/splay/node_test.go
  - go-acl/splay/set.go
  - go-acl/splay/map.go
  - go-acl/verify/predecessor_problem/verify.test.go
  - go-acl/verify/aplusb/verify.test.go
  - go-acl/verify/associative_array/verify.test.go
  - go-acl/verify/many_aplusb/verify.test.go
  - go-acl/testlib/assert.go
  - go-acl/util/unionfind.go
  - go-acl/util/sieve.go
  - go-acl/util/math_test.go
  - go-acl/util/lib_test.go
  - go-acl/util/lib.go
  - go-acl/util/segmentTree.go
  - go-acl/util/sieve_test.go
  - go-acl/util/monoid.go
  - go-acl/main.go
  isVerificationFile: false
  path: go-acl/util/math.go
  requiredBy:
  - go-acl/splay/node.go
  - go-acl/splay/set_test.go
  - go-acl/splay/node_test.go
  - go-acl/splay/set.go
  - go-acl/splay/map.go
  - go-acl/testlib/assert.go
  - go-acl/util/unionfind.go
  - go-acl/util/sieve.go
  - go-acl/util/math_test.go
  - go-acl/util/lib_test.go
  - go-acl/util/lib.go
  - go-acl/util/segmentTree.go
  - go-acl/util/sieve_test.go
  - go-acl/util/monoid.go
  - go-acl/main.go
  timestamp: '2023-08-28 00:47:54+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - go-acl/verify/predecessor_problem/verify.test.go
  - go-acl/verify/aplusb/verify.test.go
  - go-acl/verify/associative_array/verify.test.go
  - go-acl/verify/many_aplusb/verify.test.go
documentation_of: go-acl/util/math.go
layout: document
redirect_from:
- /library/go-acl/util/math.go
- /library/go-acl/util/math.go.html
title: go-acl/util/math.go
---
