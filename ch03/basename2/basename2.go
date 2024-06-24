package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/usr/share/fonts/user-fonts/Inter/Inter-VariableFont_slnt,wght.ttf"
	filename := basename(path)
	fmt.Printf("path: %s\n", path)
	fmt.Printf("file basename: %s\n", filename)
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
