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
    path: go-acl/testlib/assert.go
    title: go-acl/testlib/assert.go
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
    path: go-acl/testlib/assert.go
    title: go-acl/testlib/assert.go
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
    RuntimeError: bundler is not specified: go-acl/splay/set_test.go\n"
  code: "package splay\n\nimport (\n\t\"testing\"\n\n\t\"go-acl/testlib\"\n)\n\nfunc\
    \ TestSplaySet_Push(t *testing.T) {\n\ttests := map[string]struct {\n\t\tpushValues\
    \ []int\n\t\twant       []int\n\t}{\n\t\t\"single value\": {\n\t\t\tpushValues:\
    \ []int{10},\n\t\t\twant:       []int{10},\n\t\t},\n\n\t\t\"sorted values\": {\n\
    \t\t\tpushValues: []int{10, 20, 30},\n\t\t\twant:       []int{10, 20, 30},\n\t\
    \t},\n\t\t\"unsorted values\": {\n\t\t\tpushValues: []int{30, 10, 20},\n\t\t\t\
    want:       []int{10, 20, 30},\n\t\t},\n\t\t\"unsorted minus values\": {\n\t\t\
    \tpushValues: []int{-30, -10, -20},\n\t\t\twant:       []int{-30, -20, -10},\n\
    \t\t},\n\t}\n\n\tfor name, tc := range tests {\n\t\tt.Run(name, func(t *testing.T)\
    \ {\n\t\t\ts := NewSplayNodeFromSlice(tc.pushValues)\n\n\t\t\tgot := s.Values()\n\
    \t\t\ttestlib.AclAssert(t, tc.want, got)\n\t\t})\n\t}\n}\n\nfunc TestSplaySet_Remove(t\
    \ *testing.T) {\n\ttests := map[string]struct {\n\t\tpushValues   []int\n\t\t\
    deleteValues []int\n\t\twant         []int\n\t}{\n\t\t\"single push single delete\"\
    : {\n\t\t\tpushValues:   []int{10},\n\t\t\tdeleteValues: []int{10},\n\t\t\twant:\
    \         []int{},\n\t\t},\n\t\t\"multi push single delete\": {\n\t\t\tpushValues:\
    \   []int{20, 10, 30},\n\t\t\tdeleteValues: []int{10},\n\t\t\twant:         []int{20,\
    \ 30},\n\t\t},\n\t\t\"multi push multi delete\": {\n\t\t\tpushValues:   []int{20,\
    \ 10, 30},\n\t\t\tdeleteValues: []int{30, 20},\n\t\t\twant:         []int{10},\n\
    \t\t},\n\t}\n\n\tfor name, tc := range tests {\n\t\tt.Run(name, func(t *testing.T)\
    \ {\n\t\t\ts := NewSplayNodeFromSlice(tc.pushValues)\n\n\t\t\tfor _, v := range\
    \ tc.deleteValues {\n\t\t\t\ts.Remove(v)\n\t\t\t}\n\n\t\t\tgot := s.Values()\n\
    \t\t\ttestlib.AclAssert(t, tc.want, got)\n\t\t})\n\t}\n}\n"
  dependsOn:
  - go-acl/splay/node.go
  - go-acl/splay/node_test.go
  - go-acl/splay/set.go
  - go-acl/splay/map.go
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
  - go-acl/util/math.go
  - go-acl/util/sieve_test.go
  - go-acl/util/monoid.go
  - go-acl/main.go
  isVerificationFile: false
  path: go-acl/splay/set_test.go
  requiredBy:
  - go-acl/splay/node.go
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
  - go-acl/util/math.go
  - go-acl/util/sieve_test.go
  - go-acl/util/monoid.go
  - go-acl/main.go
  timestamp: '2023-08-25 01:19:20+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - go-acl/verify/aplusb/verify.test.go
  - go-acl/verify/associative_array/verify.test.go
  - go-acl/verify/many_aplusb/verify.test.go
documentation_of: go-acl/splay/set_test.go
layout: document
redirect_from:
- /library/go-acl/splay/set_test.go
- /library/go-acl/splay/set_test.go.html
title: go-acl/splay/set_test.go
---
