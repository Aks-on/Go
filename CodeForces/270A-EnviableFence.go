package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//https://codeforces.com/problemset/problem/270/A
	var t int
	fmt.Scan(&t)
	reader := bufio.NewReader(os.Stdin)
	angles := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &angles[i])
	}

	for i := 0; i < t; i++ {
		if (angles[i]/180+2)%1 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
