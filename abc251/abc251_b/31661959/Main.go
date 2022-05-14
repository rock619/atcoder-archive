package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, W int64, A []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	m := make(map[int64]struct{}, W)
	for i := int64(0); i < N; i++ {
		if w := A[i]; w <= W {
			m[w] = struct{}{}
		}

		for j := i + 1; j < N; j++ {
			if w := A[i] + A[j]; w <= W {
				m[w] = struct{}{}
			}
			for k := j + 1; k < N; k++ {
				if w := A[i] + A[j] + A[k]; w <= W {
					m[w] = struct{}{}
				}
			}
		}
	}
	fmt.Fprintln(w, len(m))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	W, err := strconv.ParseInt(scanner.Text(), 10, 64)
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

	solve(N, W, A)
}
