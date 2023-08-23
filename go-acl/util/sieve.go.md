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
    path: go-acl/util/lib.go
    title: go-acl/util/lib.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/lib_test.go
    title: go-acl/util/lib_test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/math.go
    title: go-acl/util/math.go
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
    path: go-acl/util/lib.go
    title: go-acl/util/lib.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/lib_test.go
    title: go-acl/util/lib_test.go
  - icon: ':heavy_check_mark:'
    path: go-acl/util/math.go
    title: go-acl/util/math.go
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
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':heavy_check_mark:'
  attributes: {}
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: go-acl/util/sieve.go\n"
  code: "package util\n\n// Eratosthenes sieve\ntype EratosthenesSieve struct {\n\t\
    isPrime   []bool\n\tminFactor []int\n\tmobius    []int\n}\n\n// NewSieve is O(N\
    \ loglog N)\nfunc NewSieve(n int) *EratosthenesSieve {\n\tisPrime := make([]bool,\
    \ n+1)\n\tminFactor := make([]int, n+1)\n\tmobius := make([]int, n+1)\n\n\tfor\
    \ i := range isPrime {\n\t\tisPrime[i] = true\n\t\tminFactor[i] = -1\n\t\tmobius[i]\
    \ = 1\n\t}\n\n\tisPrime[0] = false\n\tisPrime[1] = false\n\tminFactor[1] = 1\n\
    \n\t// sieve\n\tfor i := range isPrime {\n\t\tif !isPrime[i] {\n\t\t\tcontinue\n\
    \t\t}\n\n\t\tminFactor[i] = i\n\t\tmobius[i] = -1\n\n\t\tfor j := i * 2; j <=\
    \ n; j += i {\n\t\t\tisPrime[j] = false\n\n\t\t\tif minFactor[j] == -1 {\n\t\t\
    \t\tminFactor[j] = i\n\t\t\t}\n\n\t\t\tif (j/i)%i == 0 {\n\t\t\t\tmobius[j] =\
    \ 0\n\t\t\t} else {\n\t\t\t\tmobius[j] = -mobius[j]\n\t\t\t}\n\t\t}\n\t}\n\treturn\
    \ &EratosthenesSieve{\n\t\tisPrime:   isPrime,\n\t\tminFactor: minFactor,\n\t\t\
    mobius:    mobius,\n\t}\n}\n\n// IsPrime is O(1)\nfunc (sv *EratosthenesSieve)\
    \ IsPrime(x int) bool {\n\treturn sv.isPrime[x]\n}\n\n// Factorize is O(Sqrt(1))\n\
    // got, ret\n// 6, []Pair{{2,1}, {3.1}}\nfunc (sv *EratosthenesSieve) Factorize(x\
    \ int) []*Pair[int] {\n\tret := make([]*Pair[int], 0)\n\tn := x\n\tfor n > 1 {\n\
    \t\tp := sv.minFactor[n]\n\t\texp := 0\n\n\t\tfor sv.minFactor[n] == p {\n\t\t\
    \tn /= p\n\t\t\texp++\n\t\t}\n\t\tret = append(ret, NewPair(p, exp))\n\t}\n\t\
    return ret\n}\n\n// Mobius is O(1) return\n// 0 <= 4, 12, 18, 50\n// 1 <= 1, 6,\
    \ 210\n// -1 <= 2, 30, 140729\nfunc (sv *EratosthenesSieve) Mobius(x int) int\
    \ {\n\treturn sv.mobius[x]\n}\n\n// Divisors is O(sqrt(n)) returns\n// 2 => 1,\
    \ 2\n// 10 => 1, 2, 5, 10\nfunc (sv *EratosthenesSieve) Divisors(x int) []int\
    \ {\n\tret := []int{1}\n\n\tf := sv.Factorize(x)\n\tfor _, pe := range f {\n\t\
    \tn := len(ret)\n\t\tfor i := 0; i < n; i++ {\n\t\t\tv := 1\n\t\t\tfor j := 0;\
    \ j < pe.v; j++ {\n\t\t\t\tv *= pe.u\n\t\t\t\tret = append(ret, ret[i]*v)\n\t\t\
    \t}\n\t\t}\n\t}\n\treturn ret\n}\n\n// CountDivisors is O(1) returns len(sv.Divisors(x))\n\
    // 1 => 1\n// 2 => 2\n// 10 => 4\nfunc (sv *EratosthenesSieve) CountDivisors(x\
    \ int) int {\n\treturn CountDivisors(sv.Factorize(x))\n}\n"
  dependsOn:
  - go-acl/splay/node.go
  - go-acl/splay/node_test.go
  - go-acl/splay/set.go
  - go-acl/splay/map.go
  - go-acl/verify/aplusb/verify.test.go
  - go-acl/verify/associative_array/verify.test.go
  - go-acl/verify/many_aplusb/verify.test.go
  - go-acl/util/unionfind.go
  - go-acl/util/math_test.go
  - go-acl/util/lib_test.go
  - go-acl/util/lib.go
  - go-acl/util/segmentTree.go
  - go-acl/util/math.go
  - go-acl/util/sieve_test.go
  - go-acl/util/monoid.go
  - go-acl/main.go
  isVerificationFile: false
  path: go-acl/util/sieve.go
  requiredBy:
  - go-acl/splay/node.go
  - go-acl/splay/node_test.go
  - go-acl/splay/set.go
  - go-acl/splay/map.go
  - go-acl/util/unionfind.go
  - go-acl/util/math_test.go
  - go-acl/util/lib_test.go
  - go-acl/util/lib.go
  - go-acl/util/segmentTree.go
  - go-acl/util/math.go
  - go-acl/util/sieve_test.go
  - go-acl/util/monoid.go
  - go-acl/main.go
  timestamp: '2023-08-24 01:49:17+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - go-acl/verify/aplusb/verify.test.go
  - go-acl/verify/associative_array/verify.test.go
  - go-acl/verify/many_aplusb/verify.test.go
documentation_of: go-acl/util/sieve.go
layout: document
redirect_from:
- /library/go-acl/util/sieve.go
- /library/go-acl/util/sieve.go.html
title: go-acl/util/sieve.go
---
