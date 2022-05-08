package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(H int64, W int64, R int64, C int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	count := 0
	if R > 1 {
		count++
	}
	if R < H {
		count++
	}
	if C > 1 {
		count++
	}
	if C < W {
		count++
	}
	fmt.Fprintln(w, count)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	H, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	W, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	R, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	C, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(H, W, R, C)
}
