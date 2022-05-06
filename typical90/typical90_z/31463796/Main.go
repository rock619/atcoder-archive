package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TreeDFS struct {
	Tree   [][]int64
	Result []TreeDFSResult
}

func NewTreeDFS(tree [][]int64) *TreeDFS {
	return &TreeDFS{
		Tree:   tree,
		Result: make([]TreeDFSResult, len(tree)),
	}
}

func (t *TreeDFS) Do(init int64) []TreeDFSResult {
	t.do(init, -1, 0)
	return t.Result
}

func (t *TreeDFS) do(current, parent, depth int64) {
	t.Result[current].Depth = depth

	for _, c := range t.Tree[current] {
		if c != parent {
			t.do(c, current, depth+1)
		}
	}

	t.Result[current].SubTreeSize = 1
	for _, c := range t.Tree[current] {
		if c != parent {
			t.Result[current].SubTreeSize += t.Result[c].SubTreeSize
		}
	}
}

type TreeDFSResult struct {
	Depth       int64
	SubTreeSize int64
}

func solve(N int64, A []int64, B []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]int64, N)
	for i := int64(0); i < N-1; i++ {
		ai, bi := A[i]-1, B[i]-1
		graph[ai] = append(graph[ai], bi)
		graph[bi] = append(graph[bi], ai)
	}

	result := NewTreeDFS(graph).Do(0)
	g0, g1 := make([]int64, 0, N/2), make([]int64, 0, N/2)
	for i, r := range result {
		if n := int64(i) + 1; r.Depth%2 == 0 {
			g0 = append(g0, n)
		} else {
			g1 = append(g1, n)
		}
	}

	nums := g0
	if len(g0) < len(g1) {
		nums = g1
	}

	for i := int64(0); i < N/2; i++ {
		if i == 0 {
			fmt.Fprint(w, nums[i])
		} else {
			fmt.Fprintf(w, " %d", nums[i])
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

	var err error
	scanner.Scan()
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	A := make([]int64, N-1)
	B := make([]int64, N-1)
	for i := int64(0); i < N-1; i++ {
		scanner.Scan()
		A[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		B[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, A, B)
}
