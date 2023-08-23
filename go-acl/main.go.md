---
data:
  _extendedDependsOn:
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
    RuntimeError: bundler is not specified: go-acl/main.go\n"
  code: "package main\n\nimport \"go-acl/util\"\n\nfunc main() {\n\tdefer util.Out.Flush()\n\
    }\n"
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
  - go-acl/util/math_test.go
  - go-acl/util/lib_test.go
  - go-acl/util/lib.go
  - go-acl/util/segmentTree.go
  - go-acl/util/math.go
  - go-acl/util/sieve_test.go
  - go-acl/util/monoid.go
  isVerificationFile: false
  path: go-acl/main.go
  requiredBy:
  - go-acl/splay/node.go
  - go-acl/splay/node_test.go
  - go-acl/splay/set.go
  - go-acl/splay/map.go
  - go-acl/util/unionfind.go
  - go-acl/util/sieve.go
  - go-acl/util/math_test.go
  - go-acl/util/lib_test.go
  - go-acl/util/lib.go
  - go-acl/util/segmentTree.go
  - go-acl/util/math.go
  - go-acl/util/sieve_test.go
  - go-acl/util/monoid.go
  timestamp: '2023-08-24 01:49:17+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - go-acl/verify/aplusb/verify.test.go
  - go-acl/verify/associative_array/verify.test.go
  - go-acl/verify/many_aplusb/verify.test.go
documentation_of: go-acl/main.go
layout: document
redirect_from:
- /library/go-acl/main.go
- /library/go-acl/main.go.html
title: go-acl/main.go
---
