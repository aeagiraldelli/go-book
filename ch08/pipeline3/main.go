package main

import (
	"flag"
	"fmt"
)

var limit = flag.Int("limit", 10, "last number to get the square")

func counter(out chan<- int) {
	for x := 0; x <= *limit; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	flag.Parse()
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
