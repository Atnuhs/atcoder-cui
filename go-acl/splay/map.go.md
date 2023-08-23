---
data:
  _extendedDependsOn:
  - icon: ':heavy_check_mark:'
    path: go-acl/main.go
    title: go-acl/main.go
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
    RuntimeError: bundler is not specified: go-acl/splay/map.go\n"
  code: "package splay\n\ntype SplayMap struct {\n\troot *SplayNode\n}\n\nfunc NewSplayMap()\
    \ *SplayMap {\n\treturn &SplayMap{\n\t\troot: nil,\n\t}\n}\n\nfunc NewSplayMapNode(key\
    \ int, value int) *SplayNode {\n\treturn NewSplayNode(key, value)\n}\n\nfunc (ss\
    \ *SplayMap) Push(key int, value int) {\n\tnode := NewSplayMapNode(key, value)\n\
    \tif ss.root == nil {\n\t\tss.root = node\n\t}\n\tss.root = ss.root.Insert(node)\n\
    }\n\nfunc (ss *SplayMap) Remove(key int) int {\n\tvar removed *SplayNode\n\tss.root,\
    \ removed = ss.root.Delete(NewSplayMapNode(key, -1))\n\tif removed != nil {\n\t\
    \treturn removed.value\n\t}\n\treturn 0\n}\n\nfunc (ss *SplayMap) Has(key int)\
    \ bool {\n\treturn ss.root.Has(key)\n}\n\nfunc (ss *SplayMap) Values() (arr []int)\
    \ {\n\treturn ss.root.values()\n}\n\nfunc (ss *SplayMap) Size() int {\n\treturn\
    \ ss.root.size\n}\n\nfunc (ss *SplayMap) IsEmpty() bool {\n\treturn ss.root ==\
    \ nil\n}\n\nfunc (ss *SplayMap) At(key int) int {\n\tfound := ss.root.FindAndSplay(key)\n\
    \tif found == nil {\n\t\treturn 0\n\t}\n\tss.root = found\n\treturn ss.root.value\n\
    }\n\nfunc (ss *SplayMap) String() string {\n\tif ss.root == nil {\n\t\treturn\
    \ \"\"\n\t}\n\treturn ss.root.String()\n}\n\nfunc (ss *SplayMap) Ge(value int)\
    \ int {\n\tidx := ss.root.Ge(value)\n\tss.root = ss.root.FindAt(idx)\n\treturn\
    \ ss.root.key\n}\n\nfunc (ss *SplayMap) Gt(value int) int {\n\treturn ss.Ge(value\
    \ + 1)\n}\n\nfunc (ss *SplayMap) Le(value int) int {\n\treturn ss.Ge(value+1)\
    \ - 1\n}\n\nfunc (ss *SplayMap) Lt(value int) int {\n\treturn ss.Ge(value) - 1\n\
    }\n"
  dependsOn:
  - go-acl/splay/node.go
  - go-acl/splay/node_test.go
  - go-acl/splay/set.go
  - go-acl/verify/aplusb/verify.test.go
  - go-acl/verify/associative_array/verify.test.go
  - go-acl/verify/many_aplusb/verify.test.go
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
  path: go-acl/splay/map.go
  requiredBy:
  - go-acl/splay/node.go
  - go-acl/splay/node_test.go
  - go-acl/splay/set.go
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
  timestamp: '2023-08-24 01:49:17+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - go-acl/verify/aplusb/verify.test.go
  - go-acl/verify/associative_array/verify.test.go
  - go-acl/verify/many_aplusb/verify.test.go
documentation_of: go-acl/splay/map.go
layout: document
redirect_from:
- /library/go-acl/splay/map.go
- /library/go-acl/splay/map.go.html
title: go-acl/splay/map.go
---
