package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, A []int64, Q int64, L, R, X []int64) {
	m := make(map[int64][]int64, N)
	for i, a := range A {
		m[a] = append(m[a], int64(i))
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := int64(0); i < Q; i++ {
		indexes := m[X[i]]
		if len(indexes) == 0 {
			fmt.Fprintln(w, 0)
			continue
		}

		li := sort.Search(len(indexes), func(j int) bool {
			return indexes[j] >= L[i]-1
		})
		if li == len(indexes) {
			fmt.Fprintln(w, 0)
			continue
		}

		ri := sort.Search(len(indexes), func(j int) bool {
			return indexes[j] > R[i]-1
		})

		fmt.Fprintln(w, ri-li)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	var Q int64
	scanner.Scan()
	Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	L := make([]int64, Q)
	R := make([]int64, Q)
	X := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		L[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		R[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		X[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A, Q, L, R, X)
}
