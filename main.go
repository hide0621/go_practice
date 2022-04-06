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
	var resp GitHubResponse
	//JSONを構造体のデータに変換して第二引数に格納、第二引数の構造体は構造体のアドレス
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
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
