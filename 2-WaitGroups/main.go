package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(24)
	go task("Marcao", &waitGroup)
	go task("Duda", &waitGroup)
	go func() {
		for i := 1; i <= 4; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(2 * time.Second)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
}
