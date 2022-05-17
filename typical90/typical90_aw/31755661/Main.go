package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type DSU struct {
	parentOrSize []int64
}

func NewDSU(n int64) *DSU {
	parentOrSize := make([]int64, n)
	for i := range parentOrSize {
		parentOrSize[i] = -1
	}

	return &DSU{
		parentOrSize,
	}
}

func (dsu *DSU) Merge(a, b int64) int64 {
	x, y := dsu.Leader(a), dsu.Leader(b)
	if x == y {
		return x
	}
	if -dsu.parentOrSize[x] < -dsu.parentOrSize[y] {
		x, y = y, x
	}
	dsu.parentOrSize[x] += dsu.parentOrSize[y]
	dsu.parentOrSize[y] = x
	return x
}

func (dsu *DSU) Same(a, b int64) bool {
	return dsu.Leader(a) == dsu.Leader(b)
}

func (dsu *DSU) Leader(a int64) int64 {
	if dsu.parentOrSize[a] < 0 {
		return a
	}
	dsu.parentOrSize[a] = dsu.Leader(dsu.parentOrSize[a])
	return dsu.parentOrSize[a]
}

func (dsu *DSU) Size(a int64) int64 {
	return -dsu.parentOrSize[dsu.Leader(a)]
}

func (dsu *DSU) Groups() [][]int64 {
	l := int64(len(dsu.parentOrSize))
	m := make(map[int64][]int64, l)
	for i := int64(0); i < l; i++ {
		leader := dsu.Leader(i)
		m[leader] = append(m[leader], i)
	}

	result := make([][]int64, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

type Edge struct {
	From, To, Weight int64
}

func solve(N int64, M int64, C []int64, L []int64, R []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	dsu := NewDSU(N + 1)
	edges := make([]Edge, M)
	for i := int64(0); i < M; i++ {
		edges[i] = Edge{
			From:   L[i] - 1,
			To:     R[i],
			Weight: C[i],
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	result := int64(0)
	for _, e := range edges {
		if dsu.Same(e.From, e.To) {
			continue
		}

		dsu.Merge(e.From, e.To)
		result += e.Weight

		if dsu.Size(e.From) == N+1 {
			fmt.Fprintln(w, result)
			return
		}
	}

	fmt.Fprintln(w, -1)
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
	C := make([]int64, M)
	L := make([]int64, M)
	R := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		C[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
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

	solve(N, M, C, L, R)
}
