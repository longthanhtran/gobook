package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 4)

	messages <- "buffered"
	messages <- "channel"
	messages <- "one more"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
