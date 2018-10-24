package main

import (
	"fmt"
	"letsgo/06/mylib"
)

func main() {
	a := mylib.Embedded{}
	a.Two = "asd"
	a.Three = "yxc"

	fmt.Println(a)
}
