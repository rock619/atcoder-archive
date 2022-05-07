package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	first  = "First"
	second = "Second"
)

func solve(N int64, W []int64, B []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	maxW := int64(50)
	maxB := int64(50) * maxW * (maxW - 1) / 2
	grundy := make([][]int64, maxW+1)
	for i := int64(0); i <= maxW; i++ {
		grundy[i] = make([]int64, maxB+1)
		for j := int64(0); j <= 1500; j++ {
			mex := make([]int64, maxB+1)
			if i >= 1 {
				mex[grundy[i-1][j+i]] = 1
			}
			if j >= 2 {
				for k := int64(1); k <= j/2; k++ {
					mex[grundy[i][j-k]] = 1
				}
			}
			for k := int64(0); k <= maxB; k++ {
				if mex[k] == 0 {
					grundy[i][j] = k
					break
				}
			}
		}
	}

	sumXOR := int64(0)
	for i := int64(0); i < N; i++ {
		if i == N-1 {
			i += 0
		}
		sumXOR ^= grundy[W[i]][B[i]]
	}

	if sumXOR != 0 {
		fmt.Fprintln(w, first)
	} else {
		fmt.Fprintln(w, second)
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
	W := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		W[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	B := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		B[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, W, B)
}
