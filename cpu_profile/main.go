package main

import (
	"os"
	"runtime/pprof"
)

var num int

func main() {
	f, err := os.Create("cpu.profile")
	if err != nil {
		panic(err)
	}

	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}

	// Application code
	for i := 0; i < 1e6; i++ {
		num = Fib(10)
	}

	// Collect CPU profile
	pprof.StopCPUProfile()

}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
