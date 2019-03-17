package main

import (
	"fmt"
	"math"
)

func main() {
	var acc, initialV, initialS float64

	fmt.Printf("\nEnter acceleration value: ")
	fmt.Scan(&acc)

	fmt.Printf("\nEnter initial velocity value: ")
	fmt.Scan(&initialV)

	fmt.Printf("\nEnter initial displacement value: ")
	fmt.Scan(&initialS)

	var time float64

	fmt.Printf("\nEnter time value: ")
	fmt.Scan(&time)

	//Calculate displacement
	fn := GenDisplaceFn(acc, initialV, initialS)
	fmt.Printf("\nDisplacement = %f\n", fn(time))
}

func GenDisplaceFn(acc, initialV, initialS float64) func(float64) float64 {
	funcTime := func(time float64) float64 {
		return ((0.5) * (acc) * math.Pow(time, 2)) + (initialV * time) + initialS
	}
	return funcTime
}
