package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type FenwickTree struct {
	data []int64
}

func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		data: make([]int64, n+1),
	}
}

func (f *FenwickTree) Add(p int, x int64) {
	p++
	for p < len(f.data) {
		f.data[p-1] += x
		p += (p & -p)
	}
}

func (f *FenwickTree) Sum(l, r int) int64 {
	return f.sum(r) - f.sum(l)
}

func (f *FenwickTree) sum(r int) int64 {
	s := int64(0)
	for r > 0 {
		s += f.data[r-1]
		r -= (r & -r)
	}
	return s
}

type Query struct {
	Type, P, X, L, R int64
}

func solve(N, Q int64, a []int64, query []Query) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	ft := NewFenwickTree(int(N))
	for i := int64(0); i < N; i++ {
		ft.Add(int(i), a[i])
	}

	for _, q := range query {
		if q.Type == 0 {
			ft.Add(int(q.P), q.X)
			continue
		}

		fmt.Fprintln(w, ft.Sum(int(q.L), int(q.R)))
	}
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
	a := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		a[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	query := make([]Query, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		query[i].Type, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		if query[i].Type == 0 {
			scanner.Scan()
			query[i].P, err = strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
			scanner.Scan()
			query[i].X, err = strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
			continue
		}
		scanner.Scan()
		query[i].L, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		query[i].R, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, Q, a, query)
}
