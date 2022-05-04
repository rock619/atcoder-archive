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

	dp := make([][]int64, 2*N)
	for i := int64(0); i < 2*N; i++ {
		dp[i] = make([]int64, 2*N)
		for j := i; j < 2*N; j++ {
			dp[i][j] = 1 << 60
		}
		if i < 2*N-1 {
			dp[i][i+1] = Abs(A[i] - A[i+1])
		}
	}

	for i := int64(2); i < 2*N-1; i += 2 {
		for j := int64(0); j < 2*N-i-1; j++ {
			l, r := j, i+j+1

			for k := l; k <= r-1; k++ {
				dp[l][r] = Min(dp[l][r], dp[l][k]+dp[k+1][r])
			}
			dp[l][r] = Min(dp[l][r], dp[l+1][r-1]+Abs(A[l]-A[r]))
		}
	}

	fmt.Fprintln(w, dp[0][2*N-1])
}

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
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
	A := make([]int64, 2*N)
	for i := int64(0); i < 2*N; i++ {
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
