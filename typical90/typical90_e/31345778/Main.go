package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

func ModPow(a, b, m int64) int64 {
	p := a
	result := int64(1)
	for i := 0; i < 63; i++ {
		if b&(1<<i) != 0 {
			result *= p
			result %= m
		}
		p *= p
		p %= m
	}

	return result
}

func solve(N int64, B int64, K int64, c []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	pow10s := make([]int64, 63)
	for i := 0; i <= 62; i++ {
		pow10s[i] = ModPow(10, 1<<i, B)
	}

	dp := make([][]int64, 63)
	for i := range dp {
		dp[i] = make([]int64, B)
	}

	for _, cc := range c {
		dp[0][cc%B]++
	}

	for i := int64(0); i < 62; i++ {
		for j := int64(0); j < B; j++ {
			for k := int64(0); k < B; k++ {
				next := (j*pow10s[i] + k) % B
				dp[i+1][next] += dp[i][j] * dp[i][k]
				dp[i+1][next] %= mod
			}
		}
	}

	result := make([][]int64, 63)
	for i := range result {
		result[i] = make([]int64, B)
	}
	result[0][0] = 1
	for i := int64(0); i < 62; i++ {
		if N&(1<<i) == 0 {
			for j := int64(0); j < B; j++ {
				result[i+1][j] = result[i][j]
			}
			continue
		}

		for j := int64(0); j < B; j++ {
			for k := int64(0); k < B; k++ {
				next := (j*pow10s[i] + k) % B
				result[i+1][next] += result[i][j] * dp[i][k]
				result[i+1][next] %= mod
			}
		}
	}

	fmt.Fprintln(w, result[62][0])
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
	B, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	K, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	c := make([]int64, K)
	for i := int64(0); i < K; i++ {
		scanner.Scan()
		c[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, B, K, c)
}
