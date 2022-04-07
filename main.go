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
	return fmt.Sprintf("$%.2f", d)
}

//構造体のマップ型で定義
type database map[string]dollars

//http.Handler で使用できる ServeHTTP メソッドの実装
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	//databaseのオブジェクトを生成
	db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	//第一引数のサーバーに対して第二引数のデータを出力する
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
