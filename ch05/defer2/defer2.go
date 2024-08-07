package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStackTrace()
	f(3)
}

func printStackTrace() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
