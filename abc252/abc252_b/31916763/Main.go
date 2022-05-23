package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const yes = "Yes"
const no = "No"

func solve(N int64, K int64, A []int64, B []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	max := int64(0)
	for i := int64(0); i < N; i++ {
		if A[i] > max {
			max = A[i]
		}
	}

	m := make(map[int64]struct{})
	for i := int64(0); i < N; i++ {
		if A[i] == max {
			m[i+1] = struct{}{}
		}
	}

	for _, b := range B {
		if _, ok := m[b]; ok {
			fmt.Fprintln(w, yes)
			return
		}
	}
	fmt.Fprintln(w, no)
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
	K, err := strconv.ParseInt(scanner.Text(), 10, 64)
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
	B := make([]int64, K)
	for i := int64(0); i < K; i++ {
		scanner.Scan()
		B[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, K, A, B)
}
