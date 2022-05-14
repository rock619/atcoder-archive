package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A []int64, B []int64, C []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	mA := make(map[int64]int64, 46)
	mB := make(map[int64]int64, 46)
	mC := make(map[int64]int64, 46)
	for i := int64(0); i < N; i++ {
		mA[A[i]%46]++
		mB[B[i]%46]++
		mC[C[i]%46]++
	}

	sum := int64(0)
	for i, a := range mA {
		for j, b := range mB {
			for k, c := range mC {
				if (i+j+k)%46 != 0 {
					continue
				}

				sum += a * b * c
			}
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
	B := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		B[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	C := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		C[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, A, B, C)
}
