package main

import "fmt"

var num int

func main() {
	f := Fib(40)
	fmt.Println("Fibonacci Number: ", f)
}

// Fibonacci series
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
