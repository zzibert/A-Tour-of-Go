package main

import (
	"fmt"
	"sync"
)

var i = 0
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}

func inc() {
	i = i + 1
	wg.Done()
}
