//参考　https://qiita.com/yNavS4Ki/items/5b7a0c7c41eb8da8f12a

package main

import (
	"fmt"
	"net/http"
)

// cat型を定義
type cat int

// cat型のハンドラーを実装
func (s cat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "吾輩は猫である")
}

// dog型を定義
type dog int

// dog型のハンドラーを実装
func (d dog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "吾輩は犬である")
}

func main() {
	var c cat
	var d dog

	//マルチプレクサを作成
	mux := http.NewServeMux()

	// URLパスとハンドラを登録
	mux.Handle("/cat", c)
	mux.Handle("/dog", d)

	http.ListenAndServe(":8080", mux)
}
