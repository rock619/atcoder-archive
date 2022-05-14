package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(S string) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	if len(S) == 3 {
		fmt.Fprintln(w, S+S)
	} else if len(S) == 2 {
		fmt.Fprintln(w, S+S+S)
	} else if len(S) == 1 {
		fmt.Fprintln(w, S+S+S+S+S+S)
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
