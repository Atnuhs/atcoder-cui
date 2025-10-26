package main

import "fmt"

// SplayTreeの実用的な使用例

func ExampleSplayTreeBasicUsage() {
	fmt.Println("=== SplayTree基本使用例 ===")

	// 文字列キー、整数値のマップ
	tree := NewSplaymap[string, int]()

	// データ挿入
	tree.Insert("apple", 100)
	tree.Insert("banana", 200)
	tree.Insert("cherry", 300)
	tree.Insert("date", 400)

	fmt.Printf("サイズ: %d\n", tree.Size())

	// 検索
	if value, found := tree.Has("banana"); found {
		fmt.Printf("banana: %d\n", value)
	}

	// 存在しないキーの検索
	if _, found := tree.Has("orange"); !found {
		fmt.Println("orange: 見つかりません")
	}

	// 値の更新
	tree.Insert("banana", 250) // 既存キーなので値を更新
	if value, found := tree.Has("banana"); found {
		fmt.Printf("banana（更新後）: %d\n", value)
	}

	// 削除
	if tree.Delete("cherry") {
		fmt.Println("cherry を削除しました")
	}
	fmt.Printf("削除後のサイズ: %d\n", tree.Size())

	// 全要素の取得（ソート順）
	fmt.Println("全要素（アルファベット順）:")
	for _, item := range tree.InOrder() {
		fmt.Printf("  %s: %d\n", item.K, item.V)
	}
}

func ExampleSplayTreeWithStructs() {
	fmt.Println("\n=== 構造体を値とする例 ===")

	type Person struct {
		Name string
		Age  int
		City string
	}

	tree := NewSplaymap[int, Person]()

	// 人物データを挿入
	people := []Person{
		{"Alice", 25, "Tokyo"},
		{"Bob", 30, "Osaka"},
		{"Charlie", 28, "Kyoto"},
		{"Diana", 22, "Fukuoka"},
		{"Eve", 35, "Sapporo"},
	}
	ids := []int{1001, 1003, 1002, 1005, 1004}
	for i, person := range people {
		tree.Insert(ids[i], person)
	}

	// ID検索
	if person, found := tree.Has(1002); found {
		fmt.Printf("ID 1002: %s (%d歳, %s)\n", person.Name, person.Age, person.City)
	}

	// 全員をID順で表示
	fmt.Println("全員（ID順）:")
	for _, item := range tree.InOrder() {
		p := item.V
		fmt.Printf("  ID %d: %s (%d歳, %s)\n", item.K, p.Name, p.Age, p.City)
	}

	// 特定の条件でフィルタリング
	fmt.Println("30歳以上:")
	for _, item := range tree.InOrder() {
		p := item.V
		if p.Age >= 30 {
			fmt.Printf("  %s (%d歳)\n", p.Name, p.Age)
		}
	}
}

func ExampleSplayTreeFloatKeys() {
	fmt.Println("\n=== float64キーの例 ===")

	tree := NewSplaymap[float64, string]()

	// 温度データ
	tree.Insert(36.5, "平熱")
	tree.Insert(37.5, "微熱")
	tree.Insert(38.0, "発熱")
	tree.Insert(35.5, "低体温")
	tree.Insert(39.0, "高熱")

	// 特定の温度
	if status, found := tree.Has(37.5); found {
		fmt.Printf("37.5℃: %s\n", status)
	}

	// 温度順で表示
	fmt.Println("体温ステータス（温度順）:")
	for _, item := range tree.InOrder() {
		fmt.Printf("  %.1f℃: %s\n", item.K, item.V)
	}
}

func ExampleSplayTreeCompetitiveProgramming() {
	fmt.Println("\n=== 競技プログラミング向けの例 ===")

	// 座標管理
	type Point struct {
		X, Y int
	}

	tree := NewSplaymap[int, Point]()

	// 点を挿入（x座標をキーとする）
	points := []Point{{1, 2}, {3, 4}, {2, 1}, {5, 3}, {4, 2}}
	for _, p := range points {
		tree.Insert(p.X, p) // x座標をキーとして使用
	}

	fmt.Printf("点の数: %d\n", tree.Size())

	// x座標順で点を表示
	fmt.Println("点（x座標順）:")
	for _, item := range tree.InOrder() {
		p := item.V
		fmt.Printf("  (%d, %d)\n", p.X, p.Y)
	}

	// 特定のx座標の点を検索
	if point, found := tree.Has(3); found {
		fmt.Printf("x=3の点: (%d, %d)\n", point.X, point.Y)
	}

	// 範囲検索のシミュレーション（x座標が2以上4以下）
	fmt.Println("x座標が2-4の範囲の点:")
	for _, item := range tree.InOrder() {
		if item.K >= 2 && item.K <= 4 {
			p := item.V
			fmt.Printf("  (%d, %d)\n", p.X, p.Y)
		}
	}
}

func ExampleSplayTreePerformance() {
	fmt.Println("\n=== パフォーマンステスト ===")

	tree := NewSplaymap[int, int]()

	// 大量データの挿入
	const N = 1000
	for i := 0; i < N; i++ {
		tree.Insert(i, i*i) // 値は平方数
	}

	fmt.Printf("挿入完了: %d個の要素\n", tree.Size())

	// いくつかの値を検索
	searchKeys := []int{100, 500, 999, 1000} // 1000は存在しない
	for _, key := range searchKeys {
		if value, found := tree.Has(key); found {
			fmt.Printf("key=%d, value=%d\n", key, value)
		} else {
			fmt.Printf("key=%d: 見つかりません\n", key)
		}
	}

	// いくつかの値を削除
	deleteCount := 0
	for i := 0; i < N; i += 2 { // 偶数を削除
		if tree.Delete(i) {
			deleteCount++
		}
	}

	fmt.Printf("削除完了: %d個削除, 残り%d個\n", deleteCount, tree.Size())
}

// 全ての例を実行
func RunSplayTreeExamples() {
	ExampleSplayTreeBasicUsage()
	ExampleSplayTreeWithStructs()
	ExampleSplayTreeFloatKeys()
	ExampleSplayTreeCompetitiveProgramming()
	ExampleSplayTreePerformance()
}
