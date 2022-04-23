package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, K int64, S []string) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	max := int64(0)
	for i := uint64(0); i < (1 << N); i++ {
		m := make(map[rune]int64)
		for j := int64(0); j < N; j++ {
			if i&(1<<j) != 0 {
				for _, r := range S[j] {
					m[r]++
				}
			}
		}

		c := int64(0)
		for _, v := range m {
			if v == K {
				c++
			}
		}
		if c > max {
			max = c
		}
	}

	fmt.Fprintln(w, max)
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
	scanner.Scan()
	K, err := strconv.ParseInt(scanner.Text(), 10, 64)
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

	solve(N, K, S)
}
