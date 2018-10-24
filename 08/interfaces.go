package main

import "fmt"

type Duck struct {
	name string
}

func (duck Duck) Quack() {
	fmt.Printf("Quack, I am %s\n", duck.name)
}

type RubberDucky struct{}

func (duck *RubberDucky) Quack() {
	if duck == nil {
		fmt.Println("Silence")
	} else {
		fmt.Println("Squeak")
	}
}

type Quacker interface {
	Quack()
}

func callEmAll(quackers ...Quacker) {
	for _, item := range quackers {
		item.Quack()
	}
}

func describe(i interface{}) {
	fmt.Printf("Value: %v type: %T\n", i, i)

	rubber, ok := i.(RubberDucky)
	if ok {
		fmt.Println("Yup, it is rubber ducky")
		rubber.Quack()
	}
}

func main() {
	duck := Duck{name: "Donald"}
	rubber := RubberDucky{}
	var nilRubber *RubberDucky

	callEmAll(duck, &rubber, nilRubber)

	describe(duck)
	describe(rubber)
}
