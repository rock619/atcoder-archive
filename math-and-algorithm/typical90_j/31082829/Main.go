package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, C []int64, P []int64, Q int64, L []int64, R []int64) {
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

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := int64(0); i < Q; i++ {
		fmt.Fprintf(w, "%d %d\n", sums[0][R[i]]-sums[0][L[i]-1], sums[1][R[i]]-sums[1][L[i]-1])
	}
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
	C := make([]int64, N)
	P := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		C[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		P[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	var Q int64
	scanner.Scan()
	Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	L := make([]int64, Q)
	R := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		L[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		R[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, C, P, Q, L, R)
}
