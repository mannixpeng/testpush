package main
import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)
func main() {
	c := cache.New(5*time.Second, 10*time.Second)

	c.Set("foo", "bar", cache.DefaultExpiration)
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}
	go func() {
		c2 := cache.New(10*time.Second, 10*time.Second)
		foo2, found := c2.Get("foo")
		if found {
			fmt.Printf("foo2 found %s\n", foo2)
		} else {
			fmt.Println("foo2 not found")
		}
	}()
	time.Sleep(10*time.Second)
	fmt.Println("end")

}