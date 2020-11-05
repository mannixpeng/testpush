package main

import (
	"fmt"
	"math/rand"
	"time"
)
import "github.com/allegro/bigcache"

func main1() {

	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

	cache.Set("my-unique-key", []byte("value"))

	entry, _ := cache.Get("my-unique-key")
	fmt.Println(string(entry))

	go func() {
		for i := 0; i < 1000; i++ {
			cache.Set(fmt.Sprintf("my-key1--%d", i), []byte(fmt.Sprintf("%d", i)))
		}
	}()
	go func() {
		for i := 1000; i < i+1000; i++ {
			cache.Set(fmt.Sprintf("my-key1--%d", i), []byte(fmt.Sprintf("%d", i)))
		}
	}()
	go func() {
		for i := 0; i < 2000; i++ {
			rand.Seed(time.Now().Unix())
			n := rand.Intn(2000)

			v, err := cache.Get(fmt.Sprintf("my-key1--%d", n))
			if err != nil{
				fmt.Println(err)
				continue
			}
			fmt.Println(string(v))
		}
	}()
	time.Sleep(1*time.Minute)
}

func main() {
	ints := make([]int, 1000000)
	go func() {
		for i := 0; i < 1000000; i++ {
			ints[i] = i
		}
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			fmt.Println(ints[i])
		}
	}()
	fmt.Println("hello world")

	fmt.Println()
}