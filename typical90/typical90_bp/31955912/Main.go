package main

import (
	"bufio"
	"fmt"
	"os"
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

func solve(N int64, Q int64, T []int64, X []int64, Y []int64, V []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	sums := make([]int64, N-1)
	for i := int64(0); i < Q; i++ {
		if x := X[i] - 1; T[i] == 0 {
			sums[x] = V[i]
		}
	}

	potentials := make([]int64, N)
	for i := int64(1); i < N; i++ {
		potentials[i] = sums[i-1] - potentials[i-1]
	}

	dsu := NewDSU(N)
	for i := int64(0); i < Q; i++ {
		switch x, y, v := X[i]-1, Y[i]-1, V[i]; {
		case T[i] == 0:
			dsu.Merge(x, y)
		case !dsu.Same(x, y):
			fmt.Fprintln(w, "Ambiguous")
		case (x+y)%2 == 0:
			fmt.Fprintln(w, v+potentials[y]-potentials[x])
		default:
			fmt.Fprintln(w, potentials[x]+potentials[y]-v)
		}
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
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	T := make([]int64, Q)
	X := make([]int64, Q)
	Y := make([]int64, Q)
	V := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		T[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
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
		scanner.Scan()
		V[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, Q, T, X, Y, V)
}
