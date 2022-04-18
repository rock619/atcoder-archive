package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, T []int64) {
	sum := int64(0)
	for _, t := range T {
		sum += t
	}
	dp := make([][]bool, N+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}
	dp[0][0] = true
	for i := int64(1); i <= N; i++ {
		for j := int64(0); j <= sum; j++ {
			if dp[i-1][j] {
				dp[i][j] = true
				dp[i][j+T[i-1]] = true
			}
		}
	}

	min := sum
	for i, v := range dp[N] {
		if !v {
			continue
		}
		min = Min(min, Max(int64(i), sum-int64(i)))
	}
	fmt.Println(min)
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	T := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		T[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, T)
}
