package main

import "strings"

func wordFrequency(input string) map[string]int {
	words := strings.Fields(input)

	counts := make(map[string]int)
	for _, word := range words {
		lower := strings.ToLower(word)
		counts[lower]++
	}

	return counts
}
