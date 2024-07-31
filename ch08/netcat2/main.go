package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var address = flag.String("address", "localhost:8000", "host address")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
