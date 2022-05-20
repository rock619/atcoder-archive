package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	yes = "Yes"
	no  = "No"
)

func solve(N int64, M int64, Q int64, X []int64, Y []int64, A []int64, B []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]int64, N)
	for i := int64(0); i < M; i++ {
		from, to := X[i]-1, Y[i]-1
		graph[to] = append(graph[to], from)
	}

	ss := make([]int64, 0, 64)
	ts := make([]int64, 0, 64)
	result := make([]bool, 0, Q)
	for i := int64(0); i < Q; i++ {
		ss = append(ss, A[i]-1)
		ts = append(ts, B[i]-1)
		if len(ss) == 64 {
			result = append(result, solve64(graph, ss, ts)...)
			ss = ss[:0]
			ts = ts[:0]
		}
	}
	result = append(result, solve64(graph, ss, ts)...)

	for _, v := range result {
		if v {
			fmt.Fprintln(w, yes)
		} else {
			fmt.Fprintln(w, no)
		}
	}
}

func solve64(graph [][]int64, ss []int64, ts []int64) []bool {
	l := int64(len(ss))
	n := int64(len(graph))
	dp := make([]int64, n)
	for i := int64(0); i < l; i++ {
		s := ss[i]
		dp[s] += 1 << i
	}

	for i := int64(0); i < n; i++ {
		for _, f := range graph[i] {
			dp[i] = dp[i] | dp[f]
		}
	}

	result := make([]bool, l)
	for i, t := range ts {
		if dp[t]&(1<<i) != 0 {
			result[i] = true
		}
	}
	return result
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
	M, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	X := make([]int64, M)
	Y := make([]int64, M)
	for i := int64(0); i < M; i++ {
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
	A := make([]int64, Q)
	B := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
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

	solve(N, M, Q, X, Y, A, B)
}
