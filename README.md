# Word Frequency Counter

A simple Go program that counts how many times each word appears in a given piece of text, ignoring case differences and extra spaces.

## Objective

Given an input string, count the occurrences of each word and display the results sorted by frequency (highest first), with ties broken alphabetically.

### Example

**Input:**
```
go go java python go java
```

**Output:**
```
go: 3
java: 2
python: 1
```

## Requirements

- Ignore extra spaces between words.
- Treat uppercase and lowercase letters as the same word (e.g. `Go` and `go` count together).
- Output sorted by frequency, highest count first. Words with equal counts are sorted alphabetically.

## Project Structure

```
wordfreq/
├── main.go                    # Entry point — runs the program
├── wordfrequency.go           # Counts word occurrences into a map
└── printsortedfrequency.go    # Sorts and prints the results
```

All three files belong to `package main`, so they compile together as a single program with no imports needed between them.

## How It Works

1. **`wordfrequency.go`** — splits the input into words using `strings.Fields` (which automatically handles extra/multiple spaces), lowercases each word with `strings.ToLower`, and tallies occurrences in a `map[string]int` using `counts[word]++`.

2. **`printsortedfrequency.go`** — since Go maps have no guaranteed iteration order, this file converts the map into a slice of `wordCount` structs, sorts that slice using `sort.Slice` (by count descending, then alphabetically for ties), and prints each entry in `word: count` format.

3. **`main.go`** — calls the above two functions with the input string and runs the program.

## Running the Program

From inside the project folder:

```bash
go run .
```

> Note: use `go run .` (not `go run main.go`) since the program is split across multiple files — `.` tells Go to build everything in the directory together.

To compile a binary instead:

```bash
go build .
```

## Edge Cases Considered

- Empty input string (`""`) — returns no output without crashing.
- Input that is only whitespace — same as empty input.
- All identical words — correctly tallies into a single entry.
- Mixed-case repeats (e.g. `Go GO go`) — counted as one word, case-insensitively.
- Tied counts — resolved alphabetically for consistent, repeatable output.

## Author

Ayomide Ajisegiri 