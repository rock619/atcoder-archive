package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(H int64, W int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	if H == 1 || W == 1 {
		fmt.Fprintln(w, H*W)
	} else {
		fmt.Fprintln(w, ((H+1)/2)*((W+1)/2))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
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
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(H, W)
}
