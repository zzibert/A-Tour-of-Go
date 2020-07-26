package main

func findRestaurant(list1 []string, list2 []string) []string {
	sum := len(list1) + len(list2)
	restaurants1 := make(map[string]int)
	restaurants2 := make(map[string]int)
	restaurants := make([]string, 0)

	sumIndexes := make(map[string]int)

	for i, restaurant := range list1 {
		restaurants1[restaurant] = i
	}

	for i, restaurant := range list2 {
		restaurants2[restaurant] = i
	}

	for _, restaurant := range list1 {
		if _, ok := restaurants2[restaurant]; ok {
			localSum := restaurants1[restaurant] + restaurants2[restaurant]
			sumIndexes[restaurant] = localSum
			if localSum < sum {
				sum = localSum
			}
		}
	}

	for restaurant, count := range sumIndexes {
		if count == sum {
			restaurants = append(restaurants, restaurant)
		}
	}

	return restaurants

}
