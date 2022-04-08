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
//バッファが空の時はチャネルからの受信をブロックする

package main

import "fmt"

func send(ch chan string, message string) {
	ch <- message
}

func main() {

	size := 4
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	send(ch, "three")
	send(ch, "four")
	//fmt.Println(ch) //チャネルのアドレス
	fmt.Println("All data sent to the channel ...")

	for i := 0; i < size; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done!")

	/*
		size := 2
		ch := make(chan string, size)
		send(ch, "one")
		send(ch, "two")
		send(ch, "three")
		send(ch, "four")
		send(ch, "All data sent to the channel ...")

		for i := 0; i < 4; i++ {
			fmt.Println(<-ch)
		}

		fmt.Println("Done!")
	*/

}
