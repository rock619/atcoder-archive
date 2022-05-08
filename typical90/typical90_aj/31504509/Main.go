package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, Q int64, x []int64, y []int64, q []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	maxX, minX := int64(-1)<<60, int64(1)<<60
	maxY, minY := maxX, minX
	for i := int64(0); i < N; i++ {
		xi, yi := x[i]-y[i], x[i]+y[i]
		maxX = Max(maxX, xi)
		minX = Min(minX, xi)
		maxY = Max(maxY, yi)
		minY = Min(minY, yi)
	}

	for i := int64(0); i < Q; i++ {
		qi := q[i] - 1
		xi, yi := x[qi]-y[qi], x[qi]+y[qi]
		fmt.Fprintln(w, Max(Abs(xi-maxX), Abs(xi-minX), Abs(yi-maxY), Abs(yi-minY)))
	}
}

func Max(ints ...int64) int64 {
	if len(ints) == 0 {
		panic("Max: len(ints) == 0")
	}
	m := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] > m {
			m = ints[i]
		}
	}
	return m
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
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	x := make([]int64, N)
	y := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		x[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		y[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	q := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		q[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, Q, x, y, q)
}
