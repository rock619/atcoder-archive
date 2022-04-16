package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, A []int64) {
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})

	cumulative := make([]int64, N)
	for i := len(cumulative) - 1; i > 0; i-- {
		if i == len(cumulative)-1 {
			cumulative[i] = A[i]
			continue
		}

		cumulative[i] = cumulative[i+1] + A[i]
	}

	sum := int64(0)
	for i := int64(0); i < N-1; i++ {
		sum += (N-i-1)*A[i] - cumulative[i+1]
	}

	fmt.Println(sum)
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
	solve(N, A)
}
