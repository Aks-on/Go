package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//https://codeforces.com/problemset/problem/456/A
	var n, prise, quality int
	fmt.Scan(&n)
	check := false
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i <= n; i++ {
		fmt.Fscan(reader, &prise, &quality)
		if prise < quality {
			check = true
			break
		}
	}
	if check {
		fmt.Println("Happy Alex")
	} else {
		fmt.Println("Poor Alex")
	}
}
