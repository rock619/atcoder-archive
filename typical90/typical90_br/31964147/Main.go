package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, X []int64, Y []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	sort.Slice(X, func(i, j int) bool {
		return X[i] < X[j]
	})

	sort.Slice(Y, func(i, j int) bool {
		return Y[i] < Y[j]
	})
	medianXBy2 := 2 * X[(N-1)/2]
	if N%2 == 0 {
		medianXBy2 = X[(N-1)/2] + X[(N-1)/2+1]
	}
	medianYBy2 := 2 * Y[(N-1)/2]
	if N%2 == 0 {
		medianYBy2 = Y[(N-1)/2] + Y[(N-1)/2+1]
	}

	distBy2 := int64(0)
	for i := int64(0); i < N; i++ {
		distBy2 += Abs(2*X[i]-medianXBy2) + Abs(2*Y[i]-medianYBy2)
	}
	fmt.Fprintln(w, distBy2/2)
}

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
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
	X := make([]int64, N)
	Y := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		X[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		Y[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, X, Y)
}
