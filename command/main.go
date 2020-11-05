package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var path string
	flag.StringVar(&path, "p", "", "path")
	flag.Parse()
	if len(os.Args) <= 1 {
		panic("")
	}
	fmt.Println(path)
}
