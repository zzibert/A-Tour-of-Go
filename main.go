package main

import (
	"fmt"
	"sync"
)

var i = 0
var wg sync.WaitGroup
var once sync.Once

func main() {
	wg.Add(3)
	go hello()
	go hello()
	go hello()
	wg.Wait()
}

func setup() {
	fmt.Println("Init")
}

func hello() {
	once.Do(setup)
	fmt.Println("Hello")
	wg.Done()
}
