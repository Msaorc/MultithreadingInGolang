package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	qtdeWorkers := 100

	for i := 0; i < qtdeWorkers; i++ {
		go worker(i, data)
	}
	for i := 0; i < 1000; i++ {
		data <- i
	}
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Woker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
