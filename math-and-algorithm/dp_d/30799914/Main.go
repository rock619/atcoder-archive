package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(N int64, W int64, w []int64, v []int64) {
	dp := make([][]int64, N+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int64, W+1)
	}
	dp[0][0] = 0
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = math.MinInt64
	}

	for i := int64(1); i <= N; i++ {
		for j := int64(0); j <= W; j++ {
			if j < w[i-1] {
				dp[i][j] = dp[i-1][j]
				continue
			}

			dp[i][j] = Max(dp[i-1][j], dp[i-1][j-w[i-1]]+v[i-1])
		}
	}

	fmt.Println(Max(dp[N]...))
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var W int64
	scanner.Scan()
	W, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	w := make([]int64, N)
	v := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		w[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		v[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, W, w, v)
}
