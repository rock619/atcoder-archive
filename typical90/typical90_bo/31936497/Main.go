package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(N string, K int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	n := int64(0)
	s := N
	x := int64(1)
	for i := len(s) - 1; i >= 0; i-- {
		digit, _ := strconv.ParseInt(string(s[i]), 10, 64)
		n += digit * x
		x *= 8
	}

	for i := int64(0); i < K; i++ {
		s = strconv.FormatInt(n, 9)
		s = strings.ReplaceAll(s, "8", "5")
		n, _ = strconv.ParseInt(s, 8, 64)
	}
	fmt.Fprintln(w, s)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	N := scanner.Text()
	scanner.Scan()
	K, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, K)
}
