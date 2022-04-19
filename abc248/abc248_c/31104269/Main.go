package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 998244353

func solve(N int64, M int64, K int64) {
	dp := make([][]int64, N)
	for i := range dp {
		dp[i] = make([]int64, K)
		if i == 0 {
			for j := int64(0); j < M; j++ {
				dp[i][j] = 1
			}
		}
	}
	for i := int64(1); i < N; i++ {
		for j := int64(1); j < K; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			if j-M-1 >= 0 {
				dp[i][j] -= dp[i-1][j-M-1]
				if dp[i][j] < 0 {
					dp[i][j] += MOD
				}
			}
			dp[i][j] %= MOD
		}
	}

	sum := int64(0)
	for _, count := range dp[N-1] {
		sum += count
		sum %= MOD
	}
	fmt.Println(sum)
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
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N, M, K)
}
