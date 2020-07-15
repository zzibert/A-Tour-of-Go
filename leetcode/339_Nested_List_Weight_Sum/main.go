package main

func depthSum(nestedList []*NestedInteger) int {

	return calculateSum(nestedList, 1)
}

func calculateSum(array []*NestedInteger, depth int) int {
	sum := 0
	for _, number := range array {
		if number.IsInteger() {
			sum += number.GetInteger() * depth
		} else {
			sum += calculateSum(number.GetList(), depth+1)
		}
	}

	return sum
}
