package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(W int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	a := make([]int64, 0, 300)
	for i := int64(1); i < 1000_000; i *= 100 {
		for j := int64(1); j <= 99; j++ {
			a = append(a, i*j)
		}
	}
	fmt.Fprintln(w, len(a))
	for i := range a {
		if i == 0 {
			fmt.Fprint(w, a[i])
		} else {
			fmt.Fprintf(w, " %d", a[i])
		}
	}
	fmt.Fprintln(w)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	W, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(W)
}
