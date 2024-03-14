package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	BRL
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, el := range a {
		fmt.Printf("%d: %d\n", i, el)
	}

	for _, el := range a {
		fmt.Printf("%d\n", el)
	}

	q := [...]int{1, 2, 3}
	fmt.Println(len(q))
	fmt.Printf("%d\n", q[0])
	fmt.Printf("%T\n", q)

	symbols := [...]string{USD: "$", BRL: "R$"}
	fmt.Printf("%T\n", symbols)
	fmt.Printf("%s\n", symbols[BRL])
}
