package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, P []int64, Q []int64) {
	sum := float64(0)
	for i := range P {
		sum += float64(Q[i]) / float64(P[i])
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
	P := make([]int64, N)
	Q := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		P[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		Q[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	solve(N, P, Q)
}
