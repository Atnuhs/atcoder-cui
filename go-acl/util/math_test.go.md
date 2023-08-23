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
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':heavy_check_mark:'
  attributes: {}
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: go-acl/util/math_test.go\n"
  code: "package util\n\nimport \"testing\"\n\nfunc TestInv(t *testing.T) {\n\ttestCases\
    \ := []struct {\n\t\tdesc string\n\t\tx    int\n\t\tp    int\n\t}{\n\t\t{desc:\
    \ \"x:20, p:7\", x: 20, p: 7},\n\t\t{desc: \"x:1234567, p:10^9+7\", x: 1234567,\
    \ p: 1000000007},\n\t\t{desc: \"x:1234567, p:998244353\", x: 1234567, p: 998244353},\n\
    \t}\n\tfor _, tc := range testCases {\n\t\tt.Run(tc.desc, func(t *testing.T) {\n\
    \t\t\tinvX := Inv(tc.x, tc.p)\n\n\t\t\t// x * invX = should be 1\n\t\t\tgot :=\
    \ (tc.x * invX) % tc.p\n\t\t\tif got != 1 {\n\t\t\t\tt.Errorf(\"actual should\
    \ be 1 but got %d, invX: %d\", got, invX)\n\t\t\t}\n\t\t})\n\t}\n}\n\nfunc FuzzInv(f\
    \ *testing.F) {\n\tf.Add(4, 7)\n\tf.Add(1000000007, 97)\n\tf.Add(97, 1000000007)\n\
    \tf.Add(4, 1)\n\tf.Add(4, -11)\n\tf.Fuzz(func(f *testing.T, x, mod int) {\n\t\t\
    if !IsPrime(mod) || mod <= 1 {\n\t\t\treturn\n\t\t}\n\n\t\tif Gcd(x, mod) != 1\
    \ {\n\t\t\treturn\n\t\t}\n\n\t\tinvX := Inv(x, mod)\n\t\tgot := (invX * x) % mod\n\
    \n\t\tif got != 1 {\n\t\t\tf.Errorf(\"expected 1, but got %d, x: %d, mod: %d,\
    \ invX: %d\", got, x, mod, invX)\n\t\t}\n\t})\n}\n\nfunc TestGcd(t *testing.T)\
    \ {\n\ttestCases := []struct {\n\t\tdesc       string\n\t\tx, y, want int\n\t\
    }{\n\t\t{desc: \"gcd(2, 2) => 2\", x: 2, y: 2, want: 2},\n\t\t{desc: \"gcd(4,\
    \ 2) => 2\", x: 4, y: 2, want: 2},\n\t\t{desc: \"gcd(4, 6) => 2\", x: 4, y: 6,\
    \ want: 2},\n\t\t{desc: \"gcd(11, 13) => 1\", x: 11, y: 13, want: 1},\n\t\t{desc:\
    \ \"gcd(11, 13) => 1\", x: 11, y: 13, want: 1},\n\t}\n\tfor _, tc := range testCases\
    \ {\n\t\tt.Run(tc.desc, func(t *testing.T) {\n\t\t\tgot := Gcd(tc.x, tc.y)\n\t\
    \t\tif tc.want != got {\n\t\t\t\tt.Errorf(\"expected %d but got %d\", tc.want,\
    \ got)\n\t\t\t}\n\t\t})\n\t}\n}\n\nfunc FuzzGcd(f *testing.F) {\n\tf.Add(0, 100000)\n\
    \tf.Add(100000, 0)\n\tf.Add(0, 0)\n\tf.Add(12345678, 1000000007)\n\tf.Add(-12345678,\
    \ 1000000007)\n\tf.Add(-12345678, 0)\n\tf.Add(0, -12345678)\n\tf.Fuzz(func(f *testing.T,\
    \ x, y int) {\n\t\t_ = Gcd(x, y)\n\t})\n}\n"
  dependsOn:
  - go-acl/splay/node.go
  - go-acl/splay/node_test.go
  - go-acl/splay/set.go
  - go-acl/splay/map.go
  - go-acl/verify/aplusb/verify.test.go
  - go-acl/verify/associative_array/verify.test.go
  - go-acl/verify/many_aplusb/verify.test.go
  - go-acl/util/unionfind.go
  - go-acl/util/sieve.go
  - go-acl/util/lib_test.go
  - go-acl/util/lib.go
  - go-acl/util/segmentTree.go
  - go-acl/util/math.go
  - go-acl/util/sieve_test.go
  - go-acl/util/monoid.go
  - go-acl/main.go
  isVerificationFile: false
  path: go-acl/util/math_test.go
  requiredBy:
  - go-acl/splay/node.go
  - go-acl/splay/node_test.go
  - go-acl/splay/set.go
  - go-acl/splay/map.go
  - go-acl/util/unionfind.go
  - go-acl/util/sieve.go
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
documentation_of: go-acl/util/math_test.go
layout: document
redirect_from:
- /library/go-acl/util/math_test.go
- /library/go-acl/util/math_test.go.html
title: go-acl/util/math_test.go
---
