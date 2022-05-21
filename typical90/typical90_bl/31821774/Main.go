package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func solve(N int64, Q int64, A []int64, L []int64, R []int64, V []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	sum := int64(0)
	diffs := make([]int64, N)
	diffs[0] = A[0]
	for i := int64(1); i < N; i++ {
		diffs[i] = A[i] - A[i-1]
		sum += Abs(diffs[i])
	}

	for i := int64(0); i < Q; i++ {
		if li := L[i] - 1; li > 0 {
			c := Abs(diffs[li])
			diffs[li] += V[i]
			sum += Abs(diffs[li]) - c
		}
		if ri := R[i]; ri < N {
			c := Abs(diffs[ri])
			diffs[ri] -= V[i]
			sum += Abs(diffs[ri]) - c
		}

		fmt.Fprintln(w, sum)
	}
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
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
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
	L := make([]int64, Q)
	R := make([]int64, Q)
	V := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
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
		scanner.Scan()
		V[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, Q, A, L, R, V)
}
