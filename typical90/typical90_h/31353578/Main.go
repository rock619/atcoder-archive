package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

func solve(N int64, S string) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()
	str := []byte("atcoder")

	dp := make([][]int64, 7)
	for i := range dp {
		dp[i] = make([]int64, N)
	}

	count := int64(0)
	for j := int64(0); j < N; j++ {
		if S[j] == str[0] {
			count++
		}
		dp[0][j] = count
	}

	for i := int64(1); i < int64(len(str)); i++ {
		for j := i; j < N; j++ {
			if S[j] == str[i] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
				dp[i][j] %= mod
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	fmt.Fprintln(w, dp[len(str)-1][N-1])
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
	S := scanner.Text()
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, S)
}
