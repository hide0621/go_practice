package main

import "testing" //テスト用のパッケージの使用宣言

//mainパッケージのテストをしたくないならtrueとする
var Debug bool = false

func TestIsOne(t *testing.T) { //名前がTestから始まるとテスト用の関数
	i := 1
	//Debugがtrueならテストをスキップさせる
	if Debug {
		t.Skip("テストをスキップさせる")
	}
	//bool値を受け取る,テストが成功したら以下のif文は実行されずテストは成功
	v := IsOne(i)

	//vがfalseならエラーが出力されてテストが失敗
	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}
