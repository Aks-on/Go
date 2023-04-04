package main

import (
	"fmt"
)

func main() {
	//https://codeforces.com/problemset/problem/706/B
	//Input
	var n, q, a int
	fmt.Scan(&n)
	shops := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		shops[a]++
	}

	for i := 0; i < 100001; i++ {
		shops[i] += shops[i-1]
	}

	fmt.Scan(&q)
	money := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&money[i])
		if money[i] > 100000 {
			money[i] = 100000
		}
	}
	for i := 0; i < q; i++ {
		fmt.Println(shops[money[i]])
	}
}
