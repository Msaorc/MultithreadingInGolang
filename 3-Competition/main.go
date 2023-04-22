package main

import (
	"fmt"
	"net/http"
)

var number uint64 = 0

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		w.Write([]byte(fmt.Sprintf("Você é o visitante %d", number)))
	})

	http.ListenAndServe(":3000", mux)
}
