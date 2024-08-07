package main

import (
	"fmt"
	"go-book/ch05/links"
	"log"
	"os"
	"strings"
)

var tokens = make(chan struct{}, 20)

func main() {
	worklist := make(chan []string)
	var n int

	n++
	go func() {
		for _, link := range os.Args[1:] {
			if !strings.HasPrefix(link, "http") {
				link = "https://" + link
			}
			worklist <- []string{link}
		}
	}()

	// crawl the web concurrently
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire token
	list, err := links.Extract(url)
	<-tokens // release token
	if err != nil {
		log.Print(err)
	}
	return list
}
