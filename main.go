package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//json

//JSONから構造体へ変換

type A struct{}

type User struct {
	//JSON形式で右記の小文字の方に変換
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A,omitempty"` //構造体はポインタにするとomitemptyで消せる
}

func main() {

	//newで初期化.uはポインタ型
	u := new(User) // var u *User
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	//Marshal JSONに変換してbyteのスライスで受け取る
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	//-----------------------------------------

	fmt.Printf("%T\n", bs)

	//newで初期化.u2はポインタ型
	u2 := new(User) //var u2 *User

	//Unmarshal JSONを構造体に変換
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println("err", err)
	}

	fmt.Println(u2)

}
