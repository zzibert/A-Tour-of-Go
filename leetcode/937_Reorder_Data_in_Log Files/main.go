package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareLogs("g1 act car", "a8 act zoo"))
}

func reorderLogFiles(logs []string) []string {
	var letterLogs, digitLogs []string

	for i := 0; i < len(logs); i++ {
		log := strings.Split(logs[i], " ")[1]
		if isDigitLog(log) {
			digitLogs = append(digitLogs, logs[i])
		} else {
			letterLogs = append(letterLogs, logs[i])
		}
	}

	sortLogs(letterLogs)

	return append(letterLogs, digitLogs...)
}

func sortLogs(logs []string) {
	var temp string
	var index int
	for i := 0; i < len(logs)-1; i++ {
		temp = logs[i]
		index = i
		for j := i + 1; j < len(logs); j++ {
			if !compareLogs(temp, logs[j]) {
				temp = logs[j]
				index = j
			}
		}
		logs[index] = logs[i]
		logs[i] = temp
	}
}

func isDigitLog(log string) bool {
	for _, digit := range log {
		if _, err := strconv.Atoi(string(digit)); err != nil {
			return false
		}
	}
	return true
}

func compareLogs(log1, log2 string) bool {

	identifier1 := strings.Split(log1, " ")[0]
	words1 := strings.Split(log1, " ")[1:]

	identifier2 := strings.Split(log2, " ")[0]
	words2 := strings.Split(log2, " ")[1:]

	for i := 0; i < len(words1); i++ {
		if i >= len(words2) {
			return false
		}
		if words1[i] > words2[i] {
			return false
		}

		if words1[i] < words2[i] {
			return true
		}
	}

	for i := 0; i < len(identifier1); i++ {
		if i >= len(identifier2) {
			return false
		}
		if identifier1[i] > identifier2[i] {
			return false
		}
	}

	return true
}
