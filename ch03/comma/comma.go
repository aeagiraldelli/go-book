package main

import "fmt"

func main() {
	numStr := "2345678098"
	dottedNum := comma(numStr)
	fmt.Printf("%s\n", numStr)
	fmt.Printf("%s\n", dottedNum)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "." + s[n-3:]
}
