package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, L int64, K int64, A []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	max := int64(1)
	left, right := int64(1), L
	for right-left > 1 {
		center := (left + right) / 2
		if feasible(L, K, A, center) {
			max = Max(max, center)
			left = center
			continue
		}

		right = center
	}

	fmt.Println(max)
}

func feasible(L, K int64, A []int64, score int64) bool {
	usedIndex := -1
	for i := int64(0); i < K; i++ {
		index := sort.Search(len(A), func(i int) bool {
			if usedIndex == -1 {
				return A[i] >= score
			}
			return A[i]-A[usedIndex] >= score
		})
		if index == len(A) {
			return false
		}

		usedIndex = index
	}

	return L-A[usedIndex] >= score
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
	L, err := strconv.ParseInt(scanner.Text(), 10, 64)
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
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, L, K, A)
}
