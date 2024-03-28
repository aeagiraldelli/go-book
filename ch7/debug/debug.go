package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true // change to false and the program will crash

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}

	f(buf)
	if debug {
		fmt.Println("debug is on.")
	}
}

func f(out io.Writer) {
	// do something
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
