package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	ch := make(chan string)

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- "ping"
			fmt.Println(<-ch)
		}
		close(ch)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for value := range ch {
			fmt.Println(value)
			ch <- "pong"
		}
		wg.Done()
	}()

	wg.Wait()
}
