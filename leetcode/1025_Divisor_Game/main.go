package main

func divisorGame(N int) bool {
	divisors := make([]int, 0)
  counter := 0
  
	for i := N - 1; i > 0; i-- {
		if N%i == 0 {
			divisors = append(divisors, i)
		}
  }
  
  

}
