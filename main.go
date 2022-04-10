// goroutine

//goroutineは、軽量のスレッドである。
//goステートメントで関数を指定することで、並行処理される。
//参考：https://qiita.com/taigamikami/items/fc798cdd6a4eaf9a7d5e

//チャネル

//チャネルは、並行処理されるgoroutine間を接続するパイプ(トンネル)のイメージ。
//つまり、並行実行している関数から値を受信する。(あるgoroutineから別のgoroutineへ値を渡す。)
//参考：https://qiita.com/taigamikami/items/fc798cdd6a4eaf9a7d5e

//チャネルのselectについて
//selectを利用することで処理対象のイベントを受信するまで、プログラムの実行はブロックさせる
//受信したものから、画面に表示される。

package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch4 := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		ch4 <- "four"
		close(ch4)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
		close(ch1)
	}()

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)

		case msg4 := <-ch4:
			fmt.Println("received", msg4)

		}
	}

}
