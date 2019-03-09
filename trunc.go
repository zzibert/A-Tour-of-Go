package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	y := int(x)
	fmt.Printf("%d", y)
}
