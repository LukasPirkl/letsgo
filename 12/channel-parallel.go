package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)

	doneSending := &sync.WaitGroup{}
	doneReceiving := &sync.WaitGroup{}

	makeSender("A", ch, doneSending)
	//makeSender("B", ch, doneSending)
	//makeSender("C", ch, doneSending)
	makeReceiver("X", ch, doneReceiving)
	//makeReceiver("Y", ch, doneReceiving)
	//makeReceiver("Z", ch, doneReceiving)

	doneSending.Wait()
	fmt.Println("-- DONE SENDING --")
	close(ch)
	doneReceiving.Wait()
	fmt.Println("-- DONE RECEIVING --")
}

func makeSender(name string, ch chan<- string, wg *sync.WaitGroup) {
	wg.Add(1)
	go send(name, ch, wg)
}

func send(name string, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("data no.%d", i)
		ch <- value
		fmt.Printf("Sent by %s: %s\n", name, value)
	}
	wg.Done()
}

func makeReceiver(name string, ch <-chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	go receive(name, ch, wg)
}

func receive(name string, ch <-chan string, wg *sync.WaitGroup) {
	for value := range ch {
		fmt.Printf("Processing by %s started \t%s\n", name, value)
		time.Sleep(time.Second)
		fmt.Printf("Processing by %s ended \t%s\n", name, value)
	}
	wg.Done()
}
