package main

import "fmt"

func main() {

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	var x = make([]int, 5)
	fmt.Println(x, len(x), cap(x))
	if x == nil {
		fmt.Println("nil!")
	} else {
		fmt.Println("not nil!")
	}

	var names []string
	names = []string{"one", "two", "three"} // this creates array and then slice that references it
	colors := [2]string{"red", "blue"}      // this creates array
	_ = colors

	fmt.Printf("len=%d cap=%d %v\n", len(names), cap(names), names)
	names = append(names, "four")
	fmt.Printf("len=%d cap=%d %v\n", len(names), cap(names), names)

	fmt.Println("--- CLASSIC ---")
	for i := 0; i < len(names); i++ {
		fmt.Printf("%d - %s\n", i, names[i])
	}

	names = append(names, "four")

	fmt.Println("--- FOREACH ---")
	for index, value := range names[1:3] {
		fmt.Printf("%d - %s\n", index, value)
	}

	fmt.Println("--- WHILE ---")
	index := 0
	for index < 2 {
		fmt.Printf("%d - %s\n", index, names[index])
		index++
	}
}
