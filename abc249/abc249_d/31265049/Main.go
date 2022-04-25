package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	max := int64(0)
	counts := make(map[int64]int64, N)
	for _, n := range A {
		counts[n]++
		if n > max {
			max = n
		}
	}

	sum := int64(0)
	for i := int64(1); i <= max; i++ {
		for j := int64(1); i*j <= max; j++ {
			sum += counts[i] * counts[j] * counts[i*j]
		}
	}

	fmt.Fprintln(w, sum)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, A)
}
