package main

import (
	"fmt"
	"time"
)

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	x := make([]float64, 5, 10)
	fmt.Println(x)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
