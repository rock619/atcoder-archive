package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Job struct {
	C, D, S int64
}

func solve(N int64, D []int64, C []int64, S []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	jobs := make([]Job, 0, N)
	for i := int64(0); i < N; i++ {
		if C[i] > D[i] {
			continue
		}

		j := Job{
			C: C[i],
			D: D[i],
			S: S[i],
		}
		jobs = append(jobs, j)
	}
	if len(jobs) == 0 {
		fmt.Fprintln(w, 0)
		return
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].D < jobs[j].D
	})

	maxD := jobs[len(jobs)-1].D
	dp := make([][]int64, len(jobs))
	for i := range dp {
		dp[i] = make([]int64, maxD+1)
	}
	dp[0][jobs[0].C] = jobs[0].S

	for i := 0; i < len(dp)-1; i++ {
		for j := int64(0); j <= maxD; j++ {
			dp[i+1][j] = Max(dp[i+1][j], dp[i][j])
			if jobs[i+1].D >= j+jobs[i+1].C {
				dp[i+1][j+jobs[i+1].C] = Max(dp[i+1][j+jobs[i+1].C], dp[i][j]+jobs[i+1].S)
			}
		}
	}

	fmt.Fprintln(w, Max(dp[len(jobs)-1]...))
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
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	D := make([]int64, N)
	C := make([]int64, N)
	S := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		D[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		C[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		S[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, D, C, S)
}
