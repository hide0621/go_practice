package main

import (
	"fmt"
)

/*
type Stringer interface {
	String() string
}
*/

type Person struct {
	Name, Country string
}

/*
インターフェースで定義された関数をメソッドとして定義すると
暗黙的にインターフェースを実装することになる
*/

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

func main() {
	rs := Person{"John Doe", "USA"}
	ab := Person{"Mark Collins", "United Kingdom"}
	//String()メソッドが優先される
	fmt.Printf("%s\n%s\n", rs, ab)
}
