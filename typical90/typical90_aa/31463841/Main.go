package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, S []string) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	m := make(map[string]struct{}, N)
	for i, s := range S {
		if _, ok := m[s]; ok {
			continue
		}

		m[s] = struct{}{}
		fmt.Fprintln(w, i+1)
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
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	S := make([]string, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		S[i] = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, S)
}
