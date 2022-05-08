package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A int64, B int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	width := N * B
	height := N * A
	for i := int64(0); i < height; i++ {
		for j := int64(0); j < width; j++ {
			c := '.'
			if (j/B)%2 == 1 {
				if (i/A)%2 == 0 {
					c = '#'
				}
			} else if (i/A)%2 == 1 {
				c = '#'
			}
			fmt.Fprint(w, string(c))
		}
		fmt.Fprintln(w)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	A, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	B, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, A, B)
}
