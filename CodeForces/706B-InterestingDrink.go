package main

import "fmt"

func main() {
	//https://codeforces.com/problemset/problem/706/B
	//Input
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
	//Solution
	res := n
	for i := 0; i < q; i++ {
		for j := 0; j < n; j++ {
			if money[i] < shops[j] {
				res--
			}
		}
		fmt.Println(res)
		res = n
	}
}
