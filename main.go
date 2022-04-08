// goroutine

//goroutineは、軽量のスレッドである。
//goステートメントで関数を指定することで、並行実行される。
//参考：https://qiita.com/taigamikami/items/fc798cdd6a4eaf9a7d5e

//チャネル

//チャネルは、並行実行されるgoroutine間を接続するパイプ(トンネル)のイメージ。
//つまり、並行実行している関数から値を受信する。(あるgoroutineから別のgoroutineへ値を渡す。)
//参考：https://qiita.com/taigamikami/items/fc798cdd6a4eaf9a7d5e

//チャネルのバッファについて

//バッファは一時的に記録する場所
//バッファリングされたチャネルは、対応する受信側がいなくても決められた量までなら 値を送信することができる
//バッファが詰まるとチャネルへの送信をブロックする
//バッファが空の時はチャネルへの受信をブロックする

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3 //デッドロックが起こるようにする

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
