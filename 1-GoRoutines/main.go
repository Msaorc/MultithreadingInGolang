package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go task("Marcao")
	go task("Duda")
	go func() {
		for i := 1; i <= 4; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(2 * time.Second)
		}
	}()

	// É necessario que a thread principal (main) esteja em execução para que as demais continuem em execução.
	time.Sleep(15 * time.Second)
}
