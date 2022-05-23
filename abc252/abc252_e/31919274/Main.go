package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type DijkstraEdge struct {
	Index  int64
	Weight int64
	To     int64
}

type DijkstraEdgeHeap []DijkstraEdge

func (h DijkstraEdgeHeap) Len() int           { return len(h) }
func (h DijkstraEdgeHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h DijkstraEdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DijkstraEdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(DijkstraEdge))
}

func (h *DijkstraEdgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Dijkstra(graph [][]DijkstraEdge, start int64) []int64 {
	indexes := make([]int64, len(graph))
	dist := make([]int64, len(graph))
	for i := range dist {
		dist[i] = 1 << 62
	}
	used := make([]bool, len(graph))
	dist[start] = 0

	h := &DijkstraEdgeHeap{
		{To: start, Weight: 0},
	}
	heap.Init(h)

	for h.Len() > 0 {
		edge := heap.Pop(h).(DijkstraEdge)
		pos := edge.To
		if used[pos] {
			continue
		}
		used[pos] = true

		for _, p := range graph[pos] {
			if to, weight := p.To, dist[pos]+p.Weight; dist[to] > weight {
				dist[to] = weight
				heap.Push(h, DijkstraEdge{Weight: dist[to], To: to})
				indexes[to] = p.Index
			}
		}
	}

	return indexes
}

func solve(N int64, M int64, A []int64, B []int64, C []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	graph := make([][]DijkstraEdge, N)
	for i := int64(0); i < M; i++ {
		s, t := A[i]-1, B[i]-1
		graph[s] = append(graph[s], DijkstraEdge{Index: i, Weight: C[i], To: t})
		graph[t] = append(graph[t], DijkstraEdge{Index: i, Weight: C[i], To: s})
	}

	indexes := Dijkstra(graph, 0)

	for i := 1; i < len(indexes); i++ {
		if id := indexes[i] + 1; i == 1 {
			fmt.Fprint(w, id)
		} else {
			fmt.Fprintf(w, " %d", id)
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
