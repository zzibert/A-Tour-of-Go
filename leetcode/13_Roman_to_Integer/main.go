package main

func romanToInt(s string) int {
	var amount int

	symbols := make(map[string]int)

	symbols["I"] = 1
	symbols["V"] = 5
	symbols["X"] = 10
	symbols["L"] = 50
	symbols["C"] = 100
	symbols["D"] = 500
	symbols["M"] = 1000

	lenght := len(s)

	for i := 0; i < lenght; i++ {
		if i == lenght-1 {
			amount += symbols[string(s[i])]
		} else {
			if string(s[i]) == "I" {
				switch string(s[i+1]) {
				case "V":
					amount += 4
					i++
				case "X":
					amount += 9
					i++
				default:
					amount++
				}
			} else if string(s[i]) == "V" {
				amount += 5
			} else if string(s[i]) == "X" {
				switch string(s[i+1]) {
				case "L":
					amount += 40
					i++
				case "C":
					amount += 90
					i++
				default:
					amount += 10
				}
			} else if string(s[i]) == "L" {
				amount += 50
			} else if string(s[i]) == "C" {
				switch string(s[i+1]) {
				case "D":
					amount += 400
					i++
				case "M":
					amount += 900
					i++
				default:
					amount += 100
				}
			} else if string(s[i]) == "D" {
				amount += 500
			} else if string(s[i]) == "M" {
				amount += 1000
			}
		}
	}

	return amount
}
