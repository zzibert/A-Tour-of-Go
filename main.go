package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fakeDieRoll(size int) int {
	return 42
}

func getDieRolls() []dieRollFunc {
	return []dieRollFunc{
		dieRoll,
		fakeDieRoll,
	}
}

type dieRollFunc func(int) int

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}

func rollTwo(size1, size2 int) (int, int) {
	return dieRoll(size1), dieRoll(size2)
}

func main() {
	rolls := getDieRolls()
	for index, rollFunc := range rolls {
		fmt.Printf("Die Roll Attempt #%d, result: %d\n", index, rollFunc(10))
	}
}
