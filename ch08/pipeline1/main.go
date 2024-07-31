package main

import (
	"flag"
	"fmt"
)

var limit = flag.Int("limit", 10, "limit to get the squares values")

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x <= *limit; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}
