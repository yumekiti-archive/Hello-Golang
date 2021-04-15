package main

import (
	"fmt"
)

func main() {

	// ここはコメントです

	/* ここも
	コメントです */

	/*	型(type)
		bool						真偽値(true or false)
		int8/int16/int32/int64		nビット整数
		uint8/uint16/uint32/uint64	nビット非負整数
		float32/float64				nビット浮動小数点数
		complex64/complex128		nビット虚数
		byte						1バイトデータ(uint8と同義)
		rune						1文字(int32と同義)
		uint						uint32 または uint64
		int							int32 または int64
		uintptr						ポインタを表現するのに十分な非負整数
		string						文字列
	*/

	/*	演算子
		x + y		加算 (文字列の連結にも利用)
		x - y		減算
		x * y		乗算
		x / y		除算
		x % y		除算の余り
		x & y		論理積(AND)
		x | y		論理和(OR)
		x ^ y		排他的論理和(XOR)
		x &^ y		x AND (NOT y)
		x << y		yビット左にシフト
		x >> y		yビット右にシフト
		x = y		xにyを代入
		x := y		xにyを代入(初期化の使用可能)
		x++			x = x + 1 と同義
		x--			x = x - 1 と同義
		x += y		x = x + y と同義
		x -= y		x = x - y と同義
		x *= y		x = x * y と同義
		x /= y		x = x / y と同義
		x %= y		x = x % y と同義
		x &= y		x = x & y と同義
		x |= y		x = x | y と同義
		x ^= y		x = x ^ y と同義
		x &^= y		x = x &^ y と同義
		x <<= y		x = x << y と同義
		x >>= y		x = x >> y と同義
		x && y		xかつy(AND)
		x || y		xまたはy(OR)
		!x			xがtrueの場合false/falseの場合true(NOT)
		x == y		xとyが等しければ
		x != y		xとyが等しくなければ
		x < y		yがxより大きければ
		x < y		yがx以上であれば
		x > y		yがxより小さければ
		x >= y		yがx以下であれば
		ch <- x		chチャネルにxを送信
		x =<- ch	chチャネルからxに受信
	*/

	/*	エスケープシーケンス
		\a			ベル(U+0007)
		\b			バックスペース(U+0008)
		\t			タブ(U+0009)
		\n			改行(U+000A)
		\v			垂直タブ(U+000B)
		\f			フォームフィード(U+000C)
		\r			キャリッジリターン(U+000D)
		\"			ダブルクォート(U+0022)
		\'			シングルクォート(U+0027)
		\\			バックスラッシュ(U+005C)
		\x42		ASCII文字(U+0000～U+00FF)
		\u30A2		Unicode(U+0000～U+FFFF)
		\U0001F604	Unicode(U+0000～U+10FFFF)
	*/

	/*
		変数(var)
	*/

	var a int = 123

	var b = 456

	c := 789

	var (
		d int = 123
		e int = 456
	)

	var f int
	f = 123

	var g int
	var h string
	g, h = 456, "hoge"

	var i int
	var j string
	i = 123
	j = "piyo"

	fmt.Printf("変数 %d %d %d %d %d %d %d %s %d %s \n", a, b, c, d, e, f, g, h, i, j)

	/*
		定数(const)
	*/

	const foo = 100

	const (
		foo2 = 100
		baa  = 200
	)

	fmt.Printf("定数 %d %d %d\n", foo, foo2, baa)

	/*
		配列(array)
	*/

	a1 := [3]string{}
	a1[0] = "Red"
	a1[1] = "Green"
	a1[2] = "Blue"

	b1 := [...]string{"hoge", "piyo", "huga"}

	fmt.Println(a1[0], a1[1], a1[2])
	fmt.Println(b1[0], b1[1], b1[2])

	/*
		スライス(slice)
	*/

	c1 := []string{} // スライス。個数不定
	c1 = append(c1, "Red")
	c1 = append(c1, "Green")
	c1 = append(c1, "Blue")

	fmt.Println(c1[0], c1[1], c1[2])

	d1 := []int{}

	//len() は配列やスライスの長さ(length)、cap() は容量(capacity)を求めます。
	for i := 0; i < 10; i++ {
		d1 = append(d1, i)
		fmt.Print(d1[i]+1, "-", cap(d1), " ")
	}

	fmt.Println("配列の長さ：", len(d1))

	//make(スライス型, 初期個数, 初期容量) を用いたメモリの確保ができます。
	bufa := make([]byte, 0, 1024)
	fmt.Print("make?", bufa, "\n")

	/*
		マップ(map)
	*/

	//「map[キーの型]値の型」 を用いて連想配列のようなマップを使用することができます。

	// マップを定義する
	e1 := map[string]int{
		"x": 100,
		"y": 200, // 改行する場合はカンマ必須
	}

	// マップを参照する
	fmt.Println(e1["x"])

	// マップに要素を加える
	e1["z"] = 300

	// マップの要素を削除する
	delete(e1, "z")

	// マップの長さを求める
	fmt.Println(len(e1))

	// マップに要素が存在するかどうかを調べる
	_, ok := e1["z"]
	if ok {
		fmt.Println("Exist")
	} else {
		fmt.Println("Not exist")
	}

	// マップをループ処理する
	for key, value := range e1 {
		fmt.Printf("%s=%d\n", key, value)
	}

	/*
		最後
	*/

	fmt.Printf("end\n")

}
