package main

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	domains := make(map[string]int)

	for _, cpdomain := range cpdomains {
		count, domain := strings.Split(cpdomain, " ")[0], strings.Split(cpdomain, " ")[1]
		number, _ := strconv.Atoi(count)
		countSubDomains(domain, domains, number)
	}

	cpdomains = []string{}
	for domain, count := range domains {
		cpdomain := strconv.Itoa(count) + " " + domain
		cpdomains = append(cpdomains, cpdomain)
	}

	return cpdomains
}

func countSubDomains(domain string, domains map[string]int, count int) {
	subdomains := strings.Split(domain, ".")
	for i := 0; i < len(subdomains); i++ {
		subdomain := ""
		for j := i; j < len(subdomains); j++ {
			if j == len(subdomains)-1 {
				subdomain += subdomains[j]
			} else {
				subdomain += subdomains[j] + "."
			}
		}
		domains[subdomain] += count
	}
}
