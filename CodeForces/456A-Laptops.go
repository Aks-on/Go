package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//https://codeforces.com/problemset/problem/456/A
	var n int
	fmt.Scan(&n)
	prises := make([]int, n)
	quality := make([]int, n)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &prises[i], &quality[i])
	}

	check := false
	for i := 0; i < n; i++ {
		for j := 1; i+j < n; j++ {
			if (prises[i] < prises[i+j] && quality[i] > quality[i+j]) || (prises[i] > prises[i+j] && quality[i] < quality[i+j]) {
				check = true
			}
		}
	}
	if check {
		fmt.Println("Happy Alex")
	} else {
		fmt.Println("Poor Alex")
	}
}
