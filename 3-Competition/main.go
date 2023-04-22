package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	// CompetitionProblem()
	// CompetitionMutexSolution()
	CompetitionAtomicSolution()
}

func CompetitionProblem() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		w.Write([]byte(fmt.Sprintf("Você é o visitante %d", number)))
	})
	http.ListenAndServe(":3000", mux)
}

func CompetitionMutexSolution() {
	m := sync.Mutex{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		w.Write([]byte(fmt.Sprintf("Você é o visitante %d", number)))
	})
	http.ListenAndServe(":3000", mux)
}

func CompetitionAtomicSolution() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("Você é o visitante %d", number)))
	})
	http.ListenAndServe(":3000", mux)
}
