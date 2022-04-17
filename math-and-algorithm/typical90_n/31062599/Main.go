package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, A []int64, B []int64) {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})

	sum := int64(0)
	for i := range A {
		sum += Abs(A[i] - B[i])
	}
	fmt.Println(sum)
}

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
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
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	B := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A, B)
}
