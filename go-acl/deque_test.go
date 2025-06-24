package main

import (
	"container/list"
	"math/rand"
	"testing"
	"time"

	"github.com/Atnuhs/atcoder-cui/go-acl/testlib"
)

func TestDeque_NewDeque(t *testing.T) {
	d := NewDeque[int]()

	// 初期状態の確認
	if d == nil {
		t.Fatal("NewDeque returned nil")
	}

	if len(d.buf) != DEQUE_CAP {
		t.Errorf("Expected buffer size %d, got %d", DEQUE_CAP, len(d.buf))
	}

	if d.l != 0 || d.r != 0 {
		t.Errorf("Expected l=0, r=0, got l=%d, r=%d", d.l, d.r)
	}
}

func TestDeque_PushBack_Basic(t *testing.T) {
	d := NewDeque[int]()

	// 基本的なPushBack
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	// 内部状態の確認
	if d.l != 0 {
		t.Errorf("Expected l=0, got l=%d", d.l)
	}
	if d.r != 3 {
		t.Errorf("Expected r=3, got r=%d", d.r)
	}

	// データの確認
	if d.buf[0] != 1 || d.buf[1] != 2 || d.buf[2] != 3 {
		t.Errorf("Data not stored correctly: [%d, %d, %d]", d.buf[0], d.buf[1], d.buf[2])
	}
}

func TestDeque_PushFront_Basic(t *testing.T) {
	d := NewDeque[int]()

	// PushFrontのテスト
	d.PushFront(1)

	// 最初のPushFrontで値が設定されていない問題をテスト
	t.Logf("After PushFront(1): l=%d, r=%d", d.l, d.r)
	t.Logf("Buffer around l: buf[%d]=%d", d.l, d.buf[d.l])
}

func TestDeque_PopBack_Basic(t *testing.T) {
	d := NewDeque[int]()

	// 空のdequeからのpop
	val, ok := d.PopBack()
	if ok {
		t.Errorf("Expected PopBack to fail on empty deque, but got value %d", val)
	}

	// データを追加してpop
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	val, ok = d.PopBack()
	if !ok || val != 3 {
		t.Errorf("Expected PopBack to return (3, true), got (%d, %v)", val, ok)
	}

	val, ok = d.PopBack()
	if !ok || val != 2 {
		t.Errorf("Expected PopBack to return (2, true), got (%d, %v)", val, ok)
	}
}

func TestDeque_PopFront_Basic(t *testing.T) {
	d := NewDeque[int]()

	// 空のdequeからのpop
	val, ok := d.PopFront()
	if ok {
		t.Errorf("Expected PopFront to fail on empty deque, but got value %d", val)
	}

	// データを追加してpop
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	val, ok = d.PopFront()
	if !ok || val != 1 {
		t.Errorf("Expected PopFront to return (1, true), got (%d, %v)", val, ok)
	}

	val, ok = d.PopFront()
	if !ok || val != 2 {
		t.Errorf("Expected PopFront to return (2, true), got (%d, %v)", val, ok)
	}
}

func TestDeque_Mixed_Operations(t *testing.T) {
	d := NewDeque[int]()

	// 混合操作のテスト
	d.PushBack(1)
	d.PushFront(0)
	d.PushBack(2)
	d.PushFront(-1)

	t.Logf("After mixed pushes: l=%d, r=%d", d.l, d.r)

	// 期待される順序: [-1, 0, 1, 2]
	val, ok := d.PopFront()
	if !ok {
		t.Fatal("PopFront failed unexpectedly")
	}
	t.Logf("PopFront: %d", val)

	val, ok = d.PopBack()
	if !ok {
		t.Fatal("PopBack failed unexpectedly")
	}
	t.Logf("PopBack: %d", val)
}

func TestDeque_At_Method(t *testing.T) {
	d := NewDeque[int]()

	// データを追加
	d.PushBack(10)
	d.PushBack(20)
	d.PushBack(30)

	// Atメソッドのテスト
	if d.At(0) != 10 {
		t.Errorf("Expected At(0)=10, got %d", d.At(0))
	}
	if d.At(1) != 20 {
		t.Errorf("Expected At(1)=20, got %d", d.At(1))
	}
	if d.At(2) != 30 {
		t.Errorf("Expected At(2)=30, got %d", d.At(2))
	}
}

func TestDeque_At_WithCircularBuffer(t *testing.T) {
	d := NewDeque[int]()

	// 循環バッファでのAtテスト
	d.PushFront(1) // これによりlが変わる
	d.PushBack(2)
	d.PushBack(3)

	t.Logf("Circular buffer test: l=%d, r=%d", d.l, d.r)

	// Atが正しく動作するかテスト
	val0 := d.At(0)
	val1 := d.At(1)
	val2 := d.At(2)

	t.Logf("At(0)=%d, At(1)=%d, At(2)=%d", val0, val1, val2)
}

