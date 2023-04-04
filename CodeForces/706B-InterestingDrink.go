package main

import "fmt"

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
	fmt.Scan(&q)
	money := make([]int, q)
	res := make([]int, q)
	for i := 0; i < q; i++ {
		var x int
		fmt.Scan(&money[i])
		for key, value := range shops {
			if key <= money[i] {
				x += value
			}
		}
		res[i] = x
	}
	for i := 0; i < q; i++ {
		fmt.Println(res[i])

	}
}
