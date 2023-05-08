package main

import "fmt"

func main() {
	hello := make(chan string)
	go receive("HELLO", hello)
	send(hello)
}

// receive-only
func receive(name string, ch chan<- string) {
	ch <- name
}

// send-only
func send(ch <-chan string) {
	fmt.Println(<-ch)
}
