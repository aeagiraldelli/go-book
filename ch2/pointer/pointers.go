package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	value := 2
	fmt.Println(value)
	incr(&value)
	fmt.Println(value)

	p := Point{X: 1, Y: 1}
	fmt.Println(p)
	incrPoint(&p)
	fmt.Println(p)

	i := newInt1()
	j := newInt2()

	fmt.Printf("i: %d, j: %d\n", *i, *j)
	fmt.Println(i == j)

	var fibboVal = 1000

	fmt.Printf("Fibonacci of %d is %d\n", fibboVal, fib(fibboVal))
}

func incr(p *int) {
	*p++
}

func incrPoint(p *Point) {
	p.X++
	p.Y = p.Y * 3
}

func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
