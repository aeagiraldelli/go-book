package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func lookup(key string) string {
	cache.Lock()
	defer cache.Unlock()
	return cache.mapping[key]
}

func main() {
	cache.mapping["key"] = "map value"
	value := lookup("key")
	fmt.Println(value)
}
