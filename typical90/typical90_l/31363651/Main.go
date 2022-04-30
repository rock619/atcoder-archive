package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	yes = "Yes"
	no  = "No"
)

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
		size[i] = 1
	}
	return &UnionFind{
		parent,
		size,
	}
}

func (uf *UnionFind) root(x int) int {
	if uf.parent[x] == -1 {
		return x
	}
	uf.parent[x] = uf.root(uf.parent[x])
	return uf.parent[x]
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.root(x) == uf.root(y)
}

func (uf *UnionFind) Unite(x, y int) bool {
	x, y = uf.root(x), uf.root(y)
	if x == y {
		return false
	}

	if uf.size[x] < uf.size[y] {
		x, y = y, x
	}
	uf.parent[y] = x
	uf.size[x] += uf.size[y]
	return true
}

func solve(H, W, Q int64, q []Query) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	cells := make([][]int, H)
	for i := range cells {
		cells[i] = make([]int, W)
		for j := range cells[i] {
			cells[i][j] = -1
		}
	}
	uf := NewUnionFind(len(q))
	for i, query := range q {
		if query.Type == 1 {
			cells[query.R-1][query.C-1] = i
			if query.R >= 2 && cells[query.R-2][query.C-1] != -1 {
				uf.Unite(i, cells[query.R-2][query.C-1])
			}
			if query.R < H && cells[query.R][query.C-1] != -1 {
				uf.Unite(i, cells[query.R][query.C-1])
			}
			if query.C >= 2 && cells[query.R-1][query.C-2] != -1 {
				uf.Unite(i, cells[query.R-1][query.C-2])
			}
			if query.C < W && cells[query.R-1][query.C] != -1 {
				uf.Unite(i, cells[query.R-1][query.C])
			}
			continue
		}

		switch a, b := cells[query.RA-1][query.CA-1], cells[query.RB-1][query.CB-1]; {
		case a == -1, b == -1, !uf.Same(a, b):
			fmt.Fprintln(w, no)
		default:
			fmt.Fprintln(w, yes)
		}
	}
}

type Query struct {
	Type, R, C, RA, CA, RB, CB int64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	H, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	W, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	q := make([]Query, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		q[i].Type, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		if q[i].Type == 1 {
			scanner.Scan()
			q[i].R, err = strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
			scanner.Scan()
			q[i].C, err = strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
			continue
		}
		scanner.Scan()
		q[i].RA, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		q[i].CA, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		q[i].RB, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		q[i].CB, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(H, W, Q, q)
}
