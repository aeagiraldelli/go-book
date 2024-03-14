package main

import "fmt"

func nonempty(strs []string) []string {
	i := 0
	for _, s := range strs {
		if s != "" {
			strs[i] = s
			i++
		}
	}

	return strs[:i]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}
