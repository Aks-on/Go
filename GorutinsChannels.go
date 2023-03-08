package main

import (
	"fmt"
	"time"
)

func pinger(c chan int) {
	for i := 0; ; i++ {
		c <- i
	}
}

func ponger(c chan int) {
	for i := 0; ; i++ {
		c <- -1
	}

}

func printer(c chan int) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan int)
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
