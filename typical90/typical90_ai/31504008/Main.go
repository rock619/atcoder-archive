package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type TreeDFS struct {
	Tree      [][]int64
	Result    []TreeDFSResult
	Timestamp int64
}

func NewTreeDFS(tree [][]int64) *TreeDFS {
	return &TreeDFS{
		Tree:      tree,
		Result:    make([]TreeDFSResult, len(tree)),
		Timestamp: 0,
	}
}

func (t *TreeDFS) Do(init int64) []TreeDFSResult {
	t.do(init, -1, 0)
	return t.Result
}

func (t *TreeDFS) do(current, parent, depth int64) {
	t.Result[current].Parent = parent
	t.Result[current].Depth = depth
	t.Result[current].Timestamp = t.Timestamp

	t.Timestamp++

	for _, c := range t.Tree[current] {
		if c != parent {
			t.do(c, current, depth+1)
		}
	}
}

type TreeDFSResult struct {
	Parent    int64
	Depth     int64
	Timestamp int64
}

func solve(N int64, A, B []int64, Q int64, K []int64, V [][]int64) {
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

	dfsResult := NewTreeDFS(graph).Do(0)

	log := int64(0)
	for ; (1 << log) < N; log++ {
	}
	doubling := make([][]int64, log)
	doubling[0] = make([]int64, N)
	for i := int64(0); i < N; i++ {
		doubling[0][i] = dfsResult[i].Parent
	}
	for i := int64(1); i < log; i++ {
		doubling[i] = make([]int64, N)
		for j := int64(0); j < N; j++ {
			if doubling[i-1][j] == -1 {
				doubling[i][j] = -1
			} else {
				doubling[i][j] = doubling[i-1][doubling[i-1][j]]
			}
		}
	}

	dist := func(a, b int64) int64 {
		return distance(a, b, dfsResult[a].Depth, dfsResult[b].Depth, log, doubling, dfsResult)
	}

	for i := int64(0); i < Q; i++ {
		vs := make([]int64, K[i])
		for j, v := range V[i] {
			vs[j] = v - 1
		}
		sort.Slice(vs, func(j, k int) bool {
			return dfsResult[vs[j]].Timestamp < dfsResult[vs[k]].Timestamp
		})
		dists := int64(0)
		for j := int64(0); j < K[i]; j++ {
			k := j + 1
			if k == K[i] {
				k = 0
			}
			a := vs[j]
			b := vs[k]
			dists += dist(a, b)
		}
		fmt.Fprintln(w, dists/2)
	}
}

func distance(a, b, aDepth, bDepth, log int64, doubling [][]int64, dfsResult []TreeDFSResult) int64 {
	return aDepth + bDepth - 2*dfsResult[lca(a, b, aDepth, bDepth, log, doubling)].Depth
}

func lca(a, b, aDepth, bDepth, log int64, doubling [][]int64) int64 {
	if aDepth > bDepth {
		a, b, aDepth, bDepth = b, a, bDepth, aDepth
	}

	for i := log - 1; i >= 0; i-- {
		if bDepth-aDepth >= (1 << i) {
			b = doubling[i][b]
			bDepth -= (1 << i)
		}
	}
	if a == b {
		return a
	}

	for i := log - 1; i >= 0; i-- {
		if doubling[i][a] != doubling[i][b] {
			a, b = doubling[i][a], doubling[i][b]
		}
	}
	return doubling[0][a]
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
	A := make([]int64, N)
	B := make([]int64, N)
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
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	K := make([]int64, Q)
	V := make([][]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		K[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		V[i] = make([]int64, K[i])
		for j := int64(0); j < K[i]; j++ {
			scanner.Scan()
			V[i][j], err = strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, A, B, Q, K, V)
}
