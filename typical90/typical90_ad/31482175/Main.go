package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, K int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	if K == 1 {
		fmt.Fprintln(w, N-1)
		return
	}

	primeFactorCounts := make([]int64, N-1)
	for i := int64(2); i <= N; i++ {
		if primeFactorCounts[i-2] != 0 {
			continue
		}

		for j := i; j <= N; j += i {
			primeFactorCounts[j-2]++
		}
	}

	count := int64(0)
	for _, c := range primeFactorCounts {
		if c >= K {
			count++
		}
	}

	fmt.Fprintln(w, count)
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
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, K)
}
