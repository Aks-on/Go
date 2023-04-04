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
	angles := make([]float64, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &angles[i])
	}

	for i := 0; i < t; i++ {
		if int((360/(180-angles[i]))*100000)%100000 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
