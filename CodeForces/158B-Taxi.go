package main

import "fmt"

func main() {
	// https://codeforces.com/problemset/problem/158/B
	var n, a, cars int
	fmt.Scan(&n)
	groups := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		groups[a]++
	}
	groups[1] -= groups[3]
	if groups[1] < 0 {
		groups[1] = 0
	}
	cars += groups[4] + groups[3] + groups[2]/2 + (groups[2]%2+groups[1])/4
	if (groups[2]%2+groups[1])%4 > 0 {
		cars++
	}
	fmt.Println(cars)
}
