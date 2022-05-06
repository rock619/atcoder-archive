package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, lx []int64, ly []int64, rx []int64, ry []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	m := make([][]int64, 1001)
	for i := range m {
		m[i] = make([]int64, 1001)
	}

	for i := int64(0); i < N; i++ {
		m[ly[i]][lx[i]]++
		m[ly[i]][rx[i]]--
		m[ry[i]][lx[i]]--
		m[ry[i]][rx[i]]++
	}

	for i := 0; i < len(m); i++ {
		current := int64(0)
		for j := 0; j < len(m[i]); j++ {
			current += m[i][j]
			m[i][j] = current
		}
	}

	for i := 0; i < len(m[0]); i++ {
		current := int64(0)
		for j := 0; j < len(m); j++ {
			current += m[j][i]
			m[j][i] = current
		}
	}

	counts := make(map[int64]int64)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			counts[m[i][j]]++
		}
	}

	for i := int64(1); i <= N; i++ {
		fmt.Fprintln(w, counts[i])
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
	lx := make([]int64, N)
	ly := make([]int64, N)
	rx := make([]int64, N)
	ry := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		lx[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		ly[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		rx[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		ry[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, lx, ly, rx, ry)
}
