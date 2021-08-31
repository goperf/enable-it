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
	f, err := os.Create("cpu.profile")
	if err != nil {
		panic(err)
	}

	b, err := os.Create("block.profile")
	if err != nil {
		panic(err)
	}

	m, err := os.Create("memory.profile")
	if err != nil {
		panic(err)
	}

	runtime.MemProfileRate = 1
	runtime.SetBlockProfileRate(1)

	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}

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

	// Collect CPU profile
	pprof.StopCPUProfile()

	// collect block profile
	pprof.Lookup("block").WriteTo(b, 0)

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
