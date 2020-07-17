package main

func minCost(costs [][]int) int {
	if len(costs) == 1 {

	}
}

func findMin(house1, house2 []int) (cost, color1, color2 int) {
	cost = house1[0] + house2[1]

	for i := 0; i < len(house1); i++ {
		for j := 0; j < len(house1) && j != i; j++ {
			if (house1[i] + house2[j]) < cost {
				cost = house1[i] + house2[j]
				color1 = i
				color2 = j
			}
		}
	}
	return
}
