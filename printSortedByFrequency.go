package main

import (
	"fmt"
	"sort"
)

func printSortedByFrequency(counts map[string]int) {
	type wordCount struct {
		word  string
		count int
	}

	var list []wordCount
	for word, count := range counts {
		list = append(list, wordCount{word, count})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].count != list[j].count {
			return list[i].count > list[j].count
		}
		return list[i].word < list[j].word
	})

	for _, wc := range list {
		fmt.Printf("%s: %d\n", wc.word, wc.count)
	}
}
