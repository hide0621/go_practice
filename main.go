// goroutine

//goroutineは、軽量のスレッドである。
//goステートメントで関数を指定することで、並行実行される。
//参考：https://qiita.com/taigamikami/items/fc798cdd6a4eaf9a7d5e

//チャネル

//チャネルは、並行実行されるgoroutine間を接続するパイプ(トンネル)のイメージ。
//つまり、並行実行している関数から値を受信する。(あるgoroutineから別のgoroutineへ値を渡す。)
//参考：https://qiita.com/taigamikami/items/fc798cdd6a4eaf9a7d5e

package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
func checkAPI(api string) {
	_, err := http.Get(api)
	if err != nil {
		fmt.Printf("ERROR: %s is down!\n", api)
		//関数の実行フローを抜ける
		return
	}
	fmt.Printf("SUCCESS: %s is up and running!\n", api)
}
*/

//リファクタリング
func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		//データをチャネルに送信
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		//関数の実行フローを抜ける
		return
	}
	//データをチャネルに送信
	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func main() {

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	/*
		for _, api := range apis {
			_, err := http.Get(api)
			if err != nil {
				fmt.Printf("ERROR: %s is down!\n", api)
				//for文を抜ける
				continue
			}

			fmt.Printf("SUCCESS: %s is up and running!\n", api)
		}
	*/

	/*
		for _, api := range apis {
			//checkAPI関数は完了する時間が無かっただけで実行はされている
			go checkAPI(api)
		}
	*/

	//リファクタリング
	ch := make(chan string)

	for _, api := range apis {
		//checkAPI関数は完了する時間が無かっただけで実行はされている
		go checkAPI(api, ch)
	}
	//チャネルの作成により必要が無くなった
	//time.Sleep(3 * time.Second)

	elapsed := time.Since(start)

	fmt.Println(<-ch)

	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
