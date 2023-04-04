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

	summ := make([]int, n-k+1)
	for i := 0; i <= (n - k); i++ {
		for j := 0; j < k; j++ {
			summ[i] += boards[j+i]
		}
	}
	var indexMin int
	for key, val := range summ {
		if val < summ[indexMin] {
			indexMin = key
		}
	}
	fmt.Println(indexMin + 1)
}
