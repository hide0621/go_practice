package main

import (
	"fmt"
	"gosample/geometry"
)

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	//fmt.Println("Size", t.size)
	fmt.Println("Perimeter", t.Perimeter())
}
