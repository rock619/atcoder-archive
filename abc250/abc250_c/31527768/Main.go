package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, Q int64, x []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	nums := make([]int64, N)
	m := make(map[int64]int64, N)
	for i := int64(0); i < N; i++ {
		nums[i] = i + 1
		m[i+1] = i
	}

	for i := int64(0); i < Q; i++ {
		a := x[i]
		indexA := m[a]
		indexB := indexA + 1
		if indexB == N {
			indexB -= 2
		}
		b := nums[indexB]
		nums[indexA], nums[indexB] = nums[indexB], nums[indexA]
		m[a], m[b] = indexB, indexA
	}

	for i, v := range nums {
		if i == 0 {
			fmt.Fprint(w, v)
		} else {
			fmt.Fprintf(w, " %d", v)
		}
	}
	fmt.Fprintln(w)
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
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	x := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		x[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, Q, x)
}
