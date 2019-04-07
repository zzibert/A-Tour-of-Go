package main

import (
	"sync"
)

var i = 0
var wg sync.WaitGroup
var once sync.Once

func main() {
	var c1, c2 chan int
	wg.Add(2)
	go doStuff(c1, c2)
	go doStuff(c2, c1)
	wg.Wait()
}

func doStuff(c1 chan int, c2 chan int) {
	<-c1
	c2 <- 1
	wg.Done()
}
