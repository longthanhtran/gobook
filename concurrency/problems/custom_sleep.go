package main

import (
	"fmt"
	"time"
)

func my_sleep() {
	select {
	case m := c:
		handle(m)
	case <-time.After(5 * time.Minute):
		fmt.Println("timed out")
	}
}

func main() {

}
