// Dup loops over the lines in a file and
// prints duplicated lines
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// mapSort type is a helper struct used to sort a map[string]int.
type mapSort struct {
	srtKeys     []string
	underlaying map[string]int
}

// m.Len returns the length of the map (which should also be the length
// final & capacity of the slice).
func (m mapSort) Len() int { return len(m.underlaying) }

// m.Less provides a function to compare the keys acording to their values:
// neccessary to satisfy the sort Interface.
func (m mapSort) Less(i, j int) bool {
	return m.underlaying[m.srtKeys[i]] < m.underlaying[m.srtKeys[j]]
}

// m.Swap swaps elements in place in the string slice.
func (m mapSort) Swap(i, j int) {
	m.srtKeys[i], m.srtKeys[j] = m.srtKeys[j], m.srtKeys[i]
}

// sortMap returns a string slice containing the keys of the map
// sorted according to the map value, so that looping over the strings
// in the slice and accessing map[str] yeilds the map values in order.
// Complexity: O(n) to create slice with keys & O(n.log n) to sort => O(n.log n)
func sortMap(initial map[string]int) []string {
	keys := make([]string, 0, len(initial))

	for mapKey := range initial {
		keys = append(keys, mapKey)
	}

	m := mapSort{keys, initial}

	sort.Sort(sort.Reverse(m))

	return m.srtKeys
}

// countLines reads lines from a file pointer and counts the number of
// occurrances of each line on a given map
func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		count[input.Text()]++
	}
}

// dupV1 reads lines from stdin and then prints the repeated lines and
// the number of occurrances, in sorted order (descending).
func dupV1() {
	lines := make(map[string]int)

	countLines(os.Stdin, lines)

	for _, key := range sortMap(lines) {
		if lines[key] > 1 {
			fmt.Printf("%-d: \t%s\n", lines[key], key)
		}
	}

}

// dupV2 counts the number of duplicated lines in a given source. If arguments are
// passed in the command line, dupV2 will attempt to open those files and read from
// them. If no command line arguments are passed, the dupV2 reads from stdin.
func dupV2() {

	lines := make(map[string]int)
	if len(os.Args) < 2 {
		countLines(os.Stdin, lines)
	} else {
		for _, fname := range os.Args[1:] {
			f, err := os.Open(fname)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
			func(file *os.File) {
				defer file.Close()
				countLines(file, lines)
			}(f)
		}
	}

	for _, key := range sortMap(lines) {
		if lines[key] > 1 {
			fmt.Printf("%-d: \t%s\n", lines[key], key)
		}
	}
}

// Modify dup2 to print the name of the files in which there are duplicated lines.
func exercise1_4() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "Usage: ./dup filename [more filenames]\n")
	}

	for _, fname := range os.Args[1:] {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		func(f *os.File) {
			defer f.Close()
			lines := make(map[string]int)
			input := bufio.NewScanner(f)

			for input.Scan() {
				line := input.Text()
				if lines[line] > 0 {
					fmt.Println(fname, "has duplicated lines!")
					return
				}
				lines[line]++
			}
		}(f)
	}
}

func main() {
	exercise1_4()
}
