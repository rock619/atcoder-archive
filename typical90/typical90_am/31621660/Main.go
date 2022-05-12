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

func solve(N int64, a []int64, b []int64) {
	graph := make([][]int64, N)
	for i := int64(0); i < N-1; i++ {
		ai, bi := a[i]-1, b[i]-1
		graph[ai] = append(graph[ai], bi)
		graph[bi] = append(graph[bi], ai)
	}

	result := NewTreeDFS(graph).Do(0)

	sum := int64(0)
	for i := 1; i < len(result); i++ {
		s := result[i].SubTreeSize
		sum += s * (N - s)
	}
	fmt.Println(sum)
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
	a := make([]int64, N-1)
	b := make([]int64, N-1)
	for i := int64(0); i < N-1; i++ {
		scanner.Scan()
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, a, b)
}
