package main

import (
	"fmt"
	"go-book/ch05/links"
	"log"
	"os"
	"strings"
)

func main() {
	worklist := make(chan []string)
	unseenlinks := make(chan string)

	go func() {
		for _, link := range os.Args[1:] {
			if !strings.HasPrefix(link, "http") {
				link = "https://" + link
			}
			worklist <- []string{link}
		}
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenlinks {
				foundlinks := crawl(link)
				go func() { worklist <- foundlinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenlinks <- link
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
