package main

import (
	"./kebab" // my package カレントディレクトリから相対パス
	"fmt"     // 基本的な機能が入っている
	"strconv" // キャストに関するパッケージ
)

func main() {
	var text1 string       // 宣言と初期化
	text1 = kebab.Grill()  // glovalパッケージは大文字から始める
	num2 := kebab.Cheese() // := で宣言と代入を同時に行うことが可能

	fmt.Println(text1)
	fmt.Println(num2)

	text2 := strconv.Itoa(num2)     // int -> stringは返り値ひとつ
	num22, _ := strconv.Atoi(text2) // int -> stringは返り値ふたつ errorを捨てる場合はunderscore

	fmt.Println(text1 + text2) // 文字列の連結
	fmt.Println(num22 + 20)
}
