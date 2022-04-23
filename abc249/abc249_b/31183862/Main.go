package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	yes = "Yes"
	no  = "No"
)

func solve(S string) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	m := make(map[rune]struct{}, len(S))
	containsUpper, containsLower := false, false
	for _, r := range S {
		if _, ok := m[r]; ok {
			fmt.Fprintln(w, no)
			return
		}
		m[r] = struct{}{}

		if r >= 'A' && r <= 'Z' {
			containsUpper = true
		} else if r >= 'a' && r <= 'z' {
			containsLower = true
		}
	}

	if containsUpper && containsLower {
		fmt.Fprintln(w, yes)
	} else {
		fmt.Fprintln(w, no)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	S := scanner.Text()
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(S)
}
