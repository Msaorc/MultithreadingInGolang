package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go processOrder(ch)
	readOrder(ch, &wg)
	wg.Wait()
}

func processOrder(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func readOrder(ch chan int, wg *sync.WaitGroup) {
	for orderNumber := range ch {
		fmt.Printf("Order number: %d has been successfully processed\n", orderNumber)
		wg.Done()
	}
}
