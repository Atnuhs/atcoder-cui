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
    RuntimeError: bundler is not specified: go-acl/splay/node.go\n"
  code: "package splay\n\nimport (\n\t\"fmt\"\n\t\"strings\"\n\n\t. \"go-acl/util\"\
    \n)\n\ntype SplayNode struct {\n\tl, r, p         *SplayNode\n\tsize         \
    \   int\n\tkey             int\n\tvalue, min, max int\n}\n\nfunc NewSplayNode(key,\
    \ value int) *SplayNode {\n\tret := &SplayNode{\n\t\tl:     nil,\n\t\tr:     nil,\n\
    \t\tp:     nil,\n\t\tkey:   key,\n\t\tvalue: value,\n\t}\n\tret.update()\n\treturn\
    \ ret\n}\n\nfunc (sn *SplayNode) index() int {\n\tif sn == nil {\n\t\treturn -1\n\
    \t}\n\tif sn.l != nil {\n\t\treturn sn.l.size\n\t}\n\treturn 0\n}\n\nfunc (sn\
    \ *SplayNode) update() {\n\tif sn == nil {\n\t\treturn\n\t}\n\tsn.size = 1\n\t\
    sn.min = sn.value\n\tsn.max = sn.value\n\n\tif sn.l != nil {\n\t\tsn.size += sn.l.size\n\
    \t\tsn.min = Min(sn.min, sn.l.min)\n\t\tsn.max = Max(sn.max, sn.l.max)\n\t}\n\t\
    if sn.r != nil {\n\t\tsn.size += sn.r.size\n\t\tsn.min = Min(sn.min, sn.r.min)\n\
    \t\tsn.max = Max(sn.max, sn.r.max)\n\t}\n}\n\nfunc (sn *SplayNode) state() int\
    \ {\n\tif sn.p == nil {\n\t\treturn 0\n\t}\n\tif sn.p.l == sn {\n\t\treturn 1\n\
    \t}\n\tif sn.p.r == sn {\n\t\treturn -1\n\t}\n\treturn INF\n}\n\nfunc (sn *SplayNode)\
    \ rotate() {\n\tif sn == nil {\n\t\treturn\n\t}\n\tns := sn.state()\n\tif ns ==\
    \ 0 {\n\t\treturn\n\t}\n\n\tp := sn.p\n\tps := p.state()\n\n\t// edge 1\n\tpp\
    \ := p.p\n\tswitch ps {\n\tcase 1:\n\t\tpp.l = sn\n\tcase -1:\n\t\tpp.r = sn\n\
    \t}\n\tsn.p = pp\n\n\t// edge 2, 3\n\tvar c *SplayNode\n\tswitch ns {\n\tcase\
    \ 1:\n\t\tc = sn.r\n\t\tsn.r = p\n\t\tp.l = c\n\tcase -1:\n\t\tc = sn.l\n\t\t\
    sn.l = p\n\t\tp.r = c\n\t}\n\n\tp.p = sn\n\tif c != nil {\n\t\tc.p = p\n\t}\n\t\
    p.update()\n\tsn.update()\n}\n\nfunc (sn *SplayNode) splay() {\n\tif sn == nil\
    \ {\n\t\treturn\n\t}\n\tfor sn.p != nil {\n\t\t// sn is not root\n\n\t\tif sn.p.state()\
    \ == 0 {\n\t\t\t// sn.p is root\n\t\t\tsn.rotate()\n\t\t\tcontinue\n\t\t}\n\n\t\
    \tif sn.state() == sn.p.state() {\n\t\t\tsn.p.rotate()\n\t\t\tsn.rotate()\n\t\t\
    } else {\n\t\t\tsn.rotate()\n\t\t\tsn.rotate()\n\t\t}\n\t}\n}\n\nfunc (sn *SplayNode)\
    \ values() []int {\n\tret := make([]int, 0)\n\tif sn == nil {\n\t\treturn ret\n\
    \t}\n\tif sn.l != nil {\n\t\tret = append(ret, sn.l.values()...)\n\t}\n\tret =\
    \ append(ret, sn.key)\n\tif sn.r != nil {\n\t\tret = append(ret, sn.r.values()...)\n\
    \t}\n\treturn ret\n}\n\nfunc (sn *SplayNode) String() string {\n\tret := strings.Builder{}\n\
    \tret.WriteString(\"(\")\n\tif sn.l != nil {\n\t\tret.WriteString(fmt.Sprintf(\"\
    %s \", sn.l.String()))\n\t}\n\tret.WriteString(fmt.Sprint(sn.key))\n\tif sn.r\
    \ != nil {\n\t\tret.WriteString(fmt.Sprintf(\" %s\", sn.r.String()))\n\t}\n\t\
    ret.WriteString(\")\")\n\treturn ret.String()\n}\n\nfunc (sn *SplayNode) describe(rank\
    \ int) string {\n\tret := \"\"\n\tif sn.r != nil {\n\t\tret += sn.r.describe(rank\
    \ + 1)\n\t}\n\tret += fmt.Sprintf(\n\t\tstrings.Repeat(\"    \", rank)+\"-[k:%d,\
    \ v:%d, sz: %d, rank: %d]\\n\",\n\t\tsn.key,\n\t\tsn.value,\n\t\tsn.size,\n\t\t\
    rank,\n\t)\n\n\tif sn.l != nil {\n\t\tret += sn.l.describe(rank + 1)\n\t}\n\t\
    return ret\n}\n\nfunc (sn *SplayNode) maxRank(rank int) int {\n\tret := rank\n\
    \tif sn.r != nil {\n\t\tret = Max(ret, sn.r.maxRank(rank+1))\n\t}\n\tif sn.l !=\
    \ nil {\n\t\tret = Max(ret, sn.l.maxRank(rank+1))\n\t}\n\treturn ret\n}\n\nfunc\
    \ (sn *SplayNode) FindAt(idx int) (found *SplayNode) {\n\tif sn == nil {\n\t\t\
    return nil\n\t}\n\tif idx < 0 || sn.size <= idx {\n\t\treturn nil\n\t}\n\t// n\
    \ include [0, n)\n\tnow := sn\n\tfor now != nil {\n\t\tswitch {\n\t\tcase idx\
    \ == now.index():\n\t\t\treturn now\n\t\tcase idx < now.index():\n\t\t\tnow =\
    \ now.l\n\t\tcase idx > now.index():\n\t\t\tidx -= now.index() + 1\n\t\t\tnow\
    \ = now.r\n\t\t}\n\t}\n\tpanic(\"must not reach this code\")\n}\n\nfunc (sn *SplayNode)\
    \ FindAtAndSplay(idx int) *SplayNode {\n\tnode := sn.FindAt(idx)\n\tnode.splay()\n\
    \treturn node\n}\n\nfunc (sn *SplayNode) Find(key int) (found *SplayNode) {\n\t\
    now := sn\n\tfor now != nil {\n\t\tif now.key == key {\n\t\t\treturn now\n\t\t\
    }\n\n\t\tif now.key > key {\n\t\t\tnow = now.l\n\t\t} else {\n\t\t\tnow = now.r\n\
    \t\t}\n\t}\n\treturn nil\n}\n\nfunc (sn *SplayNode) FindAndSplay(key int) (found\
    \ *SplayNode) {\n\tfound = sn.Find(key)\n\tfound.splay()\n\treturn found\n}\n\n\
    func (sn *SplayNode) Has(key int) bool {\n\tfound := sn.Find(key)\n\tif found\
    \ == nil {\n\t\treturn false\n\t}\n\treturn found.key == key\n}\n\nfunc (sn *SplayNode)\
    \ Ge(key int) (idx int) {\n\tif sn == nil {\n\t\treturn 0\n\t}\n\tnow := sn\n\t\
    idx = sn.size\n\ti := 0\n\tfor now != nil {\n\t\tif now.key >= key {\n\t\t\tidx\
    \ = Min(idx, i+now.index())\n\t\t\tnow = now.l\n\t\t} else {\n\t\t\ti += now.index()\
    \ + 1\n\t\t\tnow = now.r\n\t\t}\n\t}\n\treturn idx\n}\n\nfunc (sn *SplayNode)\
    \ MergeR(rroot *SplayNode) *SplayNode {\n\tif rroot == nil {\n\t\treturn sn\n\t\
    }\n\tif sn == nil {\n\t\treturn rroot\n\t}\n\tsn = sn.FindAtAndSplay(sn.size -\
    \ 1) // always found\n\tsn.r = rroot\n\trroot.p = sn\n\tsn.update()\n\treturn\
    \ sn\n}\n\nfunc (sn *SplayNode) Split(idx int) (*SplayNode, *SplayNode) {\n\t\
    if sn == nil {\n\t\treturn nil, nil\n\t}\n\tif idx == sn.size {\n\t\treturn sn,\
    \ nil\n\t}\n\n\trroot := sn.FindAtAndSplay(idx)\n\tif rroot == nil {\n\t\t// idx\
    \ is out of index\n\t\treturn nil, nil\n\t}\n\n\tlroot := rroot.l\n\tif lroot\
    \ != nil {\n\t\tlroot.p = nil\n\t}\n\trroot.l = nil\n\n\trroot.update()\n\t//\
    \ lroot not need to update()\n\treturn lroot, rroot\n}\n\nfunc (sn *SplayNode)\
    \ InsertAt(idx int, node *SplayNode) *SplayNode {\n\tlroot, rroot := sn.Split(idx)\n\
    \tif lroot == nil {\n\t\treturn node.MergeR(rroot)\n\t} else {\n\t\treturn lroot.MergeR(node).MergeR(rroot)\n\
    \t}\n}\n\nfunc (sn *SplayNode) DeleteAt(idx int) (root *SplayNode, dropped *SplayNode)\
    \ {\n\tlroot, rroot := sn.Split(idx)\n\tif rroot == nil {\n\t\treturn lroot, nil\n\
    \t}\n\tdel, rroot := rroot.Split(1)\n\tif lroot == nil {\n\t\treturn rroot, del\n\
    \t} else {\n\t\troot = lroot.MergeR(rroot)\n\t\treturn root, del\n\t}\n}\n\nfunc\
    \ (sn *SplayNode) Insert(node *SplayNode) *SplayNode {\n\tidx := sn.Ge(node.key)\n\
    \tif found := sn.FindAt(idx); found != nil {\n\t\tif found.key == node.key {\n\
    \t\t\treturn sn\n\t\t}\n\t}\n\treturn sn.InsertAt(idx, node)\n}\n\nfunc (sn *SplayNode)\
    \ Delete(node *SplayNode) (root *SplayNode, removed *SplayNode) {\n\troot = sn.FindAndSplay(node.key)\n\
    \tif root == nil {\n\t\t// target not found\n\t\treturn sn, nil\n\t}\n\tif root.key\
    \ == node.key {\n\t\t// target found\n\t\troot, removed = root.DeleteAt(root.index())\n\
    \t}\n\t// target not found\n\treturn root, removed\n}\n"
  dependsOn:
  - go-acl/splay/set_test.go
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
  path: go-acl/splay/node.go
  requiredBy:
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
documentation_of: go-acl/splay/node.go
layout: document
redirect_from:
- /library/go-acl/splay/node.go
- /library/go-acl/splay/node.go.html
title: go-acl/splay/node.go
---
