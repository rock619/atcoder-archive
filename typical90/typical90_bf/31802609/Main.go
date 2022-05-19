package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod int64 = 100_000

func solve(N int64, K int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	logK := int64(1)
	for 1<<logK <= K {
		logK++
	}
	doubling := make([][]int64, logK)
	doubling[0] = make([]int64, mod)
	for j := int64(0); j < mod; j++ {
		digitsSum := int64(0)
		for k := int64(1); k < mod; k *= 10 {
			digitsSum += (j / k) % 10
		}
		doubling[0][j] = (j + digitsSum) % mod
	}
	for i := int64(1); i < logK; i++ {
		doubling[i] = make([]int64, mod)
		for j := int64(0); j < mod; j++ {
			doubling[i][j] = doubling[i-1][doubling[i-1][j]]
		}
	}

	result := N
	for n, k := int64(0), K; k > 0; n++ {
		if k&1 != 0 {
			result = doubling[n][result]
		}
		k >>= 1
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
