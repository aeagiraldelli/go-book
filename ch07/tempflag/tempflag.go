package main

import (
	"flag"
	"fmt"
	"go-book/ch07/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 25.0, "temperatura (ºC or ºF)")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
