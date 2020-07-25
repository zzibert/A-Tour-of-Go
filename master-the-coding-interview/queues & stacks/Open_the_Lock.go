package main

func openLock(deadends []string, target string) int {
	queue := make([]string, 0)
	seen := make(map[string]bool)

	start = "0000"

	if includes(deadends, start) {
		return -1
	}
}

func includes(array []string, combination string) bool {
	for i_, val := range array {
		if combination == val {
			return true
		}
	}
	return false
}
