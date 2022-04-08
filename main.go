// goroutine

//goroutineは、軽量のスレッドである。
//goステートメントで関数を指定することで、並行実行される。
//参考：https://qiita.com/taigamikami/items/fc798cdd6a4eaf9a7d5e

//チャネル

//チャネルは、並行実行されるgoroutine間を接続するパイプ(トンネル)のイメージ。
//つまり、並行実行している関数から値を受信する。(あるgoroutineから別のgoroutineへ値を渡す。)
//参考：https://qiita.com/taigamikami/items/fc798cdd6a4eaf9a7d5e

//チャネルの方向について

//チャネルは関数の引数にすることができる
//チャネルを関数の引数として使うと送信か受信のどちらを意図しているか指定(わかりやすく)することができる。
/*
chan<- int // it's a channel to only send data
<-chan int // it's a channel to only receive data
*/

package main

import "fmt"

//チャネルへの送信専用
func send(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

//チャネルからの受信先行
func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}

/*
//チャネルからの受信専用の関数でチャネルへのデータの送信をするとエラーになる
func read(ch <-chan string) {
    fmt.Printf("Receiving: %#v\n", <-ch)
    ch <- "Bye!"
}
*/

func main() {
	ch := make(chan string, 1)
	send(ch, "Hello World!")
	read(ch)
}
