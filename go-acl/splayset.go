package main

//
// type Splayset[K Ordered] struct {
// 	*Splaymap[K, struct{}]
// }
//
// func NewSplayset[K Ordered]() *Splayset[K] {
// 	return &Splayset[K]{
// 		NewSplaymap[K, struct{}](),
// 	}
// }
//
// func (ss *Splayset[K]) Has(key K) bool {
// 	_, found := ss.Splaymap.Has(key)
// 	return found
// }
// func (ss *Splayset[K]) Insert(key K) {
// 	ss.Splaymap.Insert(key, struct{}{})
// }
//
// // Kthは[0, st.Size())の範囲内でk番目の要素を返す。
// func (ss *Splayset[K]) Kth(k int) (kk K, ok bool) {
// 	kk, _, ok = ss.Splaymap.Kth(k)
// 	return
// }
//
// func (ss *Splayset[K]) Le(key K) (k K, ok bool) {
// 	return ss.Kth(ss.Splaymap.LeAt(key))
// }
//
// func (ss *Splayset[K]) Lt(key K) (k K, ok bool) {
// 	return ss.Kth(ss.Splaymap.LtAt(key))
// }
//
// func (ss *Splayset[K]) Ge(key K) (k K, ok bool) {
// 	return ss.Kth(ss.Splaymap.GeAt(key))
// }
//
// func (ss *Splayset[K]) Gt(key K) (k K, ok bool) {
// 	return ss.Kth(ss.Splaymap.GtAt(key))
// }
//
// // InOrder は中順巡回でキーと値のペアを返す
// func (ss *Splayset[K]) InOrder() []K {
// 	if ss.root == nil {
// 		return nil
// 	}
// 	n := ss.root.size
// 	ret := make([]K, 0, n)
// 	deq := NewDeque[*splaynode[K, struct{}]]()
//
// 	cur := ss.root
// 	for {
// 		if cur != nil {
// 			deq.PushBack(cur)
// 			cur = cur.l
// 			continue
// 		}
// 		if deq.Size() == 0 {
// 			break
// 		}
// 		cur, _ = deq.PopBack()
// 		ret = append(ret, cur.key)
// 		cur = cur.r
// 	}
// 	return ret
// }
