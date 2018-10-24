package main

import (
	"fmt"
)

func open(id int) {
	fmt.Printf("Open %d\n", id)
}

func close(id int) {
	fmt.Printf("Close %d\n", id)
}

func main() {
	id := 123456

	open(id)
	defer close(id)

	id = 654321

	fmt.Printf("Working %d\n", id)
}
