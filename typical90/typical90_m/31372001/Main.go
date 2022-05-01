package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Edge struct {
	Weight int64
	To     int64
}

type EdgeHeap []Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *EdgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Dijkstra(graph [][]Edge, start int64) []int64 {
	dist := make([]int64, len(graph))
	for i := range dist {
		dist[i] = 1 << 62
	}
	used := make([]bool, len(graph))
	dist[start] = 0

	h := &EdgeHeap{
		{To: start, Weight: 0},
	}
	heap.Init(h)

	for h.Len() > 0 {
		edge := heap.Pop(h).(Edge)
		pos := edge.To
		if used[pos] {
			continue
		}
		used[pos] = true

		for _, p := range graph[pos] {
			if to, weight := p.To, dist[pos]+p.Weight; dist[to] > weight {
				dist[to] = weight
				heap.Push(h, Edge{Weight: dist[to], To: to})
			}
		}
	}

	return dist
}

func solve(N int64, M int64, A []int64, B []int64, C []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]Edge, N)
	for i := int64(0); i < M; i++ {
		graph[A[i]-1] = append(graph[A[i]-1], Edge{To: B[i] - 1, Weight: C[i]})
		graph[B[i]-1] = append(graph[B[i]-1], Edge{To: A[i] - 1, Weight: C[i]})
	}

	distsTo1 := Dijkstra(graph, 0)
	distsToN := Dijkstra(graph, N-1)
	for i := int64(0); i < N; i++ {
		fmt.Fprintln(w, distsTo1[i]+distsToN[i])
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
	A := make([]int64, M)
	B := make([]int64, M)
	C := make([]int64, M)
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
		scanner.Scan()
		C[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, M, A, B, C)
}
