package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 998244353

func solve(N int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	dp := make([][]int64, N)
	for i := range dp {
		dp[i] = make([]int64, 9)
		if i == 0 {
			for j := 0; j < 9; j++ {
				dp[i][j] = 1
			}
		}
	}

	for i := int64(1); i < N; i++ {
		for j := int64(0); j < 9; j++ {
			for k := Max(j-1, 0); k <= Min(j+1, 8); k++ {
				dp[i][j] += dp[i-1][k] % mod
				dp[i][j] %= mod
			}
		}
	}

	sum := int64(0)
	for _, n := range dp[N-1] {
		sum += n
		sum %= mod
	}
	fmt.Fprintln(w, sum)
}

func Max(ints ...int64) int64 {
	if len(ints) == 0 {
		panic("Max: len(ints) == 0")
	}
	m := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] > m {
			m = ints[i]
		}
	}
	return m
}

func Min(ints ...int64) int64 {
	if len(ints) == 0 {
		panic("Min: len(ints) == 0")
	}
	m := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] < m {
			m = ints[i]
		}
	}
	return m
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
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N)
}
