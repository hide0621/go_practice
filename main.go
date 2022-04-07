//参考　https://qiita.com/yNavS4Ki/items/5b7a0c7c41eb8da8f12a

package main

import (
	"fmt"
	"net/http"
)

//ハンドラ関数catの作成
func cat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "吾輩は猫である")
}

//ハンドラ関数dogの作成
func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "吾輩は犬である")
}

func main() {

	//マルチプレクサの作成
	mux := http.NewServeMux()

	// 第二引数にハンドラ関数を渡す
	mux.HandleFunc("/cat", cat)
	mux.HandleFunc("/dog", dog)
	//webサーバーの立ち上げ
	http.ListenAndServe(":8080", mux)
}
