package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNameCounts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			fileNameCounts[filename] = counts
			counts = make(map[string]int)
			f.Close()
		}
	}
	for filename, countLines := range fileNameCounts {
		fmt.Printf("\n...:: %s ::...\n", filename)
		for line, number := range countLines {
			if number > 1 {
				fmt.Printf("%d\t\"%s\"\n", number, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
