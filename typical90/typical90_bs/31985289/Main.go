package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Queue []int64

func (q *Queue) Enqueue(v int64) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() int64 {
	if q.Empty() {
		panic("*Queue.Dequeue(): empty")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q Queue) Size() int64 {
	return int64(len(q))
}

func (q Queue) Empty() bool {
	return q.Size() == 0
}

func (q *Queue) Clear() {
	*q = (*q)[:0]
}

type DFS struct {
	N, K        int64
	graph       [][]int64
	results     [][]int64
	permutation []int64
	in          []int64
	st          []int64 // FIXME
}

func NewDFS(graph [][]int64, N, K int64, in []int64) *DFS {
	permutation := make([]int64, len(graph))
	for i := range permutation {
		permutation[i] = -1
	}
	st := make([]int64, 0, N)
	for i := int64(0); i < N; i++ {
		if in[i] == 0 {
			st = append(st, i)
		}
	}
	return &DFS{
		N:           N,
		K:           K,
		graph:       graph,
		results:     make([][]int64, 0, K),
		permutation: permutation,
		in:          in,
		st:          st,
	}
}

func (dfs *DFS) Do(depth int64) bool {
	if depth == dfs.N {
		result := make([]int64, dfs.N)
		copy(result, dfs.permutation)
		dfs.results = append(dfs.results, result)
		return true
	}
	if len(dfs.st) == 0 {
		return false
	}

	for i := len(dfs.st) - 1; i >= 0; i-- {
		if int64(len(dfs.results)) == dfs.K {
			break
		}

		x := dfs.st[i]
		dfs.st = append(dfs.st[:i], dfs.st[i+1:]...)
		for _, j := range dfs.graph[x] {
			dfs.in[j]--
			if dfs.in[j] == 0 {
				dfs.st = append(dfs.st, j)
			}
		}

		dfs.permutation[depth] = x
		sign := dfs.Do(depth + 1)
		if !sign {
			return false
		}
		for _, j := range dfs.graph[x] {
			if dfs.in[j] == 0 {
				dfs.st = dfs.st[:len(dfs.st)-1]
			}
			dfs.in[j]++
		}
		dfs.st = append(dfs.st[:i], append([]int64{x}, dfs.st[i:]...)...)
	}
	return true
}

func solve(N int64, M int64, K int64, A []int64, B []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]int64, N)
	in := make([]int64, N)
	for i := int64(0); i < M; i++ {
		a, b := A[i]-1, B[i]-1
		graph[a] = append(graph[a], b)
		in[b]++
	}

	dfs := NewDFS(graph, N, K, in)
	dfs.Do(0)

	if int64(len(dfs.results)) != K {
		fmt.Fprintln(w, -1)
		return
	}

	for _, result := range dfs.results {
		for i, v := range result {
			if n := v + 1; i == 0 {
				fmt.Fprint(w, n)
			} else {
				fmt.Fprintf(w, " %d", n)
			}
		}
		fmt.Fprintln(w)
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
	M, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	K, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	A := make([]int64, M)
	B := make([]int64, M)
	for i := int64(0); i < M; i++ {
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

	solve(N, M, K, A, B)
}