func TestDeque_Grow_Functionality(t *testing.T) {
	d := NewDeque[int]()

	// バッファサイズまで埋める
	initialCap := len(d.buf)

	// バッファをほぼ満杯にする
	for i := 0; i < initialCap-1; i++ {
		d.PushBack(i)
	}

	t.Logf("Before grow: len(buf)=%d, l=%d, r=%d", len(d.buf), d.l, d.r)

	// もう一つ追加してgrowを発生させる
	d.PushBack(999)

	t.Logf("After grow: len(buf)=%d, l=%d, r=%d", len(d.buf), d.l, d.r)

	// バッファサイズが2倍になったかチェック
	if len(d.buf) != initialCap*2 {
		t.Errorf("Expected buffer size to double to %d, got %d", initialCap*2, len(d.buf))
	}
}

func TestDeque_LargeOperations(t *testing.T) {
	d := NewDeque[int]()

	// 大量のデータでテスト
	const N = 2000

	// 大量pushback
	for i := 0; i < N; i++ {
		d.PushBack(i)
	}

	// 先頭からpop
	for i := 0; i < N; i++ {
		val, ok := d.PopFront()
		if !ok {
			t.Fatalf("PopFront failed at iteration %d", i)
		}
		if val != i {
			t.Errorf("Expected value %d, got %d at iteration %d", i, val, i)
		}
	}

	// 空になったかチェック
	val, ok := d.PopFront()
	if ok {
		t.Errorf("Expected empty deque, but got value %d", val)
	}
}

func TestDeque_EdgeCases(t *testing.T) {
	d := NewDeque[int]()

	// Edge case: push and pop immediately
	d.PushBack(42)
	val, ok := d.PopBack()
	testlib.AclAssert(t, true, ok)
	testlib.AclAssert(t, 42, val)

	// Edge case: push front and pop back
	d.PushFront(24)
	val, ok = d.PopBack()
	testlib.AclAssert(t, true, ok)
	testlib.AclAssert(t, 24, val)

	// Edge case: multiple push front then pop front
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)

	val, ok = d.PopFront()
	testlib.AclAssert(t, true, ok)
	testlib.AclAssert(t, 3, val) // 最後にpushしたのが最初にpop
}

func TestDeque_StringType(t *testing.T) {
	d := NewDeque[string]()

	// 文字列型でのテスト
	d.PushBack("hello")
	d.PushBack("world")
	d.PushFront("!")

	val, ok := d.PopFront()
	testlib.AclAssert(t, true, ok)
	testlib.AclAssert(t, "!", val)

	val, ok = d.PopBack()
	testlib.AclAssert(t, true, ok)
	testlib.AclAssert(t, "world", val)

	val, ok = d.PopBack()
	testlib.AclAssert(t, true, ok)
	testlib.AclAssert(t, "hello", val)
}

// バグの詳細検証用テスト
func TestDeque_PushFrontBugDetection(t *testing.T) {
	d := NewDeque[int]()

	t.Log("=== PushFront Bug Detection Test ===")

	// PushFrontが値を設定していない問題の検証
	d.PushFront(100)
	t.Logf("After PushFront(100): l=%d, r=%d", d.l, d.r)
	t.Logf("buf[l] = buf[%d] = %d", d.l, d.buf[d.l])

	// PopFrontで取得した値が正しいかチェック
	val, ok := d.PopFront()
	t.Logf("PopFront result: val=%d, ok=%v", val, ok)

	if !ok {
		t.Error("PopFront should succeed")
	}
	if val != 100 {
		t.Errorf("Expected PopFront to return 100, got %d", val)
	}
}

func TestDeque_At_BoundsCheck(t *testing.T) {
	d := NewDeque[int]()

	// 空のdequeでの境界チェック
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for At(0) on empty deque")
		}
	}()
	d.At(0)
}

func TestDeque_At_NegativeIndex(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)

	// 負のインデックスでの境界チェック
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for At(-1)")
		}
	}()
	d.At(-1)
}

func TestDeque_At_TooLargeIndex(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)

	// サイズを超えるインデックスでの境界チェック
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for At(2) when size is 2")
		}
	}()
	d.At(2)
}

func TestDeque_Random(t *testing.T) {
	const steps = 1_000_000

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	dq := NewDeque[int]()
	ref := list.New()

	for step := 0; step < steps; step++ {
		op := rng.Intn(4)
		val := rng.Int()

		switch op {
		case 0:
			dq.PushFront(val)
			ref.PushFront(val)
		case 1:
			dq.PushBack(val)
			ref.PushBack(val)
		case 2:
			if ref.Len() == 0 {
				continue
			}

			v1, ok1 := dq.PopFront()
			v2 := ref.Remove(ref.Front()).(int)

			if !ok1 || v1 != v2 {
				t.Fatalf("PopFront mismatch want %v got %v", v2, v1)
			}
		case 3:
			if ref.Len() == 0 {
				continue
			}
			v1, ok1 := dq.PopBack()
			v2 := ref.Remove(ref.Back()).(int)

			if !ok1 || v1 != v2 {
				t.Fatalf("PopBack mismatch want %v got %v", v2, v1)
			}
		}

		if dq.Size() != ref.Len() {
			t.Fatalf("size mismatch dq=%d, ref=%d", dq.Size(), ref.Len())
		}
	}

	for ref.Len() > 0 {
		v1, _ := dq.PopFront()
		v2 := ref.Remove(ref.Front()).(int)
		if v1 != v2 {
			t.Fatalf("final drain mismatch: %d vs %d", v1, v2)
		}
	}
}
