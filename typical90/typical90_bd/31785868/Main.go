package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const no = "Impossible"

func solve(N int64, S int64, A []int64, B []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	dp := make([][]bool, N+1)
	for i := range dp {
		dp[i] = make([]bool, S+1)
	}
	dp[0][0] = true
	for i := int64(1); i <= N; i++ {
		for j := int64(0); j <= S; j++ {
			if j-A[i-1] >= 0 && dp[i-1][j-A[i-1]] {
				dp[i][j] = true
			}
			if j-B[i-1] >= 0 && dp[i-1][j-B[i-1]] {
				dp[i][j] = true
			}
		}
	}

	if !dp[N][S] {
		fmt.Fprintln(w, no)
		return
	}

	result := make([]byte, N)
	for i, j := N, S; i > 0; i-- {
		if j-A[i-1] >= 0 && dp[i-1][j-A[i-1]] {
			result[i-1] = 'A'
			j -= A[i-1]
			continue
		}
		if j-B[i-1] >= 0 && dp[i-1][j-B[i-1]] {
			result[i-1] = 'B'
			j -= B[i-1]
		}
	}
	fmt.Fprintln(w, string(result))
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
	S, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	A := make([]int64, N)
	B := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		B[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, S, A, B)
}
