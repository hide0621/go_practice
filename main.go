package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//json

//構造体からJSONへの変換

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

}
