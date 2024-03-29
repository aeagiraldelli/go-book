package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3, 5))
	fmt.Println(sum(3, 5, 7))

	values := []int{1, 2, 3, 4, 5, 6, 7}
	total := sum(values...)
	fmt.Println(total)
}
