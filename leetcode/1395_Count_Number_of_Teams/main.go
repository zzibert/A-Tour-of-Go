package main

func numTeams(rating []int) int {

	countTeams := 0

	for i := 0; i < len(rating); i++ {
		first := rating[i]
		for j := i + 1; j < len(rating); j++ {
			if rating[j] > first {
				second := rating[j]
				for k := j + 1; k < len(rating); k++ {
					if rating[k] > second {
						countTeams++
					}
				}
			}
		}
	}

	for i := 0; i < len(rating); i++ {
		first := rating[i]
		for j := i + 1; j < len(rating); j++ {
			if rating[j] < first {
				second := rating[j]
				for k := j + 1; k < len(rating); k++ {
					if rating[k] < second {
						countTeams++
					}
				}
			}
		}
	}

	return countTeams
}
