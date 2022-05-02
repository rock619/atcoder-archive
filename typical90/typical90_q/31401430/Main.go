package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type FenwickTree struct {
	s []int64
}

func NewFenwickTree(n int64) *FenwickTree {
	return &FenwickTree{
		s: make([]int64, n+1),
	}
}

func (f *FenwickTree) Add(i, x int64) {
	for index := i; index < int64(len(f.s)); index += (index & -index) {
		f.s[index] += x
	}
}

func (f *FenwickTree) Sum(i int64) int64 {
	sum := int64(0)
	for index := i; index > 0; index -= (index & -index) {
		sum += f.s[index]
	}
	return sum
}

type Edge struct {
	L, R int64
}

func solve(N int64, M int64, L []int64, R []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	all := M * (M - 1) / 2

	rCount := make(map[int64]int64, N)
	edgeCount := make(map[int64]int64, N)
	edges := make([]Edge, M)
	for i := int64(0); i < M; i++ {
		rCount[R[i]]++
		edgeCount[L[i]]++
		edgeCount[R[i]]++
		edges[i] = Edge{L: L[i], R: R[i]}
	}
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].R == edges[j].R {
			return edges[i].L < edges[j].L
		}
		return edges[i].R < edges[j].R
	})

	type1Count := int64(0)
	for _, c := range edgeCount {
		type1Count += c * (c - 1) / 2
	}

	v := make([]int64, N+1)
	v[1] = rCount[1]
	for i := int64(2); i <= N; i++ {
		v[i] = v[i-1] + rCount[i]
	}

	type2Count := int64(0)
	type3Count := int64(0)
	fenwickTree := NewFenwickTree(N)
	for i, edge := range edges {
		if L[i]-1 >= 0 {
			type2Count += v[L[i]-1]
		}

		type3Count += int64(i) - fenwickTree.Sum(edge.L)

		fenwickTree.Add(edge.L, 1)
	}

	fmt.Fprintln(w, all-(type1Count+type2Count+type3Count))
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
	M, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	L := make([]int64, M)
	R := make([]int64, M)
	for i := int64(0); i < M; i++ {
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
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, M, L, R)
}
