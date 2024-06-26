package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	zero(&c1)
	zero_2nd(&c2)
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero_2nd(ptr *[32]byte) {
	*ptr = [32]byte{}
}
