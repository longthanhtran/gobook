package main

import (
	"fmt"
)

func main() {
	value, status := half(1)

	fmt.Printf("(%d, %t)\n", value, status)
}

func half(x uint) (result uint, err bool) {
	result, err = x/2, true

	if result == 0 {
		err = false
	}

	return
}
