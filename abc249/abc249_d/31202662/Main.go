package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, A []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	counts := make(map[int64]int64, N)
	for _, n := range A {
		counts[n]++
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	sums := make([]int64, N)
	for i, a := range A {
		if i > 0 && a == A[i-1] {
			sums[i] = sums[i-1]
			continue
		}
		for j := int64(1); j*j <= a; j++ {
			if _, ok := counts[j]; !ok || a%j != 0 {
				continue
			}
			q := a / j
			if _, ok := counts[q]; !ok {
				continue
			}

			if j*j == a {
				sums[i] += counts[j] * counts[j]
			} else {
				sums[i] += 2 * counts[j] * counts[q]
			}
		}
	}

	result := int64(0)
	for _, s := range sums {
		result += s
	}
	fmt.Fprintln(w, result)
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
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, A)
}
