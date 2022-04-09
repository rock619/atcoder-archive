package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	YES = "Yes"
	NO  = "No"
)

func solve(N int64, S int64, A []int64) {
	dp := make([][]bool, N+1)
	for i := range dp {
		dp[i] = make([]bool, S+1)
	}
	dp[0][0] = true

	for i := int64(1); i <= N; i++ {
		for j := int64(0); j <= S; j++ {
			dp[i][j] = dp[i-1][j] || (j >= A[i-1] && dp[i-1][j-A[i-1]])

			if j == S && dp[i][j] {
				fmt.Println(YES)
				return
			}
		}
	}

	fmt.Println(NO)
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
	var S int64
	scanner.Scan()
	S, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, S, A)
}
