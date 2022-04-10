package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A []int64, M int64, B []int64) {
	sums := make([]int64, N)
	sums[0] = 0
	for i := 0; i < len(A); i++ {
		sums[i+1] = sums[i] + A[i]
	}

	result := int64(0)
	for i := 1; i < len(B); i++ {
		dist := sums[B[i]-1] - sums[B[i-1]-1]
		if dist < 0 {
			result -= dist
		} else {
			result += dist
		}
	}
	fmt.Println(result)
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
	A := make([]int64, N-1)
	for i := int64(0); i < N-1; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	B := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A, M, B)
}
