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

	go func() {
		for _, link := range os.Args[1:] {
			if !strings.HasPrefix(link, "http") {
				link = "https://" + link
			}
			worklist <- []string{link}
		}
	}()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
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
