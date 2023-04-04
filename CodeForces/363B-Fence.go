package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//https://codeforces.com/problemset/problem/363/B
	var n, k int
	fmt.Scan(&n, &k)
	reader := bufio.NewReader(os.Stdin)
	boards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &boards[i])
	}

	if n == k {
		fmt.Println(1)
	} else {
		var min, index, sum int
		min = k * 100
		for i := 0; i < n; i++ {
			if i-k >= 0 {
				sum -= boards[i-k]
			}
			sum += boards[i]
			if sum < min && sum >= k && (i-k+1) >= 0 {
				min = sum
				index = i
			}
		}
		fmt.Println(index - k + 2)
	}
}
