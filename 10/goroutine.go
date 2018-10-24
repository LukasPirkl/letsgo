package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	msg := "hello"

	go func() {
		fmt.Println(msg)
		wg.Done()
	}()

	msg = "bye"
	wg.Wait()
}
