package main

import (
	"github.com/allegro/bigcache"
	"log"
	"time"
)

func main() {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		log.Println(err)
		return
	}
	entry, err := cache.Get("my-unique-key")
	if err != nil {
		log.Println(err)
		return
	}

	if err != nil {
		entry := []byte("value")
		cache.Set("my-unique-key", entry)
	}
	log.Println(string(entry))
}
