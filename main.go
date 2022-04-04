package main

import (
	"log"
	"os"
)

func main() {
	//log.SetOutput(os.Stdout)

	/*
	   log.Print("Log\n")
	   log.Println("Log2")
	   log.Printf("Log%d\n", 3)
	*/

	/*
	   log.Fatal("Log\n")
	   log.Fatalln("Log2")
	   log.Fatalf("Log%d\n", 3)
	*/

	/*
	   log.Panic("Log\n")
	   log.Panicln("Log2")
	   log.Panicf("Log%d\n", 3)
	*/

	f, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		//エラーがあればプログラムを終了する
		return
	}

	//test.logを出力先に変更
	log.SetOutput(f)

	//test.logにPrintln関数の実引数を書き込む
	log.Println("ファイルに書き込む")

}
