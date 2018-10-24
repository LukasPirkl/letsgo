package main

import (
	"fmt"
	"strings"
)

func swap(a, b int) (c, d int) {
	c = b
	d = a
	return
}

func swapShort(a, b int) (int, int) {
	return b, a
}

func builder(name string) func(...string) string {
	return func(values ...string) string {
		return fmt.Sprintf("%s - %s", name, strings.Join(values, ", "))
	}
}

type Duck struct {
	name string
}

func (duck Duck) Quack() {
	fmt.Printf("Quack, I am %s\n", duck.name)
}

func Quack(duck Duck) {
	fmt.Printf("Quack, I am %s\n", duck.name)
}

func (duck *Duck) SetName(name string) {
	duck.name = name
}

func main() {
	b := builder("Max")
	fmt.Println(b("Power", "The third"))
	fmt.Println(b("Payne"))

	duck := Duck{name: "Donald"}
	duck.Quack()
	duck.SetName("Duffy")
	Quack(duck)
}
