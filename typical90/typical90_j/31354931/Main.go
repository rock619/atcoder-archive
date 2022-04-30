package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, C []int64, P []int64, Q int64, L []int64, R []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	sums := make([][]int64, 2)
	for i := range sums {
		sums[i] = make([]int64, N+1)
	}
	for i := int64(0); i < 2; i++ {
		for j := int64(0); j < N; j++ {
			if C[j] == i+1 {
				sums[i][j+1] = sums[i][j] + P[j]
			} else {
				sums[i][j+1] = sums[i][j]
			}
		}
	}

	for i := int64(0); i < Q; i++ {
		fmt.Fprintf(w, "%d %d\n", sums[0][R[i]]-sums[0][L[i]-1], sums[1][R[i]]-sums[1][L[i]-1])
	}
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
	C := make([]int64, N)
	P := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		C[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		P[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	L := make([]int64, Q)
	R := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		L[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		R[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, C, P, Q, L, R)
}
