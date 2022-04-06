package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

//フィールド(FullName)がスライスになる
type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

type customWriter struct{}

func (w customWriter) Write(p []byte) (n int, err error) {
	//byteのスライスで出力
	//fmt.Println(p)
	var resp GitHubResponse
	//JSONを構造体のデータ(string型)に変換して第二引数に格納、第二引数の構造体は構造体のアドレス
	json.Unmarshal(p, &resp)
	//構造体のスライスの形で出力
	fmt.Println(resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	//構造体のスライスの要素数を出力
	fmt.Println(len(resp))
	//fmt.Println(len(p))
	return len(resp), nil //なんでもいいからreturnを返している?,ターミナルに出力されるのは26行目からのfor文の内容と30行目のみ
}

func main() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	writer := customWriter{}

	//respに格納されたwebページのbody部分を読み込んでコピー（byte型のスライス）、writer経由でWriteメソッドを呼び出してコピーを引数pに渡してWriteメソッドの結果が返る
	io.Copy(writer, resp.Body)
}
