package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {

	// Start profiling server in a separate goroutine so it does not block
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	for {
		Fib(10)
	}

}

func Fib(num int) int {
	if num < 2 {
		return 0
	}

	return Fib(num-2) + Fib(num-1)
}
