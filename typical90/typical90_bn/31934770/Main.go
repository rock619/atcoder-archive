package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, L []int64, R []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	e := 0.0
	for i := int64(0); i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			count := 0
			all := 0
			for k := L[i]; k <= R[i]; k++ {
				for l := L[j]; l <= R[j]; l++ {
					if k > l {
						count++
					}
					all++
				}
			}

			e += float64(count) / float64(all)
		}

	}
	fmt.Fprintf(w, "%.15f", e)
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
	L := make([]int64, N)
	R := make([]int64, N)
	for i := int64(0); i < N; i++ {
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

	solve(N, L, R)
}
