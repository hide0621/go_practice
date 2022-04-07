//学習サイト　https://docs.microsoft.com/ja-jp/learn/modules/go-methods-interfaces/2-interfaces
//参考サイト　https://qiita.com/yNavS4Ki/items/5b7a0c7c41eb8da8f12a
//Sprintfについて　https://qiita.com/Sekky0905/items/5a65602dce83551184b3
//Fprint系列やSprint系列、Print系列について　https://leben.mobi/go/fmt-print-and-format/go-programming/
//指定子はこのサイトがわかりやすい　https://leben.mobi/go/fmt-print-and-format/go-programming/

package main

import (
	"fmt"
	"log"
	"net/http"
)

//独自型の定義
type dollars float32

//独自型に紐づけられたStringメソッドを実装
func (d dollars) String() string {
	//fmt.Sprintf は、書式指定子に沿ってフォーマットした文字列を標準出力でななく戻り値で返してくれる
	//dollarsという独自型であり且つ小数点が第2位までの文字列という形で戻り値を返す
	return fmt.Sprintf("$%.2f", d)
}

//構造体のマップ型で定義
//ハンドラ関数でなくハンドラーメソッドでコードを書くので独自型で定義
type database map[string]dollars

//http.Handler で使用できる ServeHTTP メソッドの実装
//database型のハンドラー
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		//fmt.Fprint系列は任意の出力先へ表示してくれる
		//string型に変換したdbオブジェクトを引数wに格納、出力先w（localhost:8000）にそれを表示
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	//databaseのオブジェクトを生成
	db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	//第一引数のサーバーに対して第二引数のデータを出力するようにハンドラーに指示
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
