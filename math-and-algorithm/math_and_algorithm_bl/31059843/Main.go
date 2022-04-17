package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Path struct {
	Cost int64
	To   int64
}

type PathHeap []Path

func (h PathHeap) Len() int           { return len(h) }
func (h PathHeap) Less(i, j int) bool { return h[i].Cost < h[j].Cost }
func (h PathHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PathHeap) Push(x interface{}) {
	*h = append(*h, x.(Path))
}

func (h *PathHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func dijkstra(N int64, graph map[int64][]Path) int64 {
	dist := make([]int64, N+1)
	for i := int64(1); i <= N; i++ {
		dist[i] = 1 << 60
	}
	used := make([]bool, N+1)
	dist[1] = 0

	h := &PathHeap{
		{To: 1, Cost: 0},
	}
	heap.Init(h)

	for h.Len() > 0 {
		path := heap.Pop(h).(Path)
		pos := path.To
		if used[pos] {
			continue
		}
		used[pos] = true

		for _, p := range graph[pos] {
			if to, cost := p.To, dist[pos]+p.Cost; dist[to] > cost {
				dist[to] = cost
				heap.Push(h, Path{Cost: dist[to], To: to})
			}
		}
	}

	if dist[N] == (1 << 60) {
		return -1
	}
	return dist[N]
}

func solve(N int64, M int64, A []int64, B []int64, C []int64) {
	graph := make(map[int64][]Path, N)
	for i := range A {
		graph[A[i]] = append(graph[A[i]], Path{To: B[i], Cost: C[i]})
		graph[B[i]] = append(graph[B[i]], Path{To: A[i], Cost: C[i]})
	}

	fmt.Println(dijkstra(N, graph))
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
	var M int64
	scanner.Scan()
	M, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, M)
	B := make([]int64, M)
	C := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		B[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		C[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, M, A, B, C)
}
