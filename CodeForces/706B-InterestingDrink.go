package main

import "fmt"

func main() {
	//https://codeforces.com/problemset/problem/706/B
	var n, q int
	fmt.Scan(&n)
	shops := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&shops[i])
	}
	fmt.Scan(&q)
	money := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&money[i])
	}

}
