package main

func main() {
	input := "go go java python go java"
	counts := wordFrequency(input)
	printSortedByFrequency(counts)
}
