package main

func sortArrayByParity(A []int) []int {
	even := make([]int, 0)
	odd := make([]int, 0)
	for _, v := range A {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return append(even, odd...)
}

func main() {

}
