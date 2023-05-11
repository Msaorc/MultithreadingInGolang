package main

func main() {
	ch := make(chan string, 2)
	ch <- "Item 1"
	ch <- "Item 2"

	println(<-ch)
	println(<-ch)
}
