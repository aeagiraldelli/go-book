package main

import (
	"flag"
	"fmt"
)

var limit = flag.Int("limit", 10, "limit to get the squares values")

func main() {
	flag.Parse()
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x <= *limit; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
