package main

import (
	"fmt"
	"sync"
	"time"
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

	const dayDuration = time.Hour * 24
	fmt.Println(dayDuration.Seconds())
}
