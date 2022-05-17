package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

func solve(N int64, A [][]int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	result := int64(1)
	for i := int64(0); i < N; i++ {
		result *= Sum(A[i]...)
		result %= mod
	}
	fmt.Fprintln(w, result)
}

func Sum(ints ...int64) int64 {
	s := int64(0)
	for _, i := range ints {
		s += i
	}
	return s
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
	A := make([][]int64, N)
	for i := int64(0); i < N; i++ {
		A[i] = make([]int64, 6)
	}
	for i := int64(0); i < N; i++ {
		for j := int64(0); j < 6; j++ {
			scanner.Scan()
			A[i][j], err = strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, A)
}
