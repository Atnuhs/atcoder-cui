package acl

import "cmp"

// Splaysetはキー順に並んだset型のスプレー木。Splaymap[K, struct{}]の薄いラッパー。
type Splayset[K cmp.Ordered] struct {
	m *Splaymap[K, struct{}]
}

// NewSplaysetは新しいスプレー木集合を作成
func NewSplayset[K cmp.Ordered]() *Splayset[K] {
	return &Splayset[K]{m: NewSplaymap[K, struct{}]()}
}

// Sizeは要素数を返す
func (s *Splayset[K]) Size() int { return s.m.Size() }

// IsEmptyは空かどうかを返す
func (s *Splayset[K]) IsEmpty() bool { return s.m.IsEmpty() }

// Containsはkeyを含むかを返す
func (s *Splayset[K]) Contains(key K) bool {
	_, found := s.m.Get(key)
	return found
}

// Addはkeyを追加する。既に存在すれば何もしない。
func (s *Splayset[K]) Add(key K) {
	s.m.Set(key, struct{}{})
}

// Deleteはkeyを削除し、削除できたかを返す
func (s *Splayset[K]) Delete(key K) bool {
	return s.m.Delete(key)
}

// Minは最小キーを返す。空ならok=false。
func (s *Splayset[K]) Min() (key K, ok bool) {
	key, _, ok = s.m.Min()
	return
}

// Maxは最大キーを返す。空ならok=false。
func (s *Splayset[K]) Max() (key K, ok bool) {
	key, _, ok = s.m.Max()
	return
}

// PopMinは最小キーを取り出して削除する。空ならok=false。
func (s *Splayset[K]) PopMin() (key K, ok bool) {
	key, _, ok = s.m.PopMin()
	return
}

// PopMaxは最大キーを取り出して削除する。空ならok=false。
func (s *Splayset[K]) PopMax() (key K, ok bool) {
	key, _, ok = s.m.PopMax()
	return
}

// Atは昇順でk番目(0-indexed)のキーを返す。範囲外ならok=false。
func (s *Splayset[K]) At(k int) (key K, ok bool) {
	key, _, ok = s.m.At(k)
	return
}

// LowerBoundはkey以上の最小要素のindexを返す (std::lower_bound 相当)。
// 存在しなければSize()を返す。要素はAt(LowerBound(key))で取得。
func (s *Splayset[K]) LowerBound(key K) int {
	return s.m.LowerBound(key)
}

// UpperBoundはkeyより大きい最小要素のindexを返す (std::upper_bound 相当)。
// 存在しなければSize()を返す。要素はAt(UpperBound(key))で取得。
func (s *Splayset[K]) UpperBound(key K) int {
	return s.m.UpperBound(key)
}

// InOrderは昇順でキーを列挙して返す
func (s *Splayset[K]) InOrder() []K {
	entries := s.m.InOrder()
	if entries == nil {
		return nil
	}
	ret := make([]K, len(entries))
	for i, e := range entries {
		ret[i] = e.K
	}
	return ret
}
