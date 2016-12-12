package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	x := *a
	*a = *b
	*b = x
}

func main() {
	x, y := 1, 2

	swap(&x, &y)

	fmt.Printf("Then we have x = %d, and y = %d\n", x, y)
}
