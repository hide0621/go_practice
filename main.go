//参照：https://dev.classmethod.jp/articles/golang-5/

package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func main() {

	//Key:Valueで初期化。u4はポインタ型
	u4 := &User{name: "taro", age: 30} // var u4 *User
	fmt.Println("name:%s", u4.name)

	//newで初期化.u5はポインタ型
	u5 := new(User) // var u5 *User
	u5.name = "taro"
	fmt.Println("name:%s", u5.name)

}
