package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, K int64, X []int64, Y []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	d := make([][]int64, N)
	for i := int64(0); i < N; i++ {
		d[i] = make([]int64, N)
	}
	for i := int64(0); i < N; i++ {
		for j := i + 1; j < N; j++ {
			v := (X[i]-X[j])*(X[i]-X[j]) + (Y[i]-Y[j])*(Y[i]-Y[j])
			d[i][j] = v
			d[j][i] = v
		}
	}

	cost := make([]int64, 1<<N)
	for i := int64(1); i < (1 << N); i++ {
		for j := int64(0); j < N; j++ {
			for k := int64(0); k < j; k++ {
				if ((i>>j)&1) == 1 && ((i>>k)&1) == 1 {
					if d[j][k] > cost[i] {
						cost[i] = d[j][k]
					}
				}
			}
		}
	}

	dp := make([][]int64, K+1)
	for i := range dp {
		dp[i] = make([]int64, 1<<N)
		for j := range dp[i] {
			dp[i][j] = 1 << 62
		}
	}

	dp[0][0] = 0
	for i := int64(1); i <= K; i++ {
		for j := int64(1); j < (1 << N); j++ {
			for k := j; k != 0; k = (k - 1) & j {
				v := dp[i-1][j-k]
				if v < cost[k] {
					v = cost[k]
				}
				if v < dp[i][j] {
					dp[i][j] = v
				}
			}
		}
	}

	fmt.Fprintln(w, dp[K][(1<<N)-1])
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
	K, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	X := make([]int64, N)
	Y := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		X[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		Y[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, K, X, Y)
}
