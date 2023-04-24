package main

import "fmt"

func main() {
	ch := make(chan int)
	go sendMessage(ch)
	receiveMessage(ch)

}

func sendMessage(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func receiveMessage(ch chan int) {
	for x := range ch {
		fmt.Printf("Message number %d received with sucess.\n", x)
	}
}
