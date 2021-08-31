package main

import (
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
	"sync"
)

var num int
var match bool
var r = regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]+$`)

func main() {

	b, err := os.Create("block.profile")
	if err != nil {
		panic(err)
	}

	runtime.SetBlockProfileRate(1)

	// Application code

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {

		for i := 0; i < 1e6; i++ {
			num = Fib(10)
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 1e6; i++ {
			MatchRegex()
		}

		wg.Done()
	}()

	wg.Wait()

	// collect block profile
	pprof.Lookup("block").WriteTo(b, 0)

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
