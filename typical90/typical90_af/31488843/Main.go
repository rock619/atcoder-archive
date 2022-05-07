package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Pair struct {
	A, B int64
}

func solve(N int64, A [][]int64, M int64, X []int64, Y []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	m := make(map[Pair]struct{}, 2*M)

	for i := int64(0); i < M; i++ {
		x, y := X[i]-1, Y[i]-1
		m[Pair{A: x, B: y}] = struct{}{}
		m[Pair{A: y, B: x}] = struct{}{}
	}

	runners := make([]int64, N)
	for i := range runners {
		runners[i] = int64(i)
	}

	min := int64(-1)
	for {
		feasible := true
		for i := 1; i < len(runners); i++ {
			if _, ok := m[Pair{A: runners[i-1], B: runners[i]}]; ok {
				feasible = false
				break
			}
		}
		if feasible {

			sum := int64(0)
			for i, r := range runners {
				sum += A[r][i]
			}
			if min == -1 || sum < min {
				min = sum
			}
			// fmt.Fprintln(w, "feasible", "runners", runners, "sum", sum)
		}

		if !nextPermutation(runners) {
			break
		}
	}

	fmt.Fprintln(w, min)
}

func nextPermutation(s []int64) bool {
	left := len(s) - 2
	for ; left >= 0 && s[left] >= s[left+1]; left-- {
	}
	if left < 0 {
		return false
	}

	right := len(s) - 1
	for ; s[left] >= s[right]; right-- {
	}

	s[left], s[right] = s[right], s[left]
	left++
	right = len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return true
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
	A := make([][]int64, N)
	for i := int64(0); i < N; i++ {
		A[i] = make([]int64, N)
	}
	for i := int64(0); i < N; i++ {
		for j := int64(0); j < N; j++ {
			scanner.Scan()
			A[i][j], err = strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
		}
	}
	scanner.Scan()
	M, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	X := make([]int64, M)
	Y := make([]int64, M)
	for i := int64(0); i < M; i++ {
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

	solve(N, A, M, X, Y)
}
