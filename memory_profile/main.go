package main

import (
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
)

var num int
var match bool
var r = regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]+$`)

func main() {

	m, err := os.Create("memory.profile")
	if err != nil {
		panic(err)
	}

	// Set profiling rate
	runtime.MemProfileRate = 1

	// Application code

	for i := 0; i < 1e6; i++ {
		num = Fib(10)
	}

	for i := 0; i < 1e6; i++ {
		MatchRegex()
	}

	// collect memory profile
	pprof.WriteHeapProfile(m)

}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func MatchRegex() {
	match = r.MatchString("someone@pickme.lk")
}
