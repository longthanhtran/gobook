package main

import (
	"fmt"
)

func main() {
	defer func() {
		str := recover()
		fmt.Print(str)
	}()
	panic("PANIC")
}
