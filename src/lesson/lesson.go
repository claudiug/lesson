package main

import (
	"fmt"

	"time"
)

func main() {
	fmt.Printf("Hello world!\n")
	fmt.Println(time.Now().Date())
	s := make([]string, 3)
	fmt.Println(s)
	m := make(map[string] int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m["a"])
}
