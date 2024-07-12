package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	url := os.Args[1]
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	filename, _, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("file saved: %s\n", filename)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err != nil {
		err = closeErr
	}
	return local, n, err
}
