package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, A []int64, Q int64, B []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	for _, b := range B {
		switch index := sort.Search(len(A), func(i int) bool { return A[i] >= b }); index {
		case 0:
			fmt.Fprintln(w, A[0]-b)
		case len(A):
			fmt.Fprintln(w, b-A[len(A)-1])
		default:
			fmt.Fprintln(w, Min(A[index]-b, b-A[index-1]))
		}
	}
}

func Min(ints ...int64) int64 {
	if len(ints) == 0 {
		panic("Min: len(ints) == 0")
	}
	m := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] < m {
			m = ints[i]
		}
	}
	return m
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
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	B := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		B[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, A, Q, B)
}
