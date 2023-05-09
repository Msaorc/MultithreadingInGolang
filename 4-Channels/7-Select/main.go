package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id      int64
	Message string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from RabbitMA"}
			c1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Receive from RabbitMQ: ID: %d - %s\n", msg.Id, msg.Message)
		case msg := <-c2:
			fmt.Printf("Receive from Kafka: ID: %d - %s\n", msg.Id, msg.Message)
		case <-time.After(time.Second * 3):
			println("timeout")
		}
	}
}
