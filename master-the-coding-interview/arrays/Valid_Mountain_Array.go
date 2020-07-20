package main

func validMountainArray(A []int) bool {
	var top, bottom bool

	if len(A) < 3 {
		return false
	}

	for i := 0; i < len(A)-1; i++ {
		if !top {
			if decreasing(A, i) {
				top = true
			} else if !increasing(A, i) {
				return false
			} else {
				bottom = true
			}
		} else {
			if !decreasing(A, i) {
				return false
			}
		}
	}

	return top && bottom
}

func increasing(A []int, index int) bool {
	return A[index] < A[index+1]
}

func decreasing(A []int, index int) bool {
	return A[index] > A[index+1]
}
