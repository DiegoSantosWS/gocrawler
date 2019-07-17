package main

import (
	"fmt"
	"time"
)

func main() {
	var c1 = make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 10)
		c1 <- "Hello "
	}()

	go func() {
		time.Sleep(time.Millisecond * 10)
		c1 <- " world"
	}()

	fmt.Println(<-c1)
	fmt.Println(<-c1)
}
