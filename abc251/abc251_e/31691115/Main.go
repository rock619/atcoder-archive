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

	dp := make([][]int64, N+1)
	for i := range dp {
		dp[i] = make([]int64, 2)
	}
	dp[1][0] = A[N-1]
	dp[1][1] = A[0]
	for i := int64(2); i <= N; i++ {
		dp[i][0] = dp[i-1][1]
		if dp[i-1][1] > dp[i-1][0] {
			dp[i][1] = dp[i-1][0] + A[i-1]
		} else {
			dp[i][1] = dp[i-1][1] + A[i-1]
		}
	}
	// fmt.Fprintln(w, dp[N-1][1])
	// fmt.Fprintln(w, dp[N][0])

	dp2 := make([][]int64, N+1)
	for i := range dp2 {
		dp2[i] = make([]int64, 2)
	}
	dp2[1][1] = A[0]
	for i := int64(2); i <= N; i++ {
		dp2[i][0] = dp2[i-1][1]
		if dp2[i-1][1] > dp2[i-1][0] {
			dp2[i][1] = dp2[i-1][0] + A[i-1]
		} else {
			dp2[i][1] = dp2[i-1][1] + A[i-1]
		}
	}
	// fmt.Fprintln(w, dp2[N][1])
	// fmt.Fprintln(w, dp2[N][0])

	if a, b := dp[N][0], dp2[N][1]; a < b {
		fmt.Fprintln(w, a)
	} else {
		fmt.Fprintln(w, b)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

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
