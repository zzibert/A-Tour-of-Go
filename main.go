package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}

func rollTwo(size1, size2 int) (int, int) {
	return dieRoll(size1), dieRoll(size2)
}

func main() {
	result1, result2 := rollTwo(20, 6)
	fmt.Println(result1, result2)
}
