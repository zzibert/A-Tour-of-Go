package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acceleration, velocity, initial_displacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return ((0.5 * acceleration * math.Pow(time, 2)) + velocity*time + initial_displacement)
	}
}

func main() {
	var acceleration float64
	var initial_velocity float64
	var initial_displacement float64
	var time float64

	fmt.Print("Enter acceleration : ")
	fmt.Scanf("%f\n", &acceleration)

	fmt.Print("Enter initial velocity : ")
	fmt.Scanf("%f\n", &initial_velocity)

	fmt.Print("Enter initial displacement : ")
	fmt.Scanf("%f\n", &initial_displacement)

	fmt.Print("Enter time after initial displacement : ")
	fmt.Scanf("%f\n", &time)

	fn := GenDisplaceFn(acceleration, initial_velocity, initial_displacement)

	fmt.Println(fn(time))

}
