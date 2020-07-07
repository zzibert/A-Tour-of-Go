package main

func main() {}

func tribonacci(n int) int {
	t0 := 0
	t1 := 1
	t2 := 1

	if n == 0 {
		return t0
	} else if n == 1 || n == 2 {
		return t1
	}

	var value int

	for n > 2 {
		value = t2 + t1 + t0
		t0 = t1
		t1 = t2
		t2 = value
		n--
	}
	return value
}
