package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	m := make(map[int64]int64, N)
	for i := int64(0); i < N; i++ {
		m[A[i]]++
	}
	all := N * (N - 1) / 2 * (N - 2) / 3
	result := all
	for _, v := range m {
		if v <= 1 {
			continue
		}
		result -= v * (v - 1) / 2 * (N - v)
		if v >= 3 {
			result -= v * (v - 1) / 2 * (v - 2) / 3
		}
	}
	fmt.Fprintln(w, result)
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
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, A)
}
