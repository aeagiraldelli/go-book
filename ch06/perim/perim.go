package main

import (
	"fmt"
	"go-book/ch06/geometry"
)

func main() {
	perim := geometry.Path{{X: 1, Y: 1}, {X: 5, Y: 1}, {X: 5, Y: 4}, {X: 1, Y: 1}}
	fmt.Println(perim.Distance())

	p := geometry.Point{X: 1, Y: 2}
	p.ScaleBy(2.0)
	fmt.Println(p)
}
