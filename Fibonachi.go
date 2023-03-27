package main

import "fmt"

func main() {
	fmt.Println(fib(6))
	//asdasdsda
}

func fib(n int) int {
	//	f := 0
	if n == 1 {
		return 1
	} else if n <= 0 {
		return 0
	}
	return fib(n-2) + fib(n-1)
}
