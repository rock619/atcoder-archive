package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A []int64) {
	m := make(map[int64]int64)
	for _, a := range A {
		m[a]++
	}

	count := m[50000] * (m[50000] - 1) / 2
	for i := int64(1); i < int64(50000); i++ {
		count += m[i] * m[100000-i]
	}

	fmt.Println(count)
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
